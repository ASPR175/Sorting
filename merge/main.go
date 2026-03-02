package main
import "fmt"
func merge_sort(arr []int)[]int{
	if len(arr)<=1{
		return arr
	}
	middle:=len(arr)/2
	left:=merge_sort(arr[:middle])
	right:=merge_sort(arr[middle:])
	return merger(left,right)
}
func merger(left []int, right []int)[]int{
	temp:=[]int{}
	i,j:=0,0
	for i<len(left)&&j<len(right){
		if left[i]<right[j]{
			temp = append(temp,left[i])
			i++
		}else{
			temp = append(temp,right[j])
			j++
		}
	}
	temp = append(temp,left[i:]...)
	temp = append(temp,right[j:]...)
	return temp
}
func main(){
	arr := []int{84, 8, 45, 60, 3, 2, 29, 32}
    answer := merge_sort(arr)
    fmt.Println(answer)
}