package main

import "fmt"

func main() {
	//Exercise 1
	var age int
	age = 19
	fmt.Println(age)
	age = 20
	fmt.Println("День рождения:", age)
	//Exercise 2
	var height int
	height = 175
	fmt.Println("Рост:", height, "см")
	var height_in_meters float64
	height_in_meters = 1.75
	fmt.Println("Рост:", height_in_meters, "метров")
	//Exercise 3
	var isStudent bool
	isStudent = true
	fmt.Println("Является ли студентом:", isStudent)
	//Exercise 4
	var temperature int
	temperature = 25
	fmt.Println("Температура:", temperature, "°C")
	fmt.Println("Погода теплая:", temperature > 20)
	//Exercise 5
	var favoriteQuote string
	favoriteQuote = "I love rumors! Facts can be so deceiving, and rumors, true or false, often reveal a lot."
	fmt.Println(favoriteQuote)
	//Exercise 6
	var PI float64
	PI = 3.14
	fmt.Println(PI)
	// PI = "3.1415" // мы не можем присвоить строку переменной типа float64, потому что тип данных не совпадает.
	// fmt.Println(PI) 
}
