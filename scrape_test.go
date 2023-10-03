package siamauth

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

	assert.Equal(t, "GANJIL 2023/2024", c.MasaKRS)
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

func TestScrapeLoginError(t *testing.T) {
	file, err := os.Open("pages/index_login fail.html")
	require.NoError(t, err)
	defer file.Close()

	loginErrMsg, err := ScrapeLoginError(file)
	assert.NotZero(t, loginErrMsg)
	assert.Zero(t, err)
}

func TestScrapeKrs(t *testing.T){
	file, err := os.Open("pages/krs.html")
	require.NoError(t, err)
	defer file.Close()

	krs, err := ScrapeKrs(file)
	require.NoError(t, err)
	require.NotZero(t, krs)
	assert.NotZero(t, krs.MasaKRS)
	require.NotZero(t, len(krs.MataKuliah))
	
	for i := 0; i < len(krs.MataKuliah); i++ {
		matkul := krs.MataKuliah[i]
		assert.NotZero(t, matkul.Kelas)
		assert.NotZero(t, matkul.Keterangan)
		assert.NotZero(t, matkul.Kode)
		assert.NotZero(t, matkul.MataKuliah)
		assert.NotZero(t, matkul.ProgramStudi)
		assert.NotZero(t, matkul.SKS)
	}
}