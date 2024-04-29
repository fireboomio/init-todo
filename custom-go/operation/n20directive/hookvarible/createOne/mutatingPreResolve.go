package createOne

import (
	"custom-go/generated"
	"custom-go/pkg/types"
	"strings"
)

func MutatingPreResolve(hook *types.HookRequest, body generated.N20directive__hookvarible__createOneBody) (resp generated.N20directive__hookvarible__createOneBody, err error) {
	if likes := body.Input.Likes; len(likes) > 0 {
		var likesArr []string
		for _, like := range likes {
			likesArr = append(likesArr, like.Email)
		}
		body.Input.Content += "(" + strings.Join(likesArr, ",") + ")"
	}
	return body, nil
}
