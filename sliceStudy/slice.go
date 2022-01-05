package sliceStudy

import "fmt"

func testStaArrayValue(myArray [4]int) {
	myArray[0] = 100
}

func testDymArrayValue(myArray []int) {
	myArray[0] = 100
}

func DefineArray() {
	fmt.Println("-------------定义静态数组, 值传递-----------------")
	var array [10]int
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	array1 := [4]int{1, 2, 3, 4}
	for index, value := range array1 {
		fmt.Println("index:", index, "value:", value)
	}
	testStaArrayValue(array1)
	for index, value := range array1 {
		fmt.Println("index:", index, "value:", value)
	}

	fmt.Println("-------------定义动态数组（切片, sliceStudy, 大小可变）,作为参数时是引用传递-----------------")
	dymArray := []int{1, 2, 3, 4} //不指定大小 类型是[]int
	fmt.Printf("type is :%T\n", dymArray)
	for index, value := range dymArray {
		fmt.Println("index:", index, "value:", value)
	}
	testDymArrayValue(dymArray)
	for index, value := range dymArray {
		fmt.Println("index:", index, "value:", value)
	}
}

func DefineSlice() {
	fmt.Println("-------------定义slice的四种方式-----------------")
	var slice []int //这个只是声明一个slice，没有分配空间。没有默认值
	fmt.Printf("len = %d, sliceStudy = %v\n\n", len(slice), slice)

	slice1 := make([]int, 3)
	fmt.Println("make([]int, 3) 分配地址空间3, 默认值0")
	fmt.Printf("len = %d, sliceStudy = %v\n\n", len(slice1), slice1)
}

func AddValue2Slice() {
	slice := make([]int, 3, 4)
	fmt.Printf("len = %d, capacity = %d, sliceStudy = %v\n\n", len(slice), cap(slice), slice)
	slice = append(slice, 1)
	fmt.Printf("len = %d, capacity = %d, sliceStudy = %v\n\n", len(slice), cap(slice), slice)
	slice = append(slice, 2)
	fmt.Println("超出长度后，扩容：变为原来的两倍")
	fmt.Printf("len = %d, capacity = %d, sliceStudy = %v\n\n", len(slice), cap(slice), slice)
	for i := 0; i < 10; i++ {
		slice = append(slice, 2)
	}
	fmt.Printf("len = %d, capacity = %d, sliceStudy = %v\n\n", len(slice), cap(slice), slice)
}

