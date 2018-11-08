package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// set min and max constraints for the window size
const winMaxSize int = 10000
const winMinSize int = 100

// main function
func main() {
	// read the cli arguments and parse
	inputFile := flag.String("input", "", "The input data file")
	winSize := flag.Int("size", 100, "The size of the sliding window (min 2, max 10000)")

	flag.Parse()

	input := *inputFile
	size := *winSize

	// attempt to open file and error check
	file, err := os.Open(input)
	if err != nil {
		log.Printf("%#v", err)
		return
	}
	defer file.Close()

	if size < 2 {
		log.Println("Min window size not reached.")
		return
	}

	if size > winMaxSize {
		log.Println("Exceeded max window size.")
		return
	}

	// start the process
	err = start(file, size)
	if err != nil {
		log.Printf("%#v", err)
		return
	}
}

// start reading the file using the sliding window
func start(file *os.File, size int) error {
	win := NewSlidingWindow()
	win.Size(size)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		val, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		win.AddDelay(val)
		fmt.Println(win.Median())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
