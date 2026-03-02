package main
import "fmt"
func insertion_sort(arr []int)[]int{
	n:=len(arr)
	for i:=1;i<n;i++{
		for j:=i;j>0 && arr[j-1]>arr[j];j--{
			arr[j],arr[j-1] = arr[j-1],arr[j]
		}
	}
	return arr
}
func main(){
	arr := []int{40, 8, 25, 76, 32, 20, 2, 0}
    answer := insertion_sort(arr)
    fmt.Println(answer)
}