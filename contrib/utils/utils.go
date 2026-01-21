package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
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

// FormatCoordinate 格式化经纬度,保留6位小数
func FormatCoordinate(coordinate string) (string, error) {
	parts := strings.Split(coordinate, ",")
	if len(parts) != 2 {
		return "", gerror.Newf("invalid coordinate format: %s, expected 'lng,lat'", coordinate)
	}

	lng := gconv.Float64(parts[0])
	lat := gconv.Float64(parts[1])

	// 验证经纬度范围
	if lng < -180 || lng > 180 || lat < -90 || lat > 90 {
		return "", gerror.Newf("coordinate out of range: lng=%f, lat=%f", lng, lat)
	}

	return fmt.Sprintf("%.6f,%.6f", lng, lat), nil
}
