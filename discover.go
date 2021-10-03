package gokit

import "path/filepath"

// Discover 判断dir目录下面是否有glob文件， 并且返回其绝对路径
func Discover(glob, dir string) ([]string, error) {
	var err error

	// Make the directory absolute if it isn't already
	if !filepath.IsAbs(dir) {
		dir, err = filepath.Abs(dir)
		if err != nil {
			return nil, err
		}
	}

	return filepath.Glob(filepath.Join(dir, glob))
}
