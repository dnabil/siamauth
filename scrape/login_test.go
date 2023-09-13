package scrape

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScrapeLoginError(t *testing.T) {
	file, err := os.Open("pages/index_login fail.html")
	require.NoError(t, err)

	loginErrMsg, err := ScrapeLoginError(file)
	assert.NotZero(t, loginErrMsg)
	assert.Zero(t, err)
}