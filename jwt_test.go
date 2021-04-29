package toolkit

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	j := NewJWT()
	params := map[string]string{
		"id":   "1",
		"name": "song",
	}
	claims := CustomClaims{
		Params: params,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Add(time.Hour * 10000).Unix()),
			Issuer:    "QH",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
}
