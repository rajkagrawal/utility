package strconv

import (
	"fmt"
	"testing"
)

func TestStringToIntMust(t *testing.T) {

	i := StringToIntMust("10")
	if i!=10{
		t.Errorf("failed")
	}
	defer func() {
		if e := recover();e!=nil{
			fmt.Println(e)
		}
	}()
	i = StringToIntMust("1s1")
}

func TestConvertEngAlphabetToDecimalPos(t *testing.T) {
	if ConvertEngAlphabetToDecimalPos("b") != 2 {
		t.Fail()
	}
}
