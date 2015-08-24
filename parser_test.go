package main

import (
	// "github.com/stretchr/testify/assert"
	"fmt"
	"strings"
	"testing"
)

func TestProcessLine(t *testing.T) {
	myString := "255 <158>1 2015-04-02T11:52:34.520012+00:00 host heroku router - at=info method=POST path=\"/users\" host=myapp.com request_id=c1806361-2081-42e7-a8aa-92b6808eac8e fwd=\"24.76.242.18\" dyno=web.1 connect=1ms service=37ms status=201 bytes=828"
	output := strings.SplitAfter(myString, " - ")
	// fmt.Printf("%#v", output)
	headers := strings.Split(strings.TrimSpace(output[0]), " ")[3:6]
	// if headers[1] != "heroku"
	//   return
	// if headers[2] == "router"
	//  routerMetrics()
	// if headers[2] == "api"
	//  scaling()
	// else
	//  sampling()

	fmt.Printf("%#v", headers)
	// logfmt part
	byteArray := []byte(output[1])
	ProcessLine(byteArray)
	// assert.Equal(t, expected, "prefix.test")
}
