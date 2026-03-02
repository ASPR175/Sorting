package main
import "fmt"
func bubble_sort(arr []int)[]int{
	n:=len(arr)
	for i:=n-1;i>=0;i--{
		isSorted:=true
		for j:=0;j<=i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				isSorted = false
			}
		}
		if isSorted{
			break
		}
	}
	return arr
}
func main() {
	arr := []int{4, 45, 4, 60, 3, 20, 32, 17}
    answer := bubble_sort(arr)
    fmt.Println(answer)
}