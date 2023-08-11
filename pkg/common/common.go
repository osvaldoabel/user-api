package common

import (
	"strings"
)

func SliceToString(items []string, separator string) string {
	return strings.Join(items, separator)
}
