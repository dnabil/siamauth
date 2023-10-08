package main

import (
	"fmt"
	"os"

	"github.com/dnabil/siamauth"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	// ---

	fmt.Println("GetDataAndLogout, will return userdata (struct) and an err")
	dataMahasiswa, err := siamauth.NewUser().GetDataAndLogout(os.Getenv("NIM"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("result of GetDataAndLogout(): %s, %s, etc..\n", dataMahasiswa.Nama, dataMahasiswa.NIM)
	// ----

	
	// usual login, get data, get krs, etc.. :

	user := siamauth.NewUser() // constructor
	// login before doing anything else
	loginFailMsg, err := user.Login(os.Getenv("NIM"), os.Getenv("PASSWORD"))
	if err != nil {
		if err == siamauth.ErrLoginFail {
			panic("wrong credentials? login fail: " + loginFailMsg)
		} else {
			panic("error happened: " + err.Error())
		}
	}
	defer user.Logout()
	
	// when logged in, then you can start scraping siam...
	
	err = user.GetData() // user.Data will be filled if no err is given
	if err != nil {
		panic(err)
	}
	fmt.Printf("Nama: %s, Nim: %s, Fakultas: %s, etc...\n", user.Data.Nama, user.Data.NIM, user.Data.Fakultas)
	
	krs, err := user.GetKrs()
	if err != nil {
		panic(err)
	}

	fmt.Println("Masa krs: " + krs.MasaKRS)
	for i, matkul := range krs.MataKuliah {
		fmt.Printf("[%d] %v\n", i, matkul)
	}
}
