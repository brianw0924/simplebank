package token

import (
	"testing"
	"time"

	"github.com/brianw0924/simplebank/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	notBefore := issuedAt
	expirationTime := issuedAt.Add(duration)

	tokenString, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	claims, err := maker.VerifyToken(tokenString)
	require.NoError(t, err)
	require.NotEmpty(t, claims)

	claimsJWT, ok := claims.(jwt.Claims)
	require.True(t, ok)

	subject, err := claimsJWT.GetSubject()
	require.NoError(t, err)
	require.Equal(t, subject, username)

	claimsIssuedAt, err := claimsJWT.GetIssuedAt()
	require.NoError(t, err)
	require.WithinDuration(t, claimsIssuedAt.Time, issuedAt, time.Second)

	claimsGetNotBefore, err := claimsJWT.GetNotBefore()
	require.NoError(t, err)
	require.WithinDuration(t, claimsGetNotBefore.Time, notBefore, time.Second)

	claimsExpirationTime, err := claimsJWT.GetExpirationTime()
	require.NoError(t, err)
	require.WithinDuration(t, claimsExpirationTime.Time, expirationTime, time.Second)
}

func TestExpiredsJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := -time.Minute

	tokenString, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	claims, err := maker.VerifyToken(tokenString)
	require.Error(t, err)
	require.Nil(t, claims)
}
