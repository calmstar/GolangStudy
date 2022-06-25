package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("go协程执行中，i=", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
		fmt.Println("go协程结束")
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 4; i++ {
		number := <-c
		fmt.Println("main进程从channel中取出数据：", number)
	}
	fmt.Println("main进程结束")
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

	//hero := test.Hero{Age: 100, Level: 1000}
	//hero.Show()
	//hero.SetName("cwx")
	//fmt.Println(hero.GetName())
	//hero.Show()
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
