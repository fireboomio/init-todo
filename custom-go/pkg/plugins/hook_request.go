package plugins

import (
	"custom-go/pkg/types"
)

var WdgHooksAndServerConfig WunderGraphHooksAndServerConfig

type (
	WunderGraphHooksAndServerConfig struct {
		Webhooks       map[string]types.WebhookConfiguration
		Hooks          HooksConfiguration
		GraphqlServers []GraphQLServerConfig
		Options        types.ServerOptions
	}
	HooksConfiguration struct {
		Global         GlobalConfiguration
		Authentication AuthenticationConfiguration
		Queries        types.OperationHooks
		Mutations      types.OperationHooks
		Subscriptions  types.OperationHooks
		Uploads        map[string]UploadHooks
	}
)
