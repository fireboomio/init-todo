package utils

import (
	"github.com/tidwall/gjson"
	"path/filepath"
	"regexp"
	"strings"
)

func GetStringValueWithDefault(val, defaultV string) string {
	if val == "" {
		return defaultV
	}
	return val
}

func JoinString(sep string, str ...string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.Join(str, sep)
}

func JoinPathAndToSlash(path ...string) string {
	return filepath.ToSlash(filepath.Join(path...))
}

var placeholderRegexp = regexp.MustCompile(`\${([^}]+)}`)

func ReplacePlaceholder(jsonStr, str string) string {
	return placeholderRegexp.ReplaceAllStringFunc(str, func(s string) string {
		if getValue := gjson.Get(jsonStr, s[2:len(s)-1]); getValue.Exists() {
			return getValue.String()
		}

		return s
	})
}
