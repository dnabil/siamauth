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

Breaking Changes (v.0.2.0):

- Error variables moved to /siamauth/siamerr
- Login function now returns 2 data, error msg (string) and error

New:

- feat: GetDataAndLogout, basically: Login() GetData() and Logout()
-

## Usage

siamauth constructor

```go
user := siamauth.NewUser()
```

After that you can use the methods available for that user :). <a href="https://pkg.go.dev/github.com/dnabil/siamauth" target="_blank">Documentation</a>

After logging in, please defer user.Logout() so the session id won't be misused.

\*I am planning on adding more features like getting schedules, study plan (aka KRS), etc.. but still lazy.
Pull requests are welcome :)
