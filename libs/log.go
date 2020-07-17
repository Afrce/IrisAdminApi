package libs

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

func NewLogApp() *os.File {
	path := "./logs/"
	_ = CreateFile(path)
	filename := path + time.Now().Format("2006-01-02") + ".log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		color.Red(fmt.Sprintf("日志记录出错: %v", err))
	}

	return f
}
