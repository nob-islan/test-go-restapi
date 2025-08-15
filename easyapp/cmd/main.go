package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nob-islan/test-go-restapi/internal/handler/router"
)

func main() {
	// サーバーの起動
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router.Routing()))
}
