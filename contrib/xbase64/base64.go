package xbase64

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/gogf/gf/v2/util/guid"
)

func DecodeString(s string) (string, error) {
	var base64Data string
	if strings.Contains(s, ",") {
		// 分离 Base64 头部信息 (例如：'data:image/png;base64,')
		parts := strings.SplitN(s, ",", 2)
		base64Data = parts[1]
	} else {
		base64Data = s
	}

	// 解码 Base64 字符串
	bytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", fmt.Errorf("解码 Base64 字符串失败: %w", err)
	}

	// 写入二进制文件
	var outputPath = os.TempDir() + "/" + guid.S() + ".jpg"
	file, err := os.Create(outputPath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %w", err)
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return "", fmt.Errorf("写入文件失败: %w", err)
	}

	return outputPath, nil
}

func EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
