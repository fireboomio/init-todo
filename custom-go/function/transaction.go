package function

import (
	"custom-go/generated"
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
)

func init() {
	plugins.RegisterFunction[transactionInput, transactionOutput](transaction, types.OperationType_MUTATION)
}

type (
	transactionInput struct {
		Username string          `json:"username"`
		Password string          `json:"password"`
		Info     transactionInfo `json:"info,omitempty"`
	}
	transactionInfo struct {
		Code    string `json:"code,omitempty"`
		Captcha string `json:"captcha,omitempty"`
	}
	transactionOutput struct {
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}
	transactionBody = types.OperationBody[transactionInput, transactionOutput]
)

func transaction(hook *types.HookRequest, body *transactionBody) (*transactionBody, error) {
	// ExecuteWithTransaction允许传递一个函数，在函数中执行的操作将在事务中执行
	// 函数本身会添加事务标识并在执行完成后自动提交事务
	if err := plugins.ExecuteWithTransaction(hook.InternalClient, func() error {
		createPostInput := generated.N20directive__skip_include__createIfNotExistInternalInput{
			CategoryName: "fireboom养乐多",
			Content:      "营养好味道养乐多",
			Title:        "养乐多",
			UserId:       1,
		}
		_, err := generated.N20directive__skip_include__createIfNotExist.Execute(createPostInput, hook.InternalClient)
		return err
	}); err != nil {
		return nil, err
	}

	body.ResetResponse(transactionOutput{Msg: "login success"})
	return body, nil
}
