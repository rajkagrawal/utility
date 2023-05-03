package strconv

import (
	"strconv"
	"strings"
)

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func ConvertEngAlphabetToDecimalPos(s string) int {
	d := []rune(strings.ToLower(s))[0] - 96
	return int(d)
}

func StringToIntMust(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func StringConcatenate(str1 string, str ...string) string {
	sb := strings.Builder{}
	sb.WriteString(str1)
	for _, val := range str {
		sb.WriteString(val)
	}
	return sb.String()
}
