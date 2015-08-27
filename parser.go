package main

// http://godoc.org/github.com/kr/logfmt#example-package--CustomHandler
import (
	// "bytes"
	"fmt"
	// "github.com/kr/logfmt"
	// "log"
	"strings"
	// "strconv"
)

type Measurement struct {
	Key  string
	Val  string
	Unit string // (e.g. ms, MB, etc)
}

type Measurements []*Measurement

var measurePrefix = []byte("connect.")

func (mm *Measurements) HandleLogfmt(key, val []byte) error {
	fmt.Println(string(key) + ": " + string(val))
	// if !bytes.HasPrefix(key, measurePrefix) {
	// 	return nil
	// }
	// i := bytes.LastIndexFunc(val, isDigit)
	// v, err := strconv.ParseFloat(string(val[:i+1]), 10)
	// if err != nil {
	// 	return err
	// }
	// m := &Measurement{
	// 	Key:  string(key[len(measurePrefix):]),
	// 	Val:  v,
	// 	Unit: string(val[i+1:]),
	// }
	// *mm = append(*mm, m)
	return nil
}

// return true if r is an ASCII digit only, as opposed to unicode.IsDigit.
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func getLineParts(line string) []string {
	output := strings.SplitAfter(line, " - ")
	headers := strings.Join(strings.Split(strings.TrimSpace(output[0]), " ")[3:6], " ")
	return []string{headers, output[1]}
}

func ProcessLine(line string) {
	parts := getLineParts(line)
	fmt.Printf("%#v", parts)
	// if headers[1] != "heroku"
	//   return
	// if headers[2] == "router"
	//  routerMetrics()
	// if headers[2] == "api"
	//  scaling()
	// else
	//  sampling()

	// mm := make(Measurements, 0)
	// if err := logfmt.Unmarshal(line, &mm); err != nil {
	// 	log.Fatalf("err=%q", err)
	// }
}
