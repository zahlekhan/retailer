package main

import (
	"fmt"
	"github.com/joho/godotenv"
	Migration "github.com/zahlekhan/retailer/server/migration"
	Routes "github.com/zahlekhan/retailer/server/routes"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println(err)
	}

	if err := Migration.Migrations(); err != nil {
		fmt.Println(err)
	}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})
	//
	//http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
	//	fmt.Fprintf(w, "Hi")
	//})
	//
	//log.Fatal(http.ListenAndServe(":8081", nil))

	r := Routes.SetupRouter()
	//running
	_ = r.Run()

}
