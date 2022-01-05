package mapStudy

import "fmt"

func DefineMap(){
	var myMap1 map[string] string
	if myMap1 == nil{
		fmt.Println("myMap1是一个空map，没有分配地址空间")
	}

	myMap2 := make(map[string]string)
	myMap2["aa"] = "1_value"
	myMap2["bb"] = "2_value"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"aa" : "1_value",
		"bb" : "2_value",
	}
	fmt.Println(myMap3)

	for index, value := range myMap3{
		fmt.Println("index:", index, "value",value)
	}


	fmt.Println("删除key==aa后")
	delete(myMap3, "aa")
	for index, value := range myMap3{
		fmt.Println("index:", index, "value",value)
	}


	fmt.Println("修改key==bb后")
	myMap3["bb"] = "update_value"
	for index, value := range myMap3{
		fmt.Println("index:", index, "value",value)
	}

	fmt.Println("增加key==cc后")
	myMap3["cc"] = "cc_update_value"
	for index, value := range myMap3{
		fmt.Println("index:", index, "value",value)
	}

}
