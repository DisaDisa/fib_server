package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DisaDisa/fib_server.git/fib_calc"
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
	}
	response := make([]int, y-x+1)
	for i := x; i <= y; i++ {
		newVal, err := fib_calc.GetFibNimber(i)
		if err != nil {
			panic(err)
		}
		response = append(response, newVal)
	}
	fmt.Fprint(w, response)
}

//CreateServer runs server and handle /get/{x}-{y} request
func CreateServer() {
	router := mux.NewRouter()
	router.HandleFunc("/get/{x:[0-9]+}-{y:[0-9]+}", fibHandler)
	http.Handle("/", router)

	http.ListenAndServe(":8181", nil)
	fmt.Println("Server is listening...")
}
