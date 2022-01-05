package structStudy

import "fmt"

type Animal interface {
	Sleep()
	GetType() string
	GetColor() string
}

type Dog struct {
	Color string
}

func (t *Dog) Sleep() {
	fmt.Println("dog is sleeping.")
}

func (t *Dog) GetColor() string {
	return t.Color
}

func (t *Dog) GetType() string{
	return "dog."
}

type Cat struct {
	Color string
}

func (t *Cat) Sleep() {
	fmt.Println("cat is sleeping.")
}

func (t *Cat) GetColor() string {
	return t.Color
}

func (t *Cat) GetType() string{
	return "cat."
}


