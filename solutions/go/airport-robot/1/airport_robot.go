package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet (greeter string) string
}

type germanGreeter struct {}

func (obj germanGreeter) LanguageName() string {
    return "German"
}

func (obj germanGreeter) Greet(name string) string {
    return fmt.Sprintf("Hello %s!", name)
}

type Italian struct {}

func (obj Italian) LanguageName() string {
    return "Italian"
}

func (obj Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {}

func (obj Portuguese) LanguageName() string {
    return "Portuguese"
}

func (obj Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}

func SayHello (name string, greeter Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}
