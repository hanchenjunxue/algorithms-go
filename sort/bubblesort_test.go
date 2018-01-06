package sort

import (
	"testing"
	"fmt"
)

func TestBubbleSort(t *testing.T) {
	fmt.Printf("before bubble sort arr is===%v\n",arr)
	BubbleSort(arr)
	fmt.Printf("after bubble sort arr is===%v\n",arr)
}
