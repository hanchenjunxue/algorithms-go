package sort

import "algorithms-go/utils"

func BubbleSort(arr []int){
	length:=len(arr)
	for i:=0;i<length-1;i++{
		for j:=length-1;j>i;j--{
			if arr[j]<arr[j-1]{
				utils.Swap(j,j-1,arr)
			}
		}
	}

}
