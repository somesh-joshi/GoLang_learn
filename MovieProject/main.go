package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/somesh-joshi/MovieProject/routers"
)

func main() {
	r := routers.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
