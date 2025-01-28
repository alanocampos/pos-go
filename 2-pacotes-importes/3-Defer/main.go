package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	// excecutado por ultimo
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))

}
