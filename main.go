package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/exec"
	"time"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	// open file for reading
	// read line by line
	lines, err := readLines("gophers-gitclone.txt")
	if err != nil {

		log.Fatalf("readLines: %s", err)
	}
	// print file contents
	// for i, line := range lines {
	// 	fmt.Println(i, line)
	// }
	сtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(сtx, "sleep", "5").Run(); err != nil {
		for i, line := range lines {
			log.Printf("line %v", i)
			go func() {
				cmd := exec.Command(line)
				log.Printf("%s", line)
				err := cmd.Run()
				log.Printf("Command finished with error %v", err)
			}()
		}
	}

}
