package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Go Web Assembly")
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("http://localhost:9091")
	if err != nil {
		fmt.Printf("Error %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s\n", body)
}
