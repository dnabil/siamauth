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

## What's New

Breaking Changes (v0.2.0):

- Error variables moved to /siamauth/siamerr
- Login function now returns 2 data, error msg (string) and error

New:

- feat: GetDataAndLogout, basically: Login() GetData() and Logout()
- refactor: move scraping logic to a new pkg (/scrape)
- feat: Scrape addcourse.php logic :D

## Usage

siamauth constructor

```go
user := siamauth.NewUser()
```

After that you can use the methods available for that user :). <a href="https://pkg.go.dev/github.com/dnabil/siamauth" target="_blank">Documentation</a>

Please note that after logging in, please defer user.Logout() so the session id won't be misused.

Or if you just need the scraping logic for scraping siam pages, import the <a href="/dnabil/siamauth/tree/main/scrape" target="_blank">siamauth/scrape</a> pkg

\*I am planning on adding more features like getting schedules, study plan (aka KRS), etc.. but still lazy.
Pull requests are welcome :)
