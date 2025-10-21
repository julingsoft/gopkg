package xtime

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

// TimeDifference 返回本地主机和给定NTP服务器之间的时间差
func TimeDifference(server string) (time.Duration, error) {
	output, err := exec.Command("/usr/sbin/ntpdate", "-q", server).CombinedOutput()
	if err != nil {
		return time.Duration(0), err
	}

	re, _ := regexp.Compile("offset (.*) sec")
	subMatched := re.FindSubmatch(output)
	if len(subMatched) != 2 {
		return time.Duration(0), errors.New("invalid ntpdate output")
	}

	f, err := strconv.ParseFloat(string(subMatched[1]), 64)
	if err != nil {
		return time.Duration(0), err
	}

	return time.Duration(f*1000) * time.Millisecond, nil
}
