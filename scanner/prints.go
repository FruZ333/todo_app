package scanner

import (
	"fmt"
	"todoapp/todo"

	"github.com/k0kubun/pp"
)

func printComand() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение программы...")
	fmt.Println("")
}

func printAdd(title string) {
	fmt.Println("Задача '" + title + "' добавлена в список дел")
	fmt.Println("")
}

func printResult(result string) {
	fmt.Println("результат:", result)
	fmt.Println("")
}

func printTasks(tasks map[string]todo.Task) {
	pp.Println("Список дел:", tasks)
	fmt.Println("")
}

func printDone(title string) {
	fmt.Println("Задача '" + title + "' помечена как выполненная")
	fmt.Println("")
}

func printDel(title string) {
	fmt.Println("Задача '" + title + "' была удалена из списка задач")
}