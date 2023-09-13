package scrape

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/dnabil/siamauth/siamerr"
	"github.com/dnabil/siamauth/util"
)

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
		return "", siamerr.ErrNoElement
	}

	msgString := util.TrimSpace(msgElement.Text())
	
	return msgString, nil
}