package basic

import (
	"ChiMu/utils"
	"fmt"
	"os"
	"time"

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

// SaveFile SaveFile("file")
func (c BasicController) SaveFile(fileKey string) (imgURL string, err error) {
	f, h, _ := c.GetFile(fileKey)
	time := time.Now()
	dir := "./static/img/upload/" + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/")
	exist, _ := utils.Exists(dir)
	if !exist {
		os.MkdirAll(dir, 0777)
	}
	path := dir + h.Filename
	f.Close()

	if err := c.SaveToFile(fileKey, path); err != nil {
		return "", err
	}
	// imgURL = "http://www.main-zha.com/chimu/wine/image?name=" + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/") + h.Filename
	imgURL = "http://localhost:9090/chimu/wine/image?name=" + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/") + h.Filename
	return imgURL, nil
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
