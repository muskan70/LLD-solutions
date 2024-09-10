package fileSystem

import "strings"

type FileSystem interface {
	GetName() string
	IsDirectory() bool
}

type File struct {
	Name        string
	SizeInBytes int
	Extension   string
	Content     string
}

func NewFile(name string, bytes int) *File {
	fd := strings.Split(name, ".")
	ext := ""
	if len(fd) > 1 {
		ext = fd[1]
	}
	return &File{
		Name: fd[0], SizeInBytes: bytes, Extension: ext,
	}
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) GetSize() int {
	return f.SizeInBytes
}

func (f *File) IsDirectory() bool {
	return false
}

func (f *File) AddContent(c string) {
	f.Content = c
}

func (f *File) GetExtension() string {
	return f.Extension
}

type Directory struct {
	Name     string
	Children []FileSystem
}

func NewDirectory(name string, bytes int) *Directory {
	return &Directory{
		Name:     name,
		Children: []FileSystem{},
	}
}

func (d *Directory) GetName() string {
	return d.Name
}

func (d *Directory) IsDirectory() bool {
	return true
}

func (d *Directory) AddChild(fs FileSystem) {
	d.Children = append(d.Children, fs)
}

func (d *Directory) GetChildren() []FileSystem {
	return d.Children
}
