package structStudy

import "fmt"

type Human struct{
	name string
	age int
}

func (t *Human) SetName(name string) {
	t.name = name
}

func (t *Human) GetName() string {
	return t.name
}

func (t *Human) Eat() {
	fmt.Println("Human eat.")
}

func (t *Human) Walk() {
	fmt.Println("Human walk.")
}

type SuperMan struct {
	Human //继承
	ability string
}

func (t *SuperMan) Walk() {
	fmt.Println("SuperMan walk.")
}



