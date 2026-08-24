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
	events       []Event
}

func NewScanner(todoList *todo.List) scanner {
	return scanner{
		todoList:     todoList,
		bufioScanner: bufio.NewScanner(os.Stdin),
		events:       make([]Event, 0),
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

			printResult(relust)
		}

		event := NewEvent(relust, inputString)
		s.events = append(s.events, event)		
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
		return s.cmdAdd(fields)
	}

	if cmd == "list" {
		return s.cmdList(fields)
	}

	if cmd == "done" {
		return s.cmdDone(fields)
	}

	if cmd == "del" {
		return s.cmdDel(fields)
	}

	if cmd == "help" {
		return s.cmdHelp(fields)
	}

	if cmd == "events" {
		return s.cmdEvent(fields)
	}

	return unknownCommand
}

func (s *scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return wrongArgs
	}

	title := fields[1]
	text := ""

	for i := 2; i < len(fields); i++ {
		text += fields[i]
		if i+1 != len(fields) {
			text += " "
		}
	}

	task := todo.NewTask(title, text)

	s.todoList.AddTask(task)

	printAdd(title)

	return ""
}

func (s *scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	tasks := s.todoList.ListTasks()
	printTasks(tasks)

	return ""
}

func (s *scanner) cmdDone(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	title := fields[1]

	task := s.todoList.DoneTask(title)
	if task != "" {
		return task
	}

	printDone(title)

	return ""
}

func (s *scanner) cmdDel(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	title := fields[1]

	result := s.todoList.DelTask(title)
	if result != "" {
		return result
	}

	printDel(title)

	return ""
}

func (s *scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	printHelp()

	return ""
}

func (s *scanner) cmdEvent(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	printEvents(s.events)

	return ""
}