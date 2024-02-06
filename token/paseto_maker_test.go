package token

import (
	"testing"
	"time"

	"github.com/brianw0924/simplebank/util"
	"github.com/o1egl/paseto"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	notBefore := issuedAt
	expirationTime := issuedAt.Add(duration)

	tokenString, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	jsonToken, err := maker.VerifyToken(tokenString)
	require.NoError(t, err)
	require.NotEmpty(t, jsonToken)

	jsonTokenPaseto, ok := jsonToken.(*paseto.JSONToken)
	require.True(t, ok)

	require.Equal(t, jsonTokenPaseto.Subject, username)
	require.WithinDuration(t, jsonTokenPaseto.IssuedAt, issuedAt, time.Second)
	require.WithinDuration(t, jsonTokenPaseto.NotBefore, notBefore, time.Second)
	require.WithinDuration(t, jsonTokenPaseto.Expiration, expirationTime, time.Second)
}

func TestExpiredPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := -time.Minute

	tokenString, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	jsonToken, err := maker.VerifyToken(tokenString)
	require.Error(t, err)
	require.Nil(t, jsonToken)
}
