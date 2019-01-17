package main

import "strings"

func splitStringByComma(str string) []string {
	strs := strings.Split(str, ",")
	return strs
}

func trimSpaceFromString(str string) string {
	trimmedStr := strings.TrimSpace(str)
	return trimmedStr
}
