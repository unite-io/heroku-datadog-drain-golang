package main

import (
	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessLine(t *testing.T) {
	myString := "255 <158>1 2015-04-02T11:52:34.520012+00:00 host heroku router - at=info method=POST path=\"/users\" host=myapp.com request_id=c1806361-2081-42e7-a8aa-92b6808eac8e fwd=\"24.76.242.18\" dyno=web.1 connect=1ms service=37ms status=201 bytes=828"
	byteArray := []byte(myString)
	ProcessLine(byteArray)
	// assert.Equal(t, expected, "prefix.test")
}
