package scrape

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScrapeDataUser(t *testing.T) {
	file, err := os.Open("pages/akademik.html")
	require.NoError(t, err)
	defer file.Close()

	data, err := ScrapeDataUser(file)
	require.NoError(t, err)
	
	assert.Equal(t, "111111111111111", data.NIM)
	assert.Equal(t, "Nama User", data.Nama)
	assert.Equal(t, "S1", data.Jenjang)
	assert.Equal(t, "Ilmu Komputer", data.Fakultas)
	assert.Equal(t, "Teknologi Informasi", data.Jurusan)
	assert.Equal(t, "Seleksi Bersama Masuk Perguruan Tinggi Negeri Brawijaya - Malang", data.Seleksi)
	assert.Equal(t, "111111111111", data.NomorUjian)
}