package main

import (
	"flag"
	"fmt"
	"os"

	lab2 "github.com/Rembqq/CSE"
)

var (
	inputFile    = flag.Bool("f", false, "Вказування файлу з виразом")
	inputConsole = flag.Bool("e", false, "Консольне введення")
	outputFile   = flag.Bool("o", false, "Запис виразу у файл")
)

func main() {
	flag.Parse()

	var input, pkey, wkey string = "", "", ""

	if *inputFile {
		pkey = flag.Arg(0)
		if *outputFile {
			wkey = flag.Arg(1)
		} else {
			wkey = ""
		}
	}

	if *inputConsole {
		input = flag.Arg(0)
		if *outputFile {
			wkey = flag.Arg(1)
		} else {
			wkey = ""
		}
	}

	database := lab2.ComputeHandler{YourPath: "."}
	err := database.Compute(input, pkey, wkey)
	if err != "" {
		fmt.Fprintln(os.Stderr, "Помилка: ", err)
	} else {
		fmt.Println("Виконання успішне")
	}
}
