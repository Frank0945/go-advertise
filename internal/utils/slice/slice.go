package slice

import "strings"

func RmDupStrSlic(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func ConvSlicToStr(arr []string) string {
	return "{" + strings.Join(arr, ", ") + "}"
}
