package got

import (
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
