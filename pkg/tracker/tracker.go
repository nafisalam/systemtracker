package tracker

import (
	"log"
	"os"
	"path/filepath"
)

var filelist []map[string]int

type FolderList []map[string]int

type FolderData struct {
	Files FolderList `json:"files"`
}

func FilesDetails(input string) interface{} {

	err := filepath.Walk(input,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			m := map[string]int{}
			m[path] = int(info.Size())
			filelist = append(filelist, m)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	fd := FolderData{filelist}
	return fd
}
