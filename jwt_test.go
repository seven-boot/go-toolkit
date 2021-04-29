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

func TestParserToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXJhbXMiOnsiaWQiOiIxIiwibmFtZSI6InNvbmcifSwiZXhwIjoxNjU1NjY0Njg1LCJpc3MiOiJRSCIsIm5iZiI6MTYxOTY2MzY4NX0.isLU4jHbU1ydao7cTH5ra2zzNM0-QlIJD06z1OXKs2c"
	j := NewJWT()
	claims, err := j.ParserToken(token)
	if err != nil {
		t.Error(err)
		return
	}
	name := claims.Params["name"]
	t.Log(name)

}