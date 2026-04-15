package main

import "fmt"

func main() {
	//Задание 1
	temperature := 20
	if temperature < 0 {
		fmt.Println("холодно")
	} else if temperature >= 0 && temperature <= 20 {
		fmt.Println("тепло")
	} else {
		fmt.Println("жарко")
	}
	//Задание 2
	score := 50
	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 75 && score < 89 {
		fmt.Println("Хорошо")
	} else if score >= 50 && score < 69 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Не сдал")
	}
	//Задание 3
	hour := 15
	switch {
	case hour >= 0 && hour < 5:
		fmt.Println("ночь")
	case hour >= 6 && hour < 11:
		fmt.Println("утро")
	case hour >= 12 && hour < 17:
		fmt.Println("день")
	default:
		fmt.Println("вечер")
	}
	//Задание 4
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	if number % 2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
	//Задание 5
	day := "Monday"
	switch {
	case day == "Monday" || day == "Tuesday" || day == "Wednesday" || day == "Thursday" || day == "Friday":
		fmt.Println("Будний день")
	case day == "Saturday" || day == "Sunday":
		fmt.Println("Выходной день")
	default:
		fmt.Println("Некорректный день")
	}
	//Задание 6
	balance := 1000.0
	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}
	//Задание 7
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("Ребенок")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}
	//Задание 8
	var command string
	fmt.Print("Введите команду (start/stop/restart): ")
	fmt.Scanln(&command)
	switch command {
	case "start":
		fmt.Println("Система запущена")
	case "stop":
		fmt.Println("Система остановлена")
	case "restart":
		fmt.Println("Система перезапущена")
	default:
		fmt.Println("Неизвестная команда")
	}
	//Задание 9
	grade := 4
	switch grade {
	case 5:
		fmt.Println("A")
	case 4:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 2:
		fmt.Println("D")
	case 1:
		fmt.Println("F")
	default:
		fmt.Println("Некорректная оценка")
	}

}
