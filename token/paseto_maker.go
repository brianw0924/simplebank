package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	jti, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	now := time.Now()
	jsonToken := paseto.JSONToken{
		Audience:   "audience",
		Issuer:     "issuer",
		Jti:        jti.String(),
		Subject:    username,
		IssuedAt:   now,
		Expiration: now.Add(duration),
		NotBefore:  now,
	}

	return maker.paseto.Encrypt(maker.symmetricKey, jsonToken, nil)
}

func (maker *PasetoMaker) VerifyToken(tokenString string) (any, error) {
	jsonToken := &paseto.JSONToken{}
	err := maker.paseto.Decrypt(tokenString, maker.symmetricKey, jsonToken, nil)
	if err != nil {
		return nil, err
	}
	err = jsonToken.Validate()
	if err != nil {
		return nil, err
	}
	return jsonToken, nil
}
