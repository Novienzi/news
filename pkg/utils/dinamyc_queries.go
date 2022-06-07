package utils

import (
	"strconv"
	"strings"
)

func SubstitutePlaceholder(data string, startInt int) (res string) {
	placeholderCount := strings.Count(data, "?")
	res = data
	for i := startInt; i < startInt+placeholderCount; i++ {
		res = strings.Replace(res, "?", "$"+strconv.Itoa(i), 1)
	}
	return res
}
