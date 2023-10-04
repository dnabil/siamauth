package siamauth

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// page: addcourse
func ScrapeAddCourse(r io.Reader) ([]AddCourse, error) {
	var courses []AddCourse

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return courses, err
	}

	// find add course table head
	tableHeader := doc.Find("tr.textWhite")
	thLength := tableHeader.Length()
	if thLength <= 0 {
		return courses, ErrNoElement
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
		return courses, ErrNoElement
	}

	tBody := tableHeader.Parent()

	// == scrape ==
	// -scraping masa krs, ex: GANJIL 2023/2024, GENAP 2023/2024
	span := doc.Find("span.section")
	if !strings.EqualFold("Jadwal Mata Kuliah Ditawarkan", trimSpace(span.Text())){
		return courses, ErrNoElement
	}

	MasaKrsText := span.Parent().Text() // [string] : [MASA KRS]
	masaArr := strings.Split(MasaKrsText, ":")
	if len(masaArr) < 2 {
		return courses, ErrNoElement
	}

	masaKrs := trimSpace(masaArr[1])
	// end of scraping masa krs

	trs := tBody.Find("tr:nth-of-type(n+2)") // skip table header
	courses = make([]AddCourse, 0, trs.Length())

	trs.Each(func(i int, s *goquery.Selection) {
		// the tds are:
		// HARI, JAM, KELAS, KODE, MATA KULIAH,
		// THN. KURIKULUM, Kuota, SKS, RUANG, JENIS, PILIH
		tds := s.Children()
		tempArr := make([]string, tds.Length()-1) // skip 'PILIH' (last td)

		// put it into array first, proccess it then append the data to courses.

		for i := 0; i < len(tempArr); i++ {
			tempArr[i] = trimSpace(tds.Eq(i).Text())
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

		courses = append(courses, AddCourse{
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

// page: akademik (dashboard)
func ScrapeDataUser(r io.Reader) (UserData, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return UserData{}, err
	}

	bioElement := doc.Find("div[class=\"bio-info\"]")
	if bioElement.Length() == 0 {
		return UserData{}, ErrNoElement
	}

	result := make([]string, 8)
	divs := bioElement.Children()

	divs.Each(func(i int, s *goquery.Selection) {
		each := trimSpace(s.Text())
		if each != "PDDIKTI KEMDIKBUDDetail" {
			result[i] = each
		}
	})

	userData := UserData{}

	userData.NIM = trimSpace(result[0])
	userData.Nama = trimSpace(result[1])
	// result2 = Jenjang/Fakultas--/--
	jenjangFakultas := strings.Split(result[2][16:], "/")
	userData.Jenjang = jenjangFakultas[0]
	userData.Fakultas = jenjangFakultas[1]
	userData.Jurusan = trimSpace(result[3][7:])
	userData.ProgramStudi = trimSpace(result[4][13:])
	userData.Seleksi = trimSpace(result[5][7:])
	userData.NomorUjian = trimSpace(result[6][11:])
	userData.FotoProfil = fmt.Sprintf("https://siakad.ub.ac.id/dirfoto/foto/foto_20%s/%s.jpg", userData.NIM[0:2], userData.NIM)

	return userData, nil
}

// page: login 
// 
// will return error login message as string and an error
//
// error may be caused by: html parse error, siam UI changes
func ScrapeLoginError(r io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	msgElement := doc.Find("small.error-code")
	if msgElement.Length() == 0 {
		return "", ErrNoElement
	}

	msgString := trimSpace(msgElement.Text())
	
	return msgString, nil
}

//page: krs
func ScrapeKrs(r io.Reader) (Krs, error){
	var krs Krs;

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return krs, err
	}

	tableHeader := doc.Find("tr.textWhite")
	thLength := tableHeader.Length()
	if thLength <= 0 {
		return krs, ErrNoElement
	}

	// iterate to find the right table header
	found := false
	for i := 0; i < thLength; i++ {
		item := tableHeader.Eq(i)
		if item.Children().Length() == 8 { // is the number of column
			// check table header col value
			if strings.EqualFold(tableHeader.Children().Text(), "NOKODENAMA MATA KULIAHSKSKELASKETERANGANBATALPRODI JADWAL") {
				tableHeader = item
				found = true
				break
			}
		}
	}
	if !found {
		return krs, ErrNoElement
	}

	tBody := tableHeader.Parent()
	
	// == scrape ==

	span := doc.Find("span.section")
	if !strings.EqualFold("Kartu Rencana Studi", trimSpace(span.Text())){
		return krs, ErrNoElement
	}

	MasaKrsText := span.Parent().Text() // [string] : [MASA KRS]
	masaArr := strings.Split(MasaKrsText, ":")
	if len(masaArr) < 2 {
		return krs, ErrNoElement
	}
	krs.MasaKRS = trimSpace(masaArr[1])

	trs := tBody.Find("tr:nth-of-type(n+2)") // skip table header
	// ignore last 2 trs (karena 2 baris tersebut data sks)
	for i := 0; i < trs.Length()-2; i++{
		tds := trs.Eq(i).Children()
		tempArr := make([]string, 0, tds.Length()-2) // skip 'BATAL' and 'NO' column

		// put it into array first, proccess it then append the data to courses.

		for j := 1; j < tds.Length(); j++{ // skip 'NO' column
			if (j == 6) { // skip 'BATAL' column
				continue
			}
			tempArr = append(tempArr, trimSpace(tds.Eq(j).Text()))
		}

		sks, _ := strconv.ParseFloat(tempArr[2], 32)
		matkul := MataKuliahKrs{
			Kode: tempArr[0],
			MataKuliah: tempArr[1],
			SKS: int(sks),
			Kelas: tempArr[3],
			Keterangan: tempArr[4],
			ProgramStudi: tempArr[5],
		}
		krs.MataKuliah = append(krs.MataKuliah, matkul)
	}
	// == end of scraping ==

	return krs, nil
}