package scanner

import (
	"bufio"
	"os"
	"strings"
	"todoapp/todo"
)

type scanner struct {
	todoList     *todo.List
	bufioScanner *bufio.Scanner
}

func NewScanner(todoList *todo.List) scanner {
	return scanner{
		todoList:     todoList,
		bufioScanner: bufio.NewScanner(os.Stdin),
	}
}

func (s *scanner) Start() {
	for {
		printComand()

		ok := s.bufioScanner.Scan()
		if !ok {
			break
		}

		inputString := s.bufioScanner.Text()

		relust := s.process(inputString)

		if relust != "" {
			if relust == needExit {
				printExit()
				return
			}

		}

	}
}

func (s *scanner) process(inputString string) string {
	fields := strings.Fields(inputString)
	if len(fields) == 0 {
		return emptyInput
	}

	cmd := fields[0]

	if cmd == "exit" {
		return needExit
	}

	if cmd == "add" {

	}
}

func (s *scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return wrongArgs
	}

	title := fields[1]
	text := ""
	
	for i := 2; i < len(fields); i++ {
		text += fields[i]
		if i + 1 != len(fields) {
			text += " "
		}
	}

	task := todo.NewTask(title, text)
	
	s.todoList.AddTask(task)

	printAdd(title)

	return ""
}


