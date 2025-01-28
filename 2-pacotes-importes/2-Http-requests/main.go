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
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))
	req.Body.Close()

}
