package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func intBetween(min, max int) int {
	// TODO: impl&test
	return min
}

func randText(len int) string {
	// TODO: impl&test
	return ""
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func sysbenchOLTPInsert(tableIndex uint, k int, c, pad string) string {
	return fmt.Sprintf("INSERT INTO sb_oltp_insert%d (id, k, c, pad) VALUES (0, %d, %s, %s);", tableIndex, k, c, pad)
}

func main() {
	tables := flag.Uint("tables", 1, "how many tables")
	timeout := flag.Uint("timeout", 60, "how long it takes to trace")
	output := flag.String("output", "", "output target stream (default: stdout)")
	if flag.Parsed() != true {
		flag.Parse()
	}
	var stream *os.File
	var err error
	if *output == "" || !fileExists(*output) {
		stream = os.Stdout
	} else {
		stream, err = os.OpenFile(*output, os.O_CREATE|os.O_APPEND, os.ModePerm)
		defer stream.Close()
		handleError(err)
	}
	timeoutNotice := time.After(time.Duration(*timeout) * time.Second)
	c := randText(128)
	pad := randText(64)
	// context := context.Background()
	go func(tables uint, c, pad string, output *os.File) {
		k := intBetween(65536, 65535<<5)
		sql := sysbenchOLTPInsert(tables, k, c, pad)
		// TODO: mysql part
		term := time.Second
		ticker := time.NewTicker(term)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				output.WriteString(sql + "\n")
				// mysql.insert
			}
		}
	}(*tables, c, pad, stream)
	<-timeoutNotice
}
