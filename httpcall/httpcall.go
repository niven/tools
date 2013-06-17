package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"net/http/httputil"
	"strings"
)

func main() {

	var err error

	if len(os.Args) < 3 {
		fmt.Println("usage: httpcall METHOD url [body content]")
		return
	}

	method := os.Args[1]
	url := os.Args[2]

	var body io.Reader

	if len(os.Args) > 3 {
		body = strings.NewReader(os.Args[3])
	}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("User-Agent", "httpcall-cli")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.ContentLength = 9
	client := &http.Client{}

	out, _ := httputil.DumpRequest(req, true)
	fmt.Println("== What we send over the wire ==")
	fmt.Println(string(out))
	fmt.Println("================================")

	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	// pretty output a bit

	fmt.Println(response.Status)
	for header, val := range response.Header {
		fmt.Println(header+":", val)
	}
	response_body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(response_body))
	response.Body.Close()
}
