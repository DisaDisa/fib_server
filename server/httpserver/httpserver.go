package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/DisaDisa/fib_server.git/server/fibcalc"
	"github.com/gorilla/mux"
)

func fibHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	x, err := strconv.Atoi(vars["x"])

	if err != nil {
		fmt.Fprintf(w, "X parse error")
		return
	}

	y, err := strconv.Atoi(vars["y"])

	if err != nil {
		fmt.Fprintf(w, "Y parse error")
		return
	}
	if x > y {
		fmt.Fprintf(w, "X must be less than Y")
		return
	}
	response := fibcalc.GetFibRange(x, y)
	fmt.Fprint(w, response)
}

//CreateServer runs server and handle /get/{x}-{y} request
func CreateServer(wg *sync.WaitGroup) {
	defer wg.Done()
	router := mux.NewRouter()
	router.HandleFunc("/get/{x:[0-9]+}-{y:[0-9]+}", fibHandler)
	http.Handle("/", router)

	log.Println("Server is listening...")
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatalf("ListenAndServe(): %v", err)
	}

}
