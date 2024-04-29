package plugins

import (
	"bytes"
	"custom-go/pkg/types"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type (
	UploadProfile   string
	UploadMetadata  interface{}
	UploadParameter struct {
		Directory string
		Profile   UploadProfile
		Metadata  UploadMetadata
		Files     []*types.UploadFile
	}
	UploadClient types.S3UploadConfiguration
)

func NewUploadClient(Name string) *UploadClient {
	client := &UploadClient{Name: Name}
	types.AddRegisteredHook(func(_ echo.Logger) {
		for _, v := range types.WdgGraphConfig.Api.S3UploadConfiguration {
			if v.Name == Name {
				client.UseSSL = v.UseSSL
				client.Endpoint = v.Endpoint
				client.BucketName = v.BucketName
				break
			}
		}
	})
	return client
}

func buildBodyWithFileFormData(data fileFormData, optional ...func(*multipart.Writer)) (body *bytes.Buffer, contentType string, err error) {
	body = new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for field, files := range data {
		for _, item := range files {
			var formFile io.Writer
			if formFile, err = writer.CreateFormFile(field, item.Name); err != nil {
				return
			}

			if _, err = io.Copy(formFile, item.Reader); err != nil {
				return
			}
		}
	}
	for _, v := range optional {
		v(writer)
	}
	if err = writer.Close(); err != nil {
		return
	}

	contentType = writer.FormDataContentType()
	return
}

var uploadHttpClient = http.Client{Timeout: 30 * time.Second}

func (u *UploadClient) Upload(parameter *UploadParameter) (uploadResp types.UploadedFiles, err error) {
	body, contentType, err := buildBodyWithFileFormData(fileFormData{"file": parameter.Files})
	if err != nil {
		return
	}

	uploadPath := types.PrivateNodeUrl + strings.ReplaceAll(string(types.InternalEndpoint_s3upload), "{provider}", u.Name)
	if len(parameter.Directory) > 0 {
		uploadPath += "?directory=" + parameter.Directory
	}
	req, err := http.NewRequest("POST", uploadPath, body)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", contentType)
	if len(parameter.Profile) > 0 {
		req.Header.Add(string(types.InternalHeader_X_Upload_Profile), string(parameter.Profile))
	}
	if parameter.Metadata != nil {
		metadataBytes, _ := json.Marshal(parameter.Metadata)
		req.Header.Add(string(types.InternalHeader_X_Metadata), string(metadataBytes))
	}

	resp, err := uploadHttpClient.Do(req)
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = fmt.Errorf("%d: %s", resp.StatusCode, string(content))
		return
	}

	err = json.Unmarshal(content, &uploadResp)
	return
}

func (u *UploadClient) GetOssUrl(key string) string {
	var ssl string
	if u.UseSSL {
		ssl = "s"
	}
	bucketName, endpoint := types.GetConfigurationVal(u.BucketName), types.GetConfigurationVal(u.Endpoint)
	return fmt.Sprintf("http%s://%s.%s/%s", ssl, bucketName, endpoint, key)
}
