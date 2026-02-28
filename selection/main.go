package main
import "fmt"

func selection_sort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ { 
        minIdx := i
        for j := i + 1; j < n; j++ { 
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
       
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
    return arr
}

func main() {
    arr := []int{4, 89, 45, 6, 23, 2, 22, 12}
    answer := selection_sort(arr)
    fmt.Println(answer)
}