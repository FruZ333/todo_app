package scanner

import "fmt"

func printComand() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение программы...")
	fmt.Println("")
}

func printAdd(title string) {
	fmt.Println("Задача '" + title + "' помечена как выполненная")
	fmt.Println("")
}