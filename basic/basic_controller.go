package basic

import (
	"ChiMu/utils"
	"fmt"
	"io"
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

func (c *BasicController) Prepare() {
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
}

func (c *BasicController) FileDir() string {
	return "./static/img/upload/"
}

// SaveFile SaveFile("file")
func (c *BasicController) SaveFile(fileKey string) (imgURL string, err error) {
	f, h, _ := c.GetFile(fileKey)
	time := time.Now()
	dir := c.FileDir() + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/")
	exist, _ := utils.Exists(dir)
	if !exist {
		os.MkdirAll(dir, 0777)
	}
	path := dir + h.Filename
	f.Close()

	if err := c.SaveToFile(fileKey, path); err != nil {
		return "", err
	}
	// imgURL = "http://www.main-zha.com/wine/image?name=" + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/") + h.Filename
	imgURL = "http://localhost:9090/wine/image?name=" + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/") + h.Filename
	return imgURL, nil
}

func (c *BasicController) SaveFiles(filesKey string) (imgURLs []string, err error) {
	files, err := c.GetFiles(filesKey)
	if err != nil {
		return nil, err
	}
	imgURLs = make([]string, len(files))
	for i, _ := range files {
		// for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return nil, err
		}
		time := time.Now()
		dir := c.FileDir() + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/")
		exist, _ := utils.Exists(dir)
		if !exist {
			os.MkdirAll(dir, 0777)
		}
		path := dir + files[i].Filename
		dst, err := os.Create(path)
		defer dst.Close()
		if err != nil {
			return nil, err
		}
		// copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			return nil, err
		}
		// imgURL = "http://www.main-zha.com/wine/image?name=" + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/") + files[i].Filename
		imgURL := "http://localhost:9090/wine/image?name=" + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/") + files[i].Filename
		imgURLs[i] = imgURL
	}
	return imgURLs, nil
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
