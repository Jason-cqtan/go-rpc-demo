package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Input struct {
	Name string
}

type Output struct {
	Msg string
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	data,_ := ioutil.ReadAll(r.Body)
	input := &Input{}
	_ = json.Unmarshal(data,input)
	output, _ := json.Marshal(&Output{
		Msg: "hello, " + input.Name,
	})

	fmt.Fprintf(w,"%s",string(output))
}

func main() {
	http.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("server")
}
