package main

import(
	"fmt"
	"log"
	"net/http"
)

func harryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/harry" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprint(w, "Harry Was Here")
}

func formHandler(w http.ResponseWriter, r *http.Request) { // build function to handle the POST request
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}


func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver) // url route
	http.HandleFunc("/form", formHandler) // url route and cal a function
	http.HandleFunc("/harry", harryHandler)

	fmt.Println("Starting server at 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}