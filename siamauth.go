package siamauth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dnabil/siamauth/siamerr"
	"github.com/gocolly/colly"
)

var (
	loginUrl  string = "https://siam.ub.ac.id/index.php/"  //POST
	siamUrl   string = "https://siam.ub.ac.id/"            //GET
	logoutUrl string = "https://siam.ub.ac.id/logout.php/" //GET
)

type (
	User struct {
		c           *colly.Collector
		Account     Account
		LoginStatus bool
	}
	Account struct {
		NIM          string
		Nama         string
		Jenjang      string
		Fakultas     string
		Jurusan      string
		ProgramStudi string
		Seleksi      string
		NomorUjian   string
		FotoProfil   string
	}
)

// constructor
func NewUser() *User {
	return &User{c: colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36"),
	), LoginStatus: false}
}

func (s *User) AutoScrap(username, password string) (Account, error) {
	err := s.Login(username, password)
	if err != nil {
		return Account{}, err
	}

	err = s.GetData()
	if err != nil {
		return Account{}, err
	}

	err = s.Logout()
	if err != nil {
		return Account{}, err
	}

	return s.Account, nil
}

func (s *User) Login(us string, ps string) error {
	if s.LoginStatus {
		return siamerr.ErrorLoggedIn
	}

	var errLogin error
	var doc *goquery.Document

	s.c.OnResponse(func(r *colly.Response) {
		doc, errLogin = goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if errLogin != nil {
			errLogin = errors.New("couldn't read response body")
			return
		}
		temp := errors.New(strings.TrimSpace(doc.Find("small.error-code").Text()))
		if temp != nil {
			if len(temp.Error()) != 0 {
				errLogin = temp
				return
			}
		}
	})
	err := s.c.Post(loginUrl, map[string]string{
		"username": us,
		"password": ps,
		"login":    "Masuk",
	})

	if err != nil {
		if err.Error() != "Found" {
			return err
		}
	}
	if errLogin != nil {
		if len(errLogin.Error()) != 0 {
			return errLogin
		}
	}
	s.LoginStatus = true
	return nil
}

func (s *User) GetData() error {
	//scraping data mahasiswas
	result := make([]string, 8)
	s.c.OnHTML("div[class=\"bio-info\"]", func(h *colly.HTMLElement) {
		h.ForEach("div", func(i int, h *colly.HTMLElement) {
			each := strings.TrimSpace(h.Text)
			if each != "PDDIKTI KEMDIKBUDDetail" {
				result[i] = h.Text
			}
		})
	})
	err := s.c.Visit(siamUrl)
	if err != nil {
		return err
	}

	s.Account.NIM = result[0]
	s.Account.Nama = result[1]
	// result2 = Jenjang/Fakultas--/--
	jenjangFakultas := strings.Split(result[2][16:], "/")
	s.Account.Jenjang = jenjangFakultas[0]
	s.Account.Fakultas = jenjangFakultas[1]
	s.Account.Jurusan = result[3][7:]
	s.Account.ProgramStudi = result[4][13:]
	s.Account.Seleksi = result[5][7:]
	s.Account.NomorUjian = result[6][11:]
	s.Account.FotoProfil = fmt.Sprintf("https://siakad.ub.ac.id/dirfoto/foto/foto_20%s/%s.jpg", s.Account.NIM[0:2], s.Account.NIM)
	return nil
}

// make sure to defer this method after login, so the phpsessionid won't be misused
func (s *User) Logout() error {
	if !s.LoginStatus {
		return siamerr.ErrorNotLoggedIn
	}
	s.c.Visit(logoutUrl)
	return nil
}
