package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)

	count, error := resp.Body.Read(bs)

	fmt.Println(count)
	fmt.Println(error)
	fmt.Println(string(bs))

}
