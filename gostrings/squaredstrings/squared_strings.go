package squaredstrings

import (
	"gopherland/strings/utils"
	"strings"
)

func VertMirror(s string) string {
	result := []string{}
	parts := strings.Split(s, "\n")

	for _, part := range parts {
		result = append(result, utils.ReverseString(part))
	}

	return strings.Join(result, "\n")
}

func HorMirror(s string) string {
	var result []string
	parts := strings.Split(s, "\n")

	for i := len(parts) - 1; i >= 0; i-- {
		result = append(result, parts[i])
	}

	return strings.Join(result, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}
