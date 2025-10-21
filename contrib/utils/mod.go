package utils

import (
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
)

func GetModName() string {
	contents := gfile.GetContents("go.mod")

	matches, err := gregex.MatchString(`module (.+)`, contents)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(matches[1])
}
