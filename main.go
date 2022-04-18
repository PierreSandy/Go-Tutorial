package main

// used to import packages

import (
	"log"
	"os"

)

// Package List
// "fmt" allows to format basic strings, values or anything and print them to the console. ( Input / Output = I/O)
// "log" places date/timestamp onto the output of the code

// strut collection of data fields declared with data types.
type person struct {
	name string
	age  int
}

type car struct {
	vehicle string
	brand   string
	year    int
}


// Custom error loggers which
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p

}

// console Print function
func main() {

	// creates a log function that keeps track of all outputs. logs.txt file is Autocreated when runnning func first time
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	// custom log Error messages that can be created

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)

	log.Println(person{"john", 20})
	log.Println(car{"BMW", "E92", 2020})
	InfoLogger.Println("Starting First Log")
	ErrorLogger.Println("Please Check Something happened")

}
