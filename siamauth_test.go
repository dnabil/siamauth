package siamauth

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	user *User
)

func TestMain(m *testing.M) {
	err := godotenv.Load("test.env")
	if err != nil {
		log.Fatalln("failed to load test.env file: ", err)
	}

	user = NewUser()
	user.Login(os.Getenv("NIM"), os.Getenv("PASSWORD"))

	code := m.Run()

	if user.LoginStatus{
		user.Logout()
	}

	os.Exit(code)
}

func TestLogin(t *testing.T) {
	t.Run("TestLoginFail", func(t *testing.T) {
		user := NewUser()
		errLoginMsg, err := user.Login("212121211000423", "212121211000423")
		assert.Equal(t, ErrLoginFail, err)
		assert.NotZero(t, errLoginMsg)
		// assert.Zero(t, errLoginMsg)
	})
}

func TestGetKrs(t *testing.T) {
	krs, err := user.GetKrs()
	require.NoError(t, err)

	assert.NotZero(t, krs.MasaKRS)
	require.NotNil(t, krs.MataKuliah)
	require.NotZero(t, len(krs.MataKuliah))
	for _, v := range krs.MataKuliah{
		assert.NotZero(t, v.Kode)
		assert.NotZero(t, v.MataKuliah)
		assert.NotZero(t, v.SKS)
		assert.NotZero(t, v.Keterangan)
		assert.NotZero(t, v.Kelas)
		assert.NotZero(t, v.ProgramStudi)
	}
}