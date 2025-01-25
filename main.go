package main

import (
	"fmt"
	"net/http"
)

// var Name string = "patrick"

// func main() {
// 	fmt.Printf("Type: %T Value: %v\n", Name, Name)
// }

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to the home of your app! \n Let's get you started!! \nMir")
	})

	http.ListenAndServe(":3000", nil)
}
