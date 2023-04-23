package utils

import (
	"os"
)

func FilepathExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
