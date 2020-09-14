package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, er := http.Get("http://api.theysaidso.com/qod.json")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer resp.Body.Close()
	contents, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		fmt.Println(er)
		return
	}
	fmt.Println(string(contents))
}
