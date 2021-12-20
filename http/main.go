package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//my own custom type writer interface
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("this amount of bytes have been written: ", len(bs))
	return len(bs), nil
}
