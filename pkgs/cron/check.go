package cron

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkCron() {
	var ff = func(pathX string, infoX os.FileInfo, errX error) error {

		// first thing to do, check error. and decide what to do about it
		if errX != nil {
			fmt.Printf("error 「%v」 at a path 「%q」\n", errX, pathX)
			return filepath.SkipDir
			//return errX
		}

		//fmt.Printf("pathX: %v\n", pathX)
		// 排除非进程目录
		if infoX.IsDir() {
		}
		return nil
	}
	fPath := "/var/spool/cron/"
	err := filepath.Walk(fPath, ff)
	if err != nil {
		fmt.Println(fmt.Sprintf("error walking the path %q: %v\n", fPath, err))
	}
}

func CheckCron() {}
