package scrape

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScrapeAddCourse(t *testing.T) {
	file, err := os.Open("pages/addcourse.html")
	require.NoError(t, err)
	defer file.Close()

	courses, err := ScrapeAddCourse(file)
	require.NoError(t, err)
	require.NotEmpty(t, courses, "courses is not empty")

	// check one of the data 
	c := courses[0]
	hariValid := false
	switch strings.ToLower(c.Hari) {
	case "senin", "selasa", "rabu", "kamis", "jumat", "jum'at", "sabtu", "minggu":
		hariValid = true
	}
	assert.True(t, hariValid, "Hari should be valid")

	assert.NotZero(t, c.Jam)
	assert.NotZero(t, c.Kelas)
	assert.NotZero(t, c.ProgramStudi)
	assert.NotZero(t, c.Kode)
	assert.NotZero(t, c.MataKuliah)
	assert.NotZero(t, c.TahunKurikulum)
	// peminat gausa dicek
	assert.NotZero(t, c.Kuota)
	assert.NotZero(t, c.SKS)
	assert.NotZero(t, c.Ruang)
	assert.NotZero(t, c.Jenis)
}