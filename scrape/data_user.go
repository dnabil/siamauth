package scrape

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dnabil/siamauth/siamerr"
	models "github.com/dnabil/siamauth/siammodel"
	"github.com/dnabil/siamauth/util"
)

func ScrapeDataUser(r io.Reader) (models.UserData, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return models.UserData{}, err
	}

	bioElement := doc.Find("div[class=\"bio-info\"]")
	if bioElement.Length() == 0 {
		return models.UserData{}, siamerr.ErrNoElement
	}

	result := make([]string, 8)
	divs := bioElement.Children()

	divs.Each(func(i int, s *goquery.Selection) {
		each := util.TrimSpace(s.Text())
		if each != "PDDIKTI KEMDIKBUDDetail" {
			result[i] = each
		}
	})

	userData := models.UserData{}

	userData.NIM = util.TrimSpace(result[0])
	userData.Nama = util.TrimSpace(result[1])
	// result2 = Jenjang/Fakultas--/--
	jenjangFakultas := strings.Split(result[2][16:], "/")
	userData.Jenjang = jenjangFakultas[0]
	userData.Fakultas = jenjangFakultas[1]
	userData.Jurusan = util.TrimSpace(result[3][7:])
	userData.ProgramStudi = util.TrimSpace(result[4][13:])
	userData.Seleksi = util.TrimSpace(result[5][7:])
	fmt.Println(userData.Seleksi)
	userData.NomorUjian = util.TrimSpace(result[6][11:])
	userData.FotoProfil = fmt.Sprintf("https://siakad.ub.ac.id/dirfoto/foto/foto_20%s/%s.jpg", userData.NIM[0:2], userData.NIM)

	return userData, nil
}