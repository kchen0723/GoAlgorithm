package dp

import (
	"fmt"
	"testing"
)

func TestCalcualte(t *testing.T) {
	str := "1+((9-3)*4-10)/(6-4)+1"
	result := Calcualte(str)
	fmt.Println(result)
}
