package main

import "fmt"

func main() {
	hero := Hero{age: 100, level: 1000}
	hero.Show()
	hero.SetName("cwx")
	fmt.Println(hero.GetName())
	hero.Show()
}
