package main

import "fmt"

type Hero struct {
	name  string
	level int
	age   int
}

func (this *Hero) SetName(newName string) {
	this.name = newName
}

func (this *Hero) GetName() string {
	return this.name
}

func (this *Hero) Show() {
	fmt.Println("name=", this.name)
	fmt.Println("level=", this.level)
	fmt.Println("age=", this.age)
}

func main() {
	hero := Hero{age: 100, level: 1000}
	hero.Show()
	hero.SetName("cwx")
	fmt.Println(hero.GetName())
	hero.Show()
}

func example() {
	//cityMap := make(map[string]string)
	//cityMap["xx"] = "php"
	//cityMap["yy"] = "go"
	//printMyMap(cityMap)
	//fmt.Println("-----------")
	//delete(cityMap, "xx")
	//cityMap["yy"] = "aaaaa"
	//printMyMap(cityMap)
	//fmt.Println("-----------")
	//changeValue(cityMap)
	//printMyMap(cityMap)

	//var book Book
	//book.title = "cwx"
	//book.content = "success&rich&healthy"
	//changeBook(book)
	//fmt.Printf("detail: %v", book)
	//changeBook2(&book)
	//fmt.Printf("detail: %v", book)
}

type Book struct {
	title   string
	content string
}

func changeBook(book Book) {
	book.title = "ccc"
}

func changeBook2(book *Book) {
	book.title = "ccc"
}

func printMyMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key=", key, "value=", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["china"] = "guangzhou"
}
