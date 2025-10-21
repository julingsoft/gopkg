package utils

import (
	"log"

	"github.com/gogf/gf/v2/os/gfile"
	"golang.org/x/tools/imports"
)

// GoFmt formats the source file and adds or removes import statements as necessary.
func GoFmt(path string) {
	replaceFunc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			log.Printf(`error format "%s" go files: %v`, path, err)
			return content
		}
		return string(res)
	}

	var err error
	if gfile.IsFile(path) {
		// File format.
		if gfile.ExtName(path) != "go" {
			return
		}
		err = gfile.ReplaceFileFunc(replaceFunc, path)
	} else {
		// Folder format.
		err = gfile.ReplaceDirFunc(replaceFunc, path, "*.go", true)
	}
	if err != nil {
		log.Printf(`error format "%s" go files: %v`, path, err)
	}
}
