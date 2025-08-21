package slidewindow

import (
	"fmt"
	"testing"
)

func TestGetMinCover(t *testing.T) {
	source := "GetMinCover"
	target := "Mov"
	result := GetMinCover(source, target)
	fmt.Println(result)
}
