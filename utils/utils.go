package utils

func Swap(a,b int,arr []int){
   temp:=arr[a]
   arr[a]=arr[b]
   arr[b]=temp
}
