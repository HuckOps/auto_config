package common

import "strings"

func RemoveRepeatedElementAndEmpty(arr []string) []string {
	newArr := make([]string, 0)
	for _, item := range arr {
		if "" == strings.TrimSpace(item) {
			continue
		}

		repeat := false
		if len(newArr) > 0 {
			for _, v := range newArr {
				if v == item {
					repeat = true
					break
				}
			}
		}

		if repeat {
			continue
		}

		newArr = append(newArr, item)
	}
	return newArr
}
