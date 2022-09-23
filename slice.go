package main

import "fmt"

func main() {
	//khai báo 1 slice
	var slice []int
	fmt.Println(slice)
	//Khai bao va khoi tao slice
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)
	//tao mot slice tu mot array
	var array = [4]int{1, 2, 3, 4} //array[1]=2 array[2]=3
	slice2 := array[1:3]           //array[1] -> array[3-1] <-> array[2]
	fmt.Println(slice2)

	slice3 := array[:] //all
	fmt.Println(slice3)

	slice4 := array[2:] //2,3
	fmt.Println(slice4)

	slice5 := array[:3] // 0,1,2
	fmt.Println(slice5)

	//tao 1 slice tu mot slice khac
	var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice7 := slice6
	fmt.Println(slice7)
}
