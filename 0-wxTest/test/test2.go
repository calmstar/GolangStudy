package main

import "fmt"

type Hero struct {
	name  string
	Level int
	Age   int
}

func (this *Hero) SetName(newName string) {
	this.name = newName
}

func (this *Hero) GetName() string {
	return this.name
}

func (this *Hero) Show() {
	fmt.Println("name=", this.name)
	fmt.Println("level=", this.Level)
	fmt.Println("age=", this.Age)
}
