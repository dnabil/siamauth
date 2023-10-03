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

	code := m.Run()

	if user.LoginStatus{
		user.Logout()
	}

	os.Exit(code)
}

func TestLogin(t *testing.T) {
	t.Run("TestLoginSuccess", func(t *testing.T) {
		username := os.Getenv("NIM")
		require.NotZero(t, username)
		password := os.Getenv("PASSWORD")
		require.NotZero(t, password)

		errLoginMsg, err := user.Login(username, password)
		assert.NoError(t, err)
		assert.Zero(t, errLoginMsg)
	})

	t.Run("TestLoginFail", func(t *testing.T) {
		user := NewUser()
		errLoginMsg, err := user.Login("212121211000423", "212121211000423")
		assert.Equal(t, ErrLoginFail, err)
		assert.NotZero(t, errLoginMsg)
		// assert.Zero(t, errLoginMsg)
	})
}