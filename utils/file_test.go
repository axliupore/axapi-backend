package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateUUID(t *testing.T) {
	fmt.Println(CreateUUID())
}

func TestDirExistOrNot(t *testing.T) {
	path, _ := os.Getwd()
	fmt.Println(path)
	if DirExistOrNot(path) {
		fmt.Println("true")
	}
}
