package fileutil

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// PrintFile prints input file out to the terminal
func PrintFile(path string) {
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
