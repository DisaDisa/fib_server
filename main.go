package main

import (
	"/net/http"
	"fmt"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
}

func handleRequests() {

}
