package scrape

import (
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dnabil/siamauth/siamerr"
	models "github.com/dnabil/siamauth/siammodel"
	"github.com/dnabil/siamauth/util"
)



func ScrapeAddCourse(r io.Reader) ([]models.AddCourse, error) {
	var courses []models.AddCourse

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}

	// find add course table head
	tableHeader := doc.Find("tr.textWhite")
	thLength := tableHeader.Length()
	if thLength <= 0 {
		return courses, siamerr.ErrNoElement
	}

	// iterate to find the right table header
	found := false
	for i := 0; i < thLength; i++ {
		item := tableHeader.Eq(i)
		if item.Children().Length() == 11 { // 11 is the number of column
			// check table header col value
			if strings.EqualFold(tableHeader.Children().Text(), "HARIJAMKELASKODEMATA KULIAHTHN. KURIKULUMKuotaSKSRUANGJENISPILIH") {
				tableHeader = item
				found = true
				break
			}
		}
	}
	if !found {
		return courses, siamerr.ErrNoElement
	}

	tBody := tableHeader.Parent()

	// == scrape ==
	// -scraping masa krs, ex: GANJIL 2023/2024, GENAP 2023/2024
	span := doc.Find("span.section")
	if !strings.EqualFold("Jadwal Mata Kuliah Ditawarkan", util.TrimSpace(span.Text())){
		return courses, siamerr.ErrNoElement
	}

	MasaKrsText := span.Parent().Text() // [string] : [MASA KRS]
	masaArr := strings.Split(MasaKrsText, ":")
	if len(masaArr) < 2 {
		return courses, siamerr.ErrNoElement
	}

	masaKrs := util.TrimSpace(masaArr[1])
	// end of scraping masa krs

	trs := tBody.Find("tr:nth-of-type(n+2)") // skip table header
	courses = make([]models.AddCourse, 0, trs.Length())

	trs.Each(func(i int, s *goquery.Selection) {
		// the tds are:
		// HARI, JAM, KELAS, KODE, MATA KULIAH,
		// THN. KURIKULUM, Kuota, SKS, RUANG, JENIS, PILIH
		tds := s.Children()
		tempArr := make([]string, tds.Length()-1) // skip 'PILIH' (last td)

		// put it into array first, proccess it then append the data to courses.

		for i := 0; i < len(tempArr); i++ {
			tempArr[i] = util.TrimSpace(tds.Eq(i).Text())
		}

		kelasProdi := regexp.MustCompile(`^(\S+)\s+(.*)`).FindStringSubmatch(tempArr[2])
		var kelas, prodi string
		if len(kelasProdi) == 3 {
			kelas = kelasProdi[1]
			prodi = kelasProdi[2]
		}

		peminatKuota := regexp.MustCompile(`(\d+)/(\d+)`).FindStringSubmatch(tempArr[6])
		var peminat, kuota int
		if len(kelasProdi) == 3 {
			peminat, _ = strconv.Atoi(peminatKuota[1])
			kuota, _ = strconv.Atoi(peminatKuota[2])
		}

		sks, _ := strconv.Atoi(tempArr[7])

		courses = append(courses, models.AddCourse{
			MasaKRS: masaKrs,
			Hari:         tempArr[0],
			Jam:          tempArr[1],
			Kelas:        kelas,
			ProgramStudi: prodi,
			Kode:         tempArr[3],
			MataKuliah:   tempArr[4],
			TahunKurikulum: tempArr[5],
			Peminat:      peminat,
			Kuota:        kuota,
			SKS:          sks,
			Ruang:        tempArr[8],
			Jenis:        tempArr[9],
		})
	})
	// == end of scraping ==

	return courses, nil
}