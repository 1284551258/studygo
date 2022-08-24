package split_string

import (
	"strings"
)

func Split(str string, req string) []string {

	var ret []string = make([]string, 0, strings.Count(str, req)+1)

	index := strings.Index(str, req)
	for index >= 0 {
		ret = append(ret, str[:index])
		// fmt.Println(ret)
		str = str[index+len(req):]
		// fmt.Println(str)
		index = strings.Index(str, req)
		// fmt.Println(index)
	}
	ret = append(ret, str)
	return ret
}
