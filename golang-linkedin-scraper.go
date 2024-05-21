package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://linkedin-api8.p.rapidapi.com/?username=username"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "X-RapidAPI-Key")
	req.Header.Add("X-RapidAPI-Host", "linkedin-api8.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
