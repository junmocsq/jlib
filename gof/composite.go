package gof

import (
	"fmt"
	"path/filepath"
)

// 组合模式 composite design pattern

type CompositeFiler interface {
	CountNumOfFiles() int
	CountSizeOfFiles() int
	Name() string
}
type CFile struct {
	size int
	name string
}

var _ CompositeFiler = &CFile{}

func newCFile(name string, size int) *CFile {
	return &CFile{
		name: name,
		size: size,
	}
}
func (c *CFile) CountNumOfFiles() int {
	return 1
}

func (c *CFile) CountSizeOfFiles() int {
	return c.size
}

func (c *CFile) Name() string {
	return c.name
}

type CDirectory struct {
	prefix string
	name   string
	files  []*CFile
	dirs   []*CDirectory
}

var _ CompositeFiler = &CDirectory{}

func NewDir(prefix string, name string) *CDirectory {
	return &CDirectory{
		prefix: prefix,
		name:   name,
		files:  nil,
		dirs:   nil,
	}
}

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
	return c.name
}

func (c *CDirectory) Dir() string {
	return filepath.Join(c.prefix, c.name)
}

func (c *CDirectory) AddDirNode(dirname string) *CDirectory {

	for _, v := range c.dirs {
		if v.name == dirname {
			return v
		}
	}
	prefix := filepath.Join(c.prefix, c.name)
	dir := &CDirectory{
		prefix: prefix,
		name:   dirname,
		files:  nil,
		dirs:   nil,
	}
	c.dirs = append(c.dirs, dir)
	return dir
}

func (c *CDirectory) AddFileNode(filename string, size int) bool {
	for _, v := range c.files {
		if v.name == filename {
			return false
		}
	}
	c.files = append(c.files, newCFile(filename, size))
	return true
}

func (c *CDirectory) Print() {
	prefix := filepath.Join(c.prefix, c.name)
	fmt.Println(prefix, c.CountSizeOfFiles(), c.CountNumOfFiles())

	var f func(prefix string, dirs []*CDirectory, files []*CFile)

	f = func(prefix string, dirs []*CDirectory, files []*CFile) {
		for _, file := range files {
			fmt.Println(filepath.Join(prefix, file.name))
		}
		for _, dir := range dirs {
			fmt.Println(filepath.Join(prefix, dir.name), dir.CountSizeOfFiles(), dir.CountNumOfFiles())
			f(filepath.Join(dir.prefix, dir.name), dir.dirs, dir.files)
		}
	}
	f(prefix, c.dirs, c.files)
}
