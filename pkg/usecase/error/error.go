package error

import (
	"github.com/goccy/go-json"
)

type ValidationError map[string][]string

func (ve ValidationError) Error() string {
	jsonErrors, _ := json.Marshal(ve)
	return string(jsonErrors)
}
