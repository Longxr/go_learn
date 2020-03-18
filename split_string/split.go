package splitstring

import "strings"

// Split 切割字符串
// example:
// abc, b => [a c]
func Split(str string, sep string) []string {
	//str:"a:b:c" sep=":"
	// var ret []string
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
