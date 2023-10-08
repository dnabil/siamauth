# SIAM UB Auth

https://siam.ub.ac.id authentication with scraping method. +scraping some other data like krs, addcourse, etc

import:

```go
import (
	"github.com/dnabil/siamauth"
)
```

run go get command:

```
go get -u github.com/dnabil/siamauth
```

## What's New

Breaking Changes (v0.3.0):

- Login function now returns 2 data, error msg (string) and error

New:

- feat: user.GetDataAndLogout, basically: Login() GetData() and Logout()
- feat: user.GetKrs & ScrapeKrs (/krs.php)
- feat: ScrapeAddCourse logic (/addcourse.php)
- the go-colly collector is now public in user's field (user.C)
- more stable & efficient (added some tests and fixed some bad logic)

## Usage

siamauth constructor

```go
user := siamauth.NewUser()
```

After that you can use the methods available for that user :D (use user.Login() first pls n then u can start web scraping siam). <a href="https://pkg.go.dev/github.com/dnabil/siamauth" target="_blank">Documentation</a> or see this <a target="_blank" href="https://github.com/dnabil/siamauth/blob/main/docs/example/main.go">code example</a>.

Please note that after logging in, please defer user.Logout() so the session id won't be misused.

Or if you just need the scraping logic for scraping siam pages (if you have the html pages), use the funcs that starts with Scrape.. like ScrapeAddCourse to scrape /addcourse.php data, etc. Just pass in the html page as io.Reader as the argument.

Pull requests are welcome :)
