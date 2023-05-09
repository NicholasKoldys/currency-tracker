package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/nicholaskoldys/currency-tracker/service/repo"
	"github.com/nicholaskoldys/currency-tracker/util"
)

type apiHandler struct {
}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("Hello 111!\n"))
	w.Write([]byte("The time is: " + tm))
	hellow()
}

func hellow() {
	fmt.Print("%v", "helloworld")
}

func main() {

	fmt.Println("Developer: ", os.Getenv("DEVELOPER"))

	routerMux := http.NewServeMux()

	routerMux.Handle("/", apiHandler{})

	util.Debug(1, "Starting server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", routerMux); err != nil {
		fmt.Println(err)
	}
}
