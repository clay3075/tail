package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "./lib"
)

func main() {
	
	if (len(os.Args) < 2) {
		log.Fatal("Not enough arguments. \n\tEx) tail.exe <file path> <lines to read>")
	} 

	fileToTail := os.Args[1]
	linesToShow, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal("Invalid arguments.")
	}

	file, err := os.Open(fileToTail) // For read access.
	if err != nil {
		log.Fatal(err)
	}	
	defer file.Close()

	tailedLines := FixedSizeStack.New(linesToShow)
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        tailedLines.Push(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    tailedLines.Print()
}