package main

import (
	"io/ioutil"
	"net/http"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	res, err := http.Get("https://ipinfo.io/json")
	checkNilErr(err)
	defer res.Body.Close()

	databytes, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)
	println(string(databytes))
}
