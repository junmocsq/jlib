package gof

import "testing"

func TestComposite(t *testing.T) {
	dir := NewDir("/root", "work")
	d1 := dir.AddDirNode("www")
	d11 := d1.AddDirNode("gzcp")
	d11.AddFileNode("index.html", 999)
	d11.AddFileNode("index.htm", 123)
	d11.AddDirNode("work")
	d11.AddFileNode("work.txt", 222)

	d12 := d1.AddDirNode("distr")
	d12.AddFileNode("index.html", 12111)
	d12.AddFileNode("distr.work", 12111)
	d12.AddFileNode("distr.txt", 12111)

	dir.AddDirNode(".ssh")
	dir.AddFileNode("nohup.out", 8761)

	dir.Print()
}
