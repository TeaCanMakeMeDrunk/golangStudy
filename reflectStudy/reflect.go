package reflectStudy

import "fmt"
import "reflect"

type ReflectTest struct {
	Id   int
	Name string
}

func (t ReflectTest) CallReflectTest()  {
	fmt.Println("ReflectTest si called.")
	fmt.Printf("%v\n", t)
}

func Main() {
	var reTest = ReflectTest{Id: 1, Name: "TEST"}
	doFiledAndMethod(reTest)
}

func doFiledAndMethod(input interface{}) {
	// 获取输入的类型
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType:", inputType)
	// 获取输入的值
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue:", inputValue)
	// 遍历输入对象的filed
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i)
		fmt.Printf("%s: %v = %v\n", filed.Name, filed.Type, value)
	}
	// 遍历输入对象的method
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputValue.Method(i)
		fmt.Printf("%s : %v\n", m, m.Type()) //没有m.Name方法
	}
}