package models

import "strings"

func ToArray[T any](t []T, toString func(*T) string) string {
	var (
		sb   strings.Builder
		last = len(t) - 1
	)

	sb.WriteString("{")
	for index, item := range t {
		sb.WriteString("\"")
		sb.WriteString(toString(&item))
		sb.WriteString("\"")

		if index != last {
			sb.WriteString(",")
		}
	}

	sb.WriteString("}")
	return sb.String()
}
