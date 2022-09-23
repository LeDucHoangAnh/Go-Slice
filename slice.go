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

	slice8 := slice6[1:] //1->8
	fmt.Println(slice8)

	// slice => reference type
	var array1 = [5]int{1, 2, 3, 4, 5}
	slice9 := array1[:]

	slice9[0] = 9
	fmt.Println(slice9)
	fmt.Println(array1)

	//length va capacity cua slice
	countries := [...]string{"VN", "Canada", "JP", "CN", "ABC"}
	slice10 := countries[2:3]
	fmt.Println(slice10)

	fmt.Println(len(slice10)) //[JP]
	fmt.Println(cap(slice10)) // "JP", "CN", "ABC"

	// len: số lượng phần tử có trong slice
	//	capacity: số lượng phần tử underlying array bắt đầu từ vị trí start khi mà slice được tạo

	//	make, copy, append
	//make: Khai báo 1 slice cung cấp cụ thể length và cap
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	slice12 := make([]int, 2)
	fmt.Println(slice11)
	fmt.Println(len(slice12))
	fmt.Println(cap(slice12))

	//append: thêm 1 giá trị vào slice
	var slice13 []int
	slice13 = append(slice13, 100)
	fmt.Println(slice13)

}
