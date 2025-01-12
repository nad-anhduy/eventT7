package common

import (
	"regexp"
	"strings"
)

func CheckContaint(target string, condition []string) (string, bool) {

	for _, v := range condition {
		if v == target {
			return v, true
		}
	}
	return "", false
}

func RemoveEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func CheckErrorCode(str string) string {
	var re = regexp.MustCompile(`(?m)(\w\x{1ED7}\w:\S\w+)|(\w+:\S\w+)|(\w+: \S\w+)|(mã lỗi:\S\w+)|(mã:\S\w+)`)
	match := re.FindString(str)
	if match == "" {
		return match
	}
	return func(text string) string { val := strings.Split(text, ":"); return val[1] }(match)
}

func ReplaceSymbols(str, repl string) string {
	var re = regexp.MustCompile(`|‹|\*|'|_|\$|©|,|@|c\)`)
	return re.ReplaceAllString(str, repl)
}
