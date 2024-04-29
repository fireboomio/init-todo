package types

import (
	"custom-go/pkg/utils"
	"os"
	"path/filepath"
)

var WdgGraphConfig WunderGraphConfiguration

func init() {
	configJsonPath := filepath.Join("generated", "fireboom.config.json")
	_ = utils.ReadStructAndCacheFile(configJsonPath, &WdgGraphConfig)
}

func GetConfigurationVal(val *ConfigurationVariable) (result string) {
	if val == nil {
		return
	}

	switch val.Kind {
	case ConfigurationVariableKind_STATIC_CONFIGURATION_VARIABLE:
		result = val.StaticVariableContent
	case ConfigurationVariableKind_ENV_CONFIGURATION_VARIABLE:
		result = os.Getenv(val.EnvironmentVariableName)
		if result == "" && val.EnvironmentVariableDefaultValue != "" {
			result = val.EnvironmentVariableDefaultValue
		}
	case ConfigurationVariableKind_PLACEHOLDER_CONFIGURATION_VARIABLE:
	default:
		result = val.StaticVariableContent
	}
	return
}
