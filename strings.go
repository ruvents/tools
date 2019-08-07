package tools

import "strings"

func Explode(str, sep string) (str1, str2 string) {
	strComponents := strings.SplitN(str, sep, 2)
	strComponentsCount := len(strComponents)
	if strComponentsCount > 0 {
		str1 = strings.TrimSpace(strComponents[0])
		if strComponentsCount > 1 {
			str2 = strings.TrimSpace(strComponents[1])
		}
	}
	return
}
