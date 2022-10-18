package listener

import (
	"regexp"
	"strings"
)

// FirstUpper 字符串首字母大写写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]

}

func suffix(str string) string {
	return str + "Model"
}

func suffixLock(str string) string {
	return str + "Lock"
}

func suffixOnce(str string) string {
	return str + "Once"
}

func stripQuotes(s string) string {
	if s == "" || !strings.Contains(s, "\"") {
		return s
	}
	return s[1 : len(s)-1]
}

func stripArr(s string) string {
	if s == "" || s[0] != '[' || s[len(s)-1] != ']' {
		return s
	}
	return s[1 : len(s)-1]
}

func stripNewLine(s string) string {
	if s == "" || s[0] != '\n' {
		return s
	}
	return s[1:]
}

func captical(s string) string {
	if s == "" {
		return s
	}
	arr := []rune(s)
	//fmt.Println(s,arr[0],'A','a')
	if arr[0] >= 'A' && arr[0] <= 'Z' {
		return s
	}
	if arr[0] >= 'a' && arr[0] <= 'z' {
		arr[0] = arr[0] - 'a' + 'A'
		return string(arr)
	}
	return "K" + s
}

func IsNumber(k string) bool {
	f, _ := regexp.Compile("^(\\d*)$")
	return f.Match([]byte(k))
}

func camel(s string) string {
	a := strings.Split(s, "_")
	r := ""
	for _, s1 := range a {
		r += captical(s1)
	}
	return r
}

func removeEndline(s string) string {
	return strings.Replace(s, "\n", "", -1)
}
