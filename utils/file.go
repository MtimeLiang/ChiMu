package utils

import (
	"os"
)

const (
	// host  = "http://www.main-zha.com/chimu/wine/image?name="
	// local = "C:/uploadImage/"
	host  = "http://localhost:9090/image?name="
	local = "/Users/leiliang/Desktop/wine_local_imgs/"
)

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
