package basic

import (
	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	CookieName = "_ChiMu_"
	signKey    = "com.ChiMu.Go"
)

type ResInfo struct {
	Status  int         `json:"status"`
	InfoMsg string      `json:"infoMsg"`
	Data    interface{} `json:"data"`
}

type BasicController struct {
	beego.Controller
}

func (c BasicController) Prepare() {
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
}

// New token
// c := jwt.MapClaims{"username": "liang", "exp": time.Now().Add(time.Hour * 72).Unix(),}
// token, err := c.NewToken(c)
func (ct *BasicController) NewToken(c jwt.MapClaims) (string, error) {
	return ct.newToken(signKey, c)
}

// Parse token
// t, err := ParseToken(token)
func (ct *BasicController) ParseToken(tokenStr string) (token *jwt.Token, err error) {
	return ct.parseToken(tokenStr, signKey)
}

// Get MapClaims
// m, ok := GetMapClaims(token)
func (ct *BasicController) GetMapClaims(t *jwt.Token) (m jwt.MapClaims, ok bool) {
	m, ok = interface{}(t.Claims).(jwt.MapClaims)
	return m, ok
}

func (ct *BasicController) newToken(signKey string, c jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = c
	tokenString, err := token.SignedString([]byte(signKey))
	return tokenString, err
}

func (ct *BasicController) parseToken(tokenStr, signKey string) (token *jwt.Token, err error) {
	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err == nil && jwtToken.Valid {
		return jwtToken, nil
	}
	return nil, err
}
