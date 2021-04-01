package gof

import (
	"os"
	"path/filepath"
)

// 组合模式 composite design pattern

type CompositeFiler interface {
	CountNumOfFiles() int
	CountSizeOfFiles() int
	Name() string
	Dir() string
}
type CFile struct {
	path string
}

var _ CompositeFiler = &CFile{}

func (c *CFile) CountNumOfFiles() int {
	return 1
}

func (c *CFile) CountSizeOfFiles() int {
	f, e := os.Open(c.path)
	if e != nil {
		return 0
	}
	stats, e := f.Stat()
	if e != nil {
		return 0
	}
	return int(stats.Size())
}

func (c *CFile) Name() string {
	return filepath.Base(c.path)
}

func (c *CFile) Dir() string {
	return filepath.Dir(c.path)
}

type CDirectory struct {
	path  string
	files []*CFile
	dirs  []*CDirectory
}

var _ CompositeFiler = &CDirectory{}

func (c *CDirectory) CountNumOfFiles() int {
	num := 0
	for _, f := range c.files {
		num += f.CountNumOfFiles()
	}

	for _, d := range c.dirs {
		num += d.CountNumOfFiles()
	}
	return num
}

func (c *CDirectory) CountSizeOfFiles() int {
	num := 0
	for _, f := range c.files {
		num += f.CountSizeOfFiles()
	}

	for _, d := range c.dirs {
		num += d.CountSizeOfFiles()
	}
	return num
}

func (c *CDirectory) Name() string {
	return filepath.Base(c.path)
}

func (c *CDirectory) Dir() string {
	return filepath.Dir(c.path)
}

func (c *CDirectory) AddNodes() bool {

}
