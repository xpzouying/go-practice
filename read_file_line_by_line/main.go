package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

const logfile = "logfile.log"

func scanFile() {
	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	total := 0 // count lines

	begin := time.Now()
	defer func() {
		log.Printf("scan file, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)
	}()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		_ = sc.Text()
		total++
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}

func readFileLines() {
	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	total := 0
	begin := time.Now()
	defer func() {
		log.Printf("reader read string in file, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)
	}()
	rd := bufio.NewReader(f)
	for {
		if _, err := rd.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		total++
	}
}

func main() {
	scanFile()
	readFileLines()
}
