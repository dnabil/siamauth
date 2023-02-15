package main

import (
	"fmt"
	"os"

	"github.com/dnabil/siamauth"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	account, err := siamauth.NewUser().AutoScrap(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(account)
}
