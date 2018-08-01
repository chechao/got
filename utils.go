package got

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DirPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func UniqueInt(is []int) []int {
	i := []int{}
	unq := map[int]bool{}
	for _, v := range is {
		if _, ok := unq[v]; ok {
		} else {
			unq[v] = true
			i = append(i, v)
		}
	}
	return i
}

func Md5(val string) string {
	m := md5.New()
	m.Write([]byte(val))
	return hex.EncodeToString(m.Sum(nil))
}
