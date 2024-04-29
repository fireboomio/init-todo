package types

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type (
	OperationArgsWithInput[I any] struct {
		Input   I `json:"input"`
		Context context.Context
	}
	InternalClient struct {
		ExtraHeaders RequestHeaders
		*BaseRequestBodyWg
	}
)

var randSource = rand.NewSource(time.Now().UnixNano())

func NewEmptyInternalClient() *InternalClient {
	randNumber := randSource.Int63()
	uberTraceId := fmt.Sprintf("%016x:%016x:%016x:%x", randNumber, randNumber, 0, 1)
	return &InternalClient{
		BaseRequestBodyWg: &BaseRequestBodyWg{
			ClientRequest: &WunderGraphRequest{Headers: RequestHeaders{}},
		},
		ExtraHeaders: RequestHeaders{
			string(InternalHeader_X_Request_Id):    uuid.New().String(),
			string(InternalHeader_X_uber_trace_id): uberTraceId,
		},
	}
}

func (i *InternalClient) WithHeaders(headers RequestHeaders) *InternalClient {
	if len(i.ExtraHeaders) == 0 {
		i.ExtraHeaders = headers
	} else {
		for k, v := range headers {
			i.ExtraHeaders[k] = v
		}
	}
	return i
}

func InternalClientFactoryCall(headers RequestHeaders, wg *BaseRequestBodyWg) *InternalClient {
	client := &InternalClient{
		BaseRequestBodyWg: wg,
		ExtraHeaders:      headers,
	}
	return client
}
