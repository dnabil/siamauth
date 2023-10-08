package siamauth

import (
	"bytes"

	"github.com/gocolly/colly/v2"
)

var (
	pathLogin		string = "index.php"
	pathLogout		string = "logout.php"
	pathKrs			string = "krs.php"
	// addCoursePath	string = "addcourse.php"

	urlSiam			string = "https://siam.ub.ac.id/"			//GET
	urlLogin		string = urlSiam + pathLogin				//POST
	urlLogout		string = urlSiam + pathLogout				//GET
	urlKrs			string = urlSiam + pathKrs					//GET
	// addCourseUrl	string = siamUrl + addCoursePath
)

type (
	User struct {
		C           *colly.Collector
		Data     	UserData
		LoginStatus bool
	}
)

// constructor
func NewUser() *User {
	return &User{C: colly.NewCollector(
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
	c := s.C.Clone()
	var errOnResponse error
	var loginErrMsg string

	c.OnResponse(func(r *colly.Response) {
		// TODO: check status codes (500,400,etc)

		// may visit this path if login failed
		if r.FileName() == pathLogin {
			loginErrMsg, errOnResponse = ScrapeLoginError(bytes.NewReader(r.Body))
		}
	})

	err := c.Post(urlLogin, map[string]string{
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
	c := s.cloneCollector()

	var errOnResponse error
	var data UserData

	// scraping data mahasiswas
	c.OnResponse(func(r *colly.Response) {
		data, errOnResponse = ScrapeDataUser(bytes.NewReader(r.Body))
	})
	err := c.Visit(urlSiam)
	if err := s.checkLoginStatus(); err != nil{ return err }
	if err != nil { return err }
	if errOnResponse != nil { return errOnResponse }
	s.Data = data
	return nil
}


func (s *User) GetKrs() (Krs, error){
	c := s.cloneCollector()
		
	var krs Krs
	var errOnResponse error

	c.OnResponse(func(r *colly.Response) {
		if r.FileName() == pathKrs {
			krs, errOnResponse = ScrapeKrs(bytes.NewReader(r.Body))
		}
	})
	
	err := c.Visit(urlKrs)
	if err := s.checkLoginStatus(); err != nil{ return krs, err }
	if err != nil {
		return krs, err
	}
	if errOnResponse != nil {
		return krs, errOnResponse
	}
	
	return krs, nil
}

// Make sure to defer this method after login, so the phpsessionid won't be misused
func (s *User) Logout() error {
	if !s.LoginStatus {
		return ErrNotLoggedIn
	}
	s.C.Visit(urlLogout)
	s.LoginStatus = false
	return nil
}

// use this to clone collector for every web scraping
func (s *User) cloneCollector() *colly.Collector{
	cloned := s.C.Clone()

	// insert global callbacks here :
	
	// user auth callback
	cloned.OnResponse(func(r *colly.Response) {
		if r.Request.URL.String() == urlSiam + pathLogin {
			s.LoginStatus = false
		}
	})

	// end of callbacks ---
	return cloned
}

func (s *User) checkLoginStatus() error{
	if !s.LoginStatus{
		return ErrNotLoggedIn
	}
	return nil
}
