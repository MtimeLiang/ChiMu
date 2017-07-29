package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

const (
	// host  = "http://www.main-zha.com/chimu/wine/image?name="
	// local = "C:/uploadImage/"
	host  = "http://localhost:9090/image?name="
	local = "/Users/leiliang/Desktop/wine_local_imgs/"
)

func SaveFile(mf multipart.File, h multipart.FileHeader) (url string, err error) {
	defer mf.Close()
	if mf == nil {
		err = errors.New("Error : Arg f is nil")
		return "", err
	}
	time := time.Now()
	dir := local + fmt.Sprint(time.Year(), "/", time.Month(), "/", time.Day(), "/")
	tofile := dir + h.Filename
	exist, _ := Exists(dir)
	if !exist {
		os.MkdirAll(dir, 0777)
	}
	rf, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer rf.Close()
	io.Copy(rf, mf)
	return tofile, nil
}

func RemoveFile(path string) error {
	if exist, _ := Exists(path); exist == true {
		return os.RemoveAll(path)
	}
	return nil
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
