package strconv

import (
	"strconv"
	"strings"
)

func StringToInt(str string) (int, error)   {
	return strconv.Atoi(str)
}

func ConvertEngAlphabetToDecimalPos(s string) int {
	d := []rune(strings.ToLower(s))[0] - 96
	return int(d)
}

func StringToIntMust(str string) (int)   {
	i , err :=  strconv.Atoi(str)
	if err!=nil{
		panic(err)
	}
	return i
}
