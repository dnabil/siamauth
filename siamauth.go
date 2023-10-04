package siamauth

import (
	"bytes"

	"github.com/gocolly/colly"
)

var (
	loginPath		string = "index.php"
	logoutPath		string = "logout.php"
	// addCoursePath	string = "addcourse.php"

	siamUrl			string = "https://siam.ub.ac.id/"			//GET
	loginUrl		string = siamUrl + loginPath				//POST
	logoutUrl		string = siamUrl + logoutPath				//GET
	// addCourseUrl	string = siamUrl + addCoursePath
)

type (
	User struct {
		c           *colly.Collector
		Data     	UserData
		LoginStatus bool
	}
)

// constructor
func NewUser() *User {
	return &User{c: colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36"),
	), LoginStatus: false}
}


// no need to login first & defer logout.
//
// if you just need to get the data and bounce, use this ;)
func (s *User) GetDataAndLogout(username, password string) (UserData, error) {
	errMsg, err := s.Login(username, password)
	if err != nil {
		return UserData{}, err
	}
	if errMsg != "" {
		return UserData{}, ErrLoginFail
	}
	defer s.Logout()

	err = s.GetData()
	if err != nil {
		return UserData{}, err
	}

	return s.Data, nil
} 


// Please defer Logout() after this function is called
//
// will return a login error message (from siam) and an error (already logged in/login error/siam ui changes/server down/etc)
func (s *User) Login(us string, ps string) (string, error) {
	var errOnResponse error
	var loginErrMsg string

	s.c.OnResponse(func(r *colly.Response) {
		// TODO: check status codes (500,400,etc)

		// may visit this path if login failed
		if r.FileName() == loginPath {
			response := bytes.NewReader(r.Body)
			loginErrMsg, errOnResponse = ScrapeLoginError(response)
		}
	})

	err := s.c.Post(loginUrl, map[string]string{
		"username": us,
		"password": ps,
		"login":    "Masuk",
	})
	if err != nil {
		if err.Error() != "Found" {
			return "", err
		}
	}

	if errOnResponse != nil {
		return "", errOnResponse
	}

	// login fail
	if loginErrMsg != "" {
		return loginErrMsg, ErrLoginFail
	}

	s.LoginStatus = true
	return "", nil
}

// GetData will fill in user's Data or return an error
func (s *User) GetData() error {
	if !s.LoginStatus {
		return ErrNotLoggedIn
	}
	var onScrapeErr error
	var data UserData

	// scraping data mahasiswas
	s.c.OnHTML("*", func(h *colly.HTMLElement) {
		data, onScrapeErr = ScrapeDataUser(bytes.NewReader(h.Response.Body))
	})
	err := s.c.Visit(siamUrl)
	if err != nil { return err }
	if onScrapeErr != nil { return onScrapeErr }
	s.Data = data
	return nil
}

// Make sure to defer this method after login, so the phpsessionid won't be misused
func (s *User) Logout() error {
	if !s.LoginStatus {
		return ErrNotLoggedIn
	}
	s.c.Visit(logoutUrl)
	return nil
}
