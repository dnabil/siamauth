# SIAM UB Auth

https://siam.ub.ac.id authentication with scraping method

import:

```go
import (
	"github.com/dnabil/siamauth"
)
```

run go get command:

```
go get github.com/dnabil/siamauth
```

## Usage

siamauth constructor

```go
user := siamauth.NewUser()
```

Currently there's 2 methods which is Login(uname string, pass string) and GetData()
-to use GetData(), must be logged in first otherwise will return an error.

ex:

```go
func main(){
	user := siamauth.NewUser()
	err := user.Login("NIM", "PASSWORD")
	if err != nil {
		panic(err)
	}
	err = user.GetData()
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Account.Nama)
	fmt.Println(user.Account.NIM)

	user.Logout()
}
```

the GetData() method will fill the User struct which looks like this:

```go
type User struct{
	c *colly.Collector

  	Account struct{
		NIM string
		Nama string
		Jenjang string
		Fakultas string
		Jurusan string
		ProgramStudi string
		Seleksi string
		NomorUjian string
	}

	LoginStatus bool
}
```

You can also check if the user logged in or not with LoginStatus.

That's pretty much it

# UPDATE!

You can use AutoScrap method for scrap siam because this way is efficient

```go
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
```
\*I am planning on adding more features like getting schedules, study plan (aka KRS), etc.. but still lazy.
Pull requests are welcome :)
