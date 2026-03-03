package main
import "fmt"
func quick_sort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    } 
    pivotIndex := len(arr) / 2
    pivotValue := arr[pivotIndex]

    left := []int{}
    right := []int{}

    for i, v := range arr {
        if i == pivotIndex {
            continue
        }
        if v < pivotValue {
            left = append(left, v)
        } else {
            right = append(right, v)
        }
    }

    return append(append(quick_sort(left), pivotValue), quick_sort(right)...)
}
func main(){
	arr := []int{14, 80, 75, 64, 2, 49, 22, 100}
    answer := quick_sort(arr)
    fmt.Println(answer)
}