package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(6)
	hashedPass, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPass)
	require.NotEqual(t, password, hashedPass)
}

func TestComparePassword(t *testing.T) {
	password := RandomString(6)
	hashedPass, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPass)

	err = ComparePassword(hashedPass, password)
	require.NoError(t, err)

	err = ComparePassword(hashedPass, RandomString(6))
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
