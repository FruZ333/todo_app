package scanner

import "time"

type Event struct {
	describe  string
	userInput string
	DateAt    time.Time
}

func NewEvent(descibe string, userinput string) Event {
	return Event{
		describe: descibe,
		userInput: userinput,
		DateAt: time.Now(),
	}
}