package lib

import (
	"fmt"
	"testing"
)

func TestGetRandCode(t *testing.T) {
	code := GetRandCode()
	fmt.Println(code)
}
