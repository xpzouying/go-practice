package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const logfile = "/tmp/golang_read_file_line_by_line/Login.log.2017111023"

func scanFile() error {
	total := 0 // count lines
	begin := time.Now()

	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		_ = sc.Text()
		total++
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return err
	}

	log.Printf("scan file, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)

	return nil
}

func readFileLines() error {
	total := 0
	begin := time.Now()

	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		if _, err := rd.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return err
		}
		total++
	}

	log.Printf("reader read string in file, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)
	return nil
}

func readFileOnce() error {
	total := 0
	begin := time.Now()

	data, err := ioutil.ReadFile(logfile)
	if err != nil {
		return err
	}
	ss := strings.Split(string(data), "\n")
	for _, s := range ss {
		_ = s
		total++
	}

	log.Printf("read file once and split strings, time_used: %v, lines=%v\n", time.Since(begin).Seconds(), total)
	return nil
}

func main() {
	scanFile()
	readFileLines()
	readFileOnce()
}
