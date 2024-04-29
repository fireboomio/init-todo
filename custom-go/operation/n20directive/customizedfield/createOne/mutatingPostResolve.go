package createOne

import (
	"custom-go/generated"
	"custom-go/pkg/types"
	"encoding/json"
)

func MutatingPostResolve(hook *types.HookRequest, body generated.N20directive__customizedfield__createOneBody) (resp generated.N20directive__customizedfield__createOneBody, err error) {
	body.Response.Data.Data.ContentValid = json.Valid([]byte(body.Input.Content))
	return body, nil
}
