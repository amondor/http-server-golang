package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//list all entries

func listEntry(w http.ResponseWriter, req *http.Request) {
	body, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Fprintf(w, string(body))
}

//  fill data.txt
func addEntry(author, message string) {
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(author + ":" + message + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}
}

// return now time
func index(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	} else {
		fmt.Fprintf(w, time.Now().Format("15:04"))
	}
}

// post /add
func add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	author := req.Form.Get("author")
	message := req.Form.Get("message")

	if len(author) > 0 && len(message) > 0 {
		addEntry(author, message)
		fmt.Fprintf(w, author+":"+message)
	} else {
		fmt.Println("champs vide")
	}
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/add", add)
	http.HandleFunc("/listEntry", listEntry)
	fmt.Println("Server started on http://localhost:4567")

	if err := http.ListenAndServe(":4567", nil); err != nil {
		log.Fatal(err)
	}

}
