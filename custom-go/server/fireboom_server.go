package server

import (
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
	"github.com/joho/godotenv"

	"custom-go/generated"

	operation_n20directive__customizedfield__createOne "custom-go/operation/n20directive/customizedfield/createOne"
	operation_n20directive__hookvarible__createOne "custom-go/operation/n20directive/hookvarible/createOne"
)

func init() {
	_ = godotenv.Overload("../.env")

	plugins.WdgHooksAndServerConfig = plugins.WunderGraphHooksAndServerConfig{
		Hooks: plugins.HooksConfiguration{
			Global:         plugins.GlobalConfiguration{},
			Authentication: plugins.AuthenticationConfiguration{},
			Queries:        types.OperationHooks{},
			Mutations: types.OperationHooks{
				"n20directive/customizedfield/createOne": {
					MutatingPostResolve: plugins.ConvertBodyFunc[generated.N20directive__customizedfield__createOneInternalInput, generated.N20directive__customizedfield__createOneResponseData](operation_n20directive__customizedfield__createOne.MutatingPostResolve),
				},
				"n20directive/hookvarible/createOne": {
					MutatingPreResolve: plugins.ConvertBodyFunc[generated.N20directive__hookvarible__createOneInternalInput, generated.N20directive__hookvarible__createOneResponseData](operation_n20directive__hookvarible__createOne.MutatingPreResolve),
				},
			},
		},
	}
}
