package sort

import (
	"testing"
	"fmt"
)

func TestSeletionSort(t *testing.T) {

	fmt.Printf("before sort arr is===%v\n",arr)
	SeletionSort(arr)
	fmt.Printf("after sort arr is===%v\n",arr)
}
