package sort

import "algorithms-go/utils"

func SeletionSort(arr []int){
	len:=len(arr)
	var startIndex,minIndex int
	for i:=0;i<len-1;i++{
		startIndex,minIndex=i,i
       for j:=i+1;j<len;j++{
       	if arr[j]<arr[minIndex]{
       		minIndex=j
		}
	   }

	   if startIndex!=minIndex{
	   	utils.Swap(startIndex,minIndex,arr)
	   }

	}
}
