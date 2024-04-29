package types

type (
	OperationBody[I, O any] struct {
		Canceled                bool                      `json:"canceled"`
		Op                      string                    `json:"op,omitempty"`
		Hook                    MiddlewareHook            `json:"hook,omitempty"`
		Input                   I                         `json:"input,omitempty"`
		Response                *OperationBodyResponse[O] `json:"response"`
		SetClientRequestHeaders RequestHeaders            `json:"setClientRequestHeaders,omitempty"`
	}
	OperationBodyResponse[O any] struct {
		DataAny any            `json:"dataAny,omitempty"`
		Data    O              `json:"data"`
		Errors  []RequestError `json:"errors"`
	}
)

func (o *OperationBody[I, O]) ResetResponse(data ...O) {
	o.Response = &OperationBodyResponse[O]{}
	if len(data) > 0 {
		o.Response.Data = data[0]
	}
}

type (
	OperationHookFunction  func(hook *HookRequest, body *OperationBody[any, any]) (*OperationBody[any, any], error)
	OperationHooks         map[string]OperationConfiguration
	OperationConfiguration struct {
		MockResolve         OperationHookFunction
		PreResolve          OperationHookFunction
		PostResolve         OperationHookFunction
		MutatingPreResolve  OperationHookFunction
		MutatingPostResolve OperationHookFunction
		CustomResolve       OperationHookFunction
	}
)
