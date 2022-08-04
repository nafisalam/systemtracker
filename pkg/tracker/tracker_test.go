package tracker

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFilesDetails(t *testing.T) {
	if err := os.MkdirAll("a/b/c/d", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	fd := FilesDetails("./a")
	fmt.Println(fd)
	if fd == nil {
		t.Error("error to execute the FileDetails method")
	}
	os.RemoveAll("./a")
}
