package plugins

import (
	"bytes"
	"context"
	"custom-go/pkg/types"
	"custom-go/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/labstack/echo/v4"
	"github.com/r3labs/sse/v2"
	"github.com/tidwall/sjson"
	"golang.org/x/exp/maps"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

func internalRequest[I any](client *types.InternalClient, path string, options types.OperationArgsWithInput[I]) (resp *http.Response, err error) {
	if client == nil {
		client = defaultInternalClient
	}
	var (
		bodyBuffer  *bytes.Buffer
		contentType string
	)
	url := fetchInternalRequestUrl(path)
	baseBodyWg := &types.BaseRequestBodyWg{
		ClientRequest: &types.WunderGraphRequest{
			RequestURI: url,
			Method:     http.MethodPost,
			Headers:    client.ClientRequest.Headers,
		},
		User: client.User,
	}
	formData, ok := options.Context.Value(fileFormDataKey).(fileFormData)
	if ok {
		optional := func(writer *multipart.Writer) {
			inputBytes, _ := json.Marshal(options.Input)
			for _, key := range maps.Keys(formData) {
				inputBytes, _ = sjson.DeleteBytes(inputBytes, key)
			}
			_ = writer.WriteField("input", string(inputBytes))
			baseBodyWgBytes, _ := json.Marshal(baseBodyWg)
			_ = writer.WriteField("__wg", string(baseBodyWgBytes))
		}
		if bodyBuffer, contentType, err = buildBodyWithFileFormData(formData, optional); err != nil {
			return
		}
	} else {
		var jsonData []byte
		if jsonData, err = json.Marshal(types.OperationHookPayload{Input: options.Input, Wg: baseBodyWg}); err != nil {
			return
		}
		bodyBuffer, contentType = bytes.NewBuffer(jsonData), echo.MIMEApplicationJSON
	}

	req, err := http.NewRequest(http.MethodPost, url, bodyBuffer)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", contentType)
	for k, v := range client.ExtraHeaders {
		req.Header.Set(k, v)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		err = errors.New(string(bodyBytes))
		return
	}

	return
}

func executeInternalRequest[I, OD any](client *types.InternalClient, path string, input I) (result OD, err error) {
	options := types.OperationArgsWithInput[I]{Input: input}
	inputValue := reflect.ValueOf(input)
	if inputValue.Kind() == reflect.Ptr {
		inputValue = inputValue.Elem()
	}
	inputType := inputValue.Type()
	formData := make(fileFormData)
	for i := 0; i < inputValue.NumField(); i++ {
		inputFieldValue := inputValue.Field(i)
		if !inputFieldValue.IsValid() || !inputFieldValue.CanInterface() || inputFieldValue.IsZero() {
			continue
		}

		var files []*types.UploadFile
		switch v := inputFieldValue.Interface().(type) {
		case *types.UploadFile:
			files = []*types.UploadFile{v}
		case []*types.UploadFile:
			files = v
		}
		if len(files) > 0 {
			inputFieldTag := inputType.Field(i).Tag.Get("json")
			formData[inputFieldTag] = files
		}
	}
	options.Context = context.Background()
	if len(formData) > 0 {
		options.Context = context.WithValue(options.Context, fileFormDataKey, formData)
	}

	resp, err := internalRequest[I](client, path, options)
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()

	var operationResp types.OperationBodyResponse[OD]
	if err = json.NewDecoder(resp.Body).Decode(&operationResp); err != nil {
		return
	}

	if len(operationResp.Errors) > 0 {
		err = errors.New(operationResp.Errors[0].Message)
		return
	}

	result = operationResp.Data
	return
}

const fileFormDataKey = "fileFormData"

var (
	operations            = make(map[types.OperationType][]string)
	defaultInternalClient = types.NewEmptyInternalClient()
)

type (
	fileFormData   map[string][]*types.UploadFile
	Meta[I, O any] struct {
		Path string
		Type types.OperationType
	}
	Subscriber[I, O any] struct {
		Path string
	}
	SubscriberData[O any] struct {
		Data   O
		Errors []gqlerrors.FormattedError
	}
)

func FetchOperations(logger echo.Logger, operationType types.OperationType, printUrlRequired bool) []string {
	paths := operations[operationType]
	if printUrlRequired {
		for _, path := range paths {
			logger.Debugf(`Built internalRequest (%s)`, fetchInternalRequestUrl(path))
		}
	}
	return paths
}

func fetchInternalRequestUrl(path string) string {
	return types.PrivateNodeUrl + strings.ReplaceAll(string(types.InternalEndpoint_internalRequest), "{path}", path)
}

func NewOperationMeta[I, O any](path string, operationType types.OperationType) *Meta[I, O] {
	operations[operationType] = append(operations[operationType], path)
	return &Meta[I, O]{Path: path, Type: operationType}
}

func (m *Meta[I, O]) Execute(input I, client *types.InternalClient) (O, error) {
	return executeInternalRequest[I, O](client, m.Path, input)
}

func NewOperationSubscriber[I, O any](path string) *Subscriber[I, O] {
	operations[types.OperationType_SUBSCRIPTION] = append(operations[types.OperationType_SUBSCRIPTION], path)
	return &Subscriber[I, O]{Path: path}
}

func (m *Subscriber[I, O]) Subscribe(input I, client *types.InternalClient) (dataChan chan SubscriberData[O], err error) {
	options := types.OperationArgsWithInput[I]{Input: input, Context: context.Background()}
	resp, err := internalRequest[I](client, m.Path, options)
	if err != nil {
		return
	}

	dataChan = make(chan SubscriberData[O])
	go func() {
		defer func() { _ = resp.Body.Close() }()
		reader := sse.NewEventStreamReader(resp.Body, math.MaxInt)
		var (
			readMsg, lineData []byte
			data              O
		)
		for {
			if readMsg, err = reader.ReadEvent(); err != nil {
				if err == io.EOF {
					return
				}

				dataChan <- SubscriberData[O]{Errors: []gqlerrors.FormattedError{{Message: internalError}}}
				return
			}
			if len(readMsg) == 0 {
				continue
			}

			// normalize the crlf to lf to make it easier to split the lines.
			// split the line by "\n" or "\r", per the spec.
			lines := bytes.FieldsFunc(readMsg, func(r rune) bool { return r == '\n' || r == '\r' })
			for _, line := range lines {
				if bytes.HasPrefix(line, headerData) {
					if lineData = trim(line[len(headerData):]); len(lineData) == 0 {
						continue
					}
					if err = json.Unmarshal(lineData, &data); err != nil {
						dataChan <- SubscriberData[O]{Errors: []gqlerrors.FormattedError{{Message: internalError}}}
						return
					}
					dataChan <- SubscriberData[O]{Data: data}
				}
			}
		}
	}()
	return
}

func ExecuteWithTransaction(client *types.InternalClient, execute func() error) error {
	transactionId := uuid.New().String()
	client.WithHeaders(types.RequestHeaders{
		string(types.TransactionHeader_X_Transaction_Manually): "true",
		string(types.TransactionHeader_X_Transaction_Id):       transactionId,
	})
	var body []byte
	executeErr := execute()
	if executeErr != nil {
		body = []byte(fmt.Sprintf(`{"error": "%s"}`, executeErr.Error()))
	}
	url := types.PrivateNodeUrl + string(types.InternalEndpoint_internalTransaction)
	if _, err := utils.HttpPost(url, body, client.ExtraHeaders); err != nil {
		return err
	}
	return executeErr
}
