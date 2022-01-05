package main

import (
	"fmt"
	"test_project1/goroutineStudy"
	"test_project1/mapStudy"
	"test_project1/reflectStudy"
	"test_project1/sliceStudy"
	"test_project1/structStudy"
	_ "test_project1/test_package1" //匿名导入包 仅调用init方法，Go语言规定导入包必须使用其中除init之外的函数
	"time"
)

const (
	BEIJING = 11 * iota //iota只能const使用
	SHENZHEN
	GUANGZHOU
)

func foo1(a string, b int) (a1 int, a2 int) {
	a1 = 123
	a2 = 321
	return
	//return 123, 321
}

func testInterface(arg interface{}) {

	fmt.Println("Go语言中的interface{} 相当于 Java中的T，任意类型（int, float, struct{}....）")
	fmt.Println("interface{}判断是哪种类型， 通过断言判断")
	value, ok := arg.(string)
	if ok {
		fmt.Println("是string类型, 值是", value)
	} else {
		fmt.Println("不是string类型")
	}
}

func main() {
	// defer xxFunc()
	//var a = "xx"

	defer fmt.Println("defer类似Java的finally最后执行")
	fmt.Println("Hello Go!")
	time.Sleep(1)

	var aa string = "xx" //可以全局
	a := "ss"            //只能局部
	fmt.Println("A" + a + "11")
	fmt.Println("A" + aa)
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("BEIJING", BEIJING)
	fmt.Println("SHENZHEN", SHENZHEN)
	fmt.Println("GUANGZHOU", GUANGZHOU)

	ret1, ret2 := foo1("xx", 1)
	fmt.Println("ret1:", ret1, "ret2:", ret2)

	sliceStudy.DefineArray()
	sliceStudy.DefineSlice()
	sliceStudy.AddValue2Slice()

	mapStudy.DefineMap()

	fmt.Println("-------------------------------")
	human := structStudy.Human{}
	human.SetName("zhang3")
	fmt.Println(human.GetName())

	superMan := structStudy.SuperMan{}
	superMan.SetName("superMan_Li4")
	superMan.Walk()
	superMan.Eat()
	fmt.Println(superMan.GetName())
	fmt.Println("-------------------------------")

	fmt.Println("-------------------------------")
	testInterface("是的")
	testInterface(111)
	fmt.Println("-------------------------------")

	fmt.Println("-------------------------------")
	var animal structStudy.Animal
	animal = &structStudy.Dog{Color: "yellow Dog"}
	animal.Sleep()
	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())

	animal = &structStudy.Cat{Color: "white Cat"}
	animal.Sleep()
	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())
	fmt.Println("-------------------------------")

	fmt.Println("-------------------------------")
	reflectStudy.Main()
	reflectStudy.JsonConversion()
	fmt.Println("-------------------------------")

	fmt.Println("-------------------------------")
	goroutineStudy.Main()
	goroutineStudy.ChannelMain()
	goroutineStudy.ChannelSelect()
	fmt.Println("-------------------------------")

}
