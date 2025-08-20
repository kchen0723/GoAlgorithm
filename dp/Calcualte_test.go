package dp

import (
	"fmt"
	"testing"
)

func TestCalcualte(t *testing.T) {
	// str := "9-10+11"
	// str := "1+((9-3)*4-10)/(6-4)+1"
	str := "1+((3*((9-3)*4-10)/(6-4)+1)+2)/3+5"
	result := Calcualte(str)
	fmt.Println(result)
}
