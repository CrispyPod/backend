package helpers

import (
	"encoding/json"
	"slices"
	"strings"
)

func JsonValidate(headers string) bool {
	return json.Valid([]byte(headers))
}

func URLValidate(webURL string) bool {
	if strings.HasPrefix(webURL, "http://") {
		return true
	} else if strings.HasPrefix(webURL, "https://") {
		return true
	} else {
		return false
	}
}

var methods = []string{"GET", "POST", "DELETE", "PUT", "HEAD", "PATCH"}

func MethodValidate(method string) bool {
	return slices.Contains(methods, strings.ToUpper(method))
}
