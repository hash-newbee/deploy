package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/siddontang/go-mysql/client"
)

func intBetween(min, max int) int {
	// TODO: impl&test
	return min
}

func randText(len int) string {
	// TODO: impl&test
	return "test"
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

func (gen *insertGenerator) traceSysbenchOLTPInsert() string {
	// TODO: TiDB prepare statement + arguments replace SQL string
	tableIdx := intBetween(1, int(gen.tables))
	k := intBetween(65536, 65535<<5)
	c := randText(128)
	pad := randText(64)
	return fmt.Sprintf("trace format=\"row\" INSERT INTO sbtest%d (id, k, c, pad) VALUES (0, %d, \"%s\", \"%s\");", tableIdx, k, c, pad)
}

type insertGenerator struct {
	tables uint
}

func traceLoop(ctx context.Context, term time.Duration, conn *client.Conn, gen *insertGenerator, output *os.File) {
	ticker := time.NewTicker(term)
	num := 1
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			sql := gen.traceSysbenchOLTPInsert()
			r, err := conn.Execute(sql)
			handleError(err)
			defer r.Close()
			output.WriteString(fmt.Sprintf("#%d %s\n", num, sql))
			num++
			for _, row := range r.Values {
				for _, val := range row {
					output.Write(val.AsString())
					output.Write([]byte{'\t'})
				}
				output.Write([]byte{'\n'})
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	tables := flag.Uint("tables", 1, "how many tables")
	timeout := flag.Uint("timeout", 60, "how long it takes to trace")
	output := flag.String("output", "", "output target stream (default: stdout)")
	host := flag.String("host", "106.75.175.149:4000", "TiDB host IP address")
	user := flag.String("user", "root", "TiDB login user")
	password := flag.String("password", "", "TiDB login user's password")
	database := flag.String("db", "sb_oltp_insert_test", "Auto trace database")
	term := flag.Duration("term", time.Second, "SQL request term duration")
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
	gen := &insertGenerator{tables: *tables}
	conn, err := client.Connect(*host, *user, *password, *database)
	defer conn.Close()
	handleError(err)
	ctx := context.TODO()
	loopCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	timeoutNotice := time.After(time.Duration(*timeout) * time.Second)
	go traceLoop(loopCtx, *term, conn, gen, stream)
	select {
	case <-timeoutNotice:
	case <-ctx.Done():
		return
	}
}
