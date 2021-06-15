package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var Word string

type LineInfo struct {
	lineNo int
	line   string
}

func ProcFileList(paths ...string) (out chan string) {
	out = make(chan string)
	go func() {
		for _, path := range paths {
			filelist, err := filepath.Glob(path)
			if err != nil {
				continue
			}

			for _, filename := range filelist {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}

func ProcFile(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for filename := range in {
			file, err := os.Open(filename)
			if err != nil {
				continue
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				out <- scanner.Text()
			}
		}

		close(out)
	}()
	return out
}

func ProcLine(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for v := range in {
			if strings.Contains(v, Word) {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행인수필요. ex) ex26.1 word filepath")
		return
	}

	Word = os.Args[1]
	files := os.Args[2:]
	fmt.Println("WORD: ", Word, "FILE ", files)

	for n := range ProcLine(ProcFile(ProcFileList(files...))) {
		fmt.Println(n)
	}
}
