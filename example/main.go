package main

import (
	"fmt"
	"strings"

	"github.com/wNee/aidi"
)

func main() {
	fmt.Println("Aidi!\n")

	aidi.CreateCase("Test GET Go homepage").
		Get("http://golang.org").
		Send().
		ExpectStatus(200)

	body := strings.NewReader(`{"test_key":"test_value"}`)
	aidi.CreateCase("Test POST").
		Post("http://golang.org").
		SetBody(body).
		Send().
		ExpectStatus(200)

	aidi.CreateCase("Test GET Go homepage").
		Get("http://golang.org").
		Send().
		ExpectStatus(400)
	aidi.Global.PrintReport()
}