package utils

import (
	"fmt"
	"github.com/oledakotajoe/codenvi-core/config"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func InitEnvi() {
	enviRoot := GetEnviRoot()

	mErr := os.MkdirAll(enviRoot, os.ModePerm)
	if mErr != nil {
		log.Println(mErr)
	}
}

func GetHomeDir() (string, error) {
	if runtime.GOOS == "windows" {
		return filepath.Abs(fmt.Sprintf("%s\\%s", os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH")))
	} else {
		return os.Getenv("HOME"), nil
	}
}

func GetEnviRoot() string {
	homeDir, _ := GetHomeDir()
	enviRoot := homeDir + config.Global().HomeDirName
	return enviRoot
}
