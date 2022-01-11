package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type req struct {
	Text string `json:"text"`
}
type res struct {
	Sucess bool `json:"success"`
	Data   struct {
		MostUsedWord string `json:"most-used-word"`
	} `json:"data"`
}

func main() {

	http.HandleFunc("/api/most-used-word", HandleParagraph)
	http.ListenAndServe(":8080", nil)

}

func HandleParagraph(w http.ResponseWriter, r *http.Request) {

	var Request req
	json.NewDecoder(r.Body).Decode(&Request)

	fmt.Println(Request)

	wCount := make(map[string]int)

	for _, w := range strings.Fields(Request.Text) {
		c, _ := wCount[w] // if not found then value of c is 0
		wCount[w] = c + 1
	}

	max := 0
	maxWord := ""

	for w, v := range wCount {
		if v > max {
			max = v
			maxWord = w
		}
	}

	var Response res
	Response.Sucess = true
	Response.Data.MostUsedWord = maxWord
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(Response)
}
