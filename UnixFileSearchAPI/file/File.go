package file

import "strings"

type File struct {
	Name        string
	SizeInBytes int
	IsDir       bool
	Extension   string
	Content     string
	Children    []*File
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) GetSize() int {
	return f.SizeInBytes
}

func (f *File) IsDirectory() bool {
	return f.IsDir
}

func NewFile(name string, bytes int, isDir bool) *File {
	fd := strings.Split(name, ".")
	ext := ""
	if len(fd) > 1 {
		ext = fd[1]
	}
	return &File{
		Name: fd[0], SizeInBytes: bytes, IsDir: isDir, Extension: ext,
	}
}

func (f *File) AddContent(c string) {
	f.Content = c
}

func (f *File) GetExtension() string {
	return f.Extension
}

func (f *File) AddChild(fs *File) {
	f.Children = append(f.Children, fs)
}

func (f *File) GetChildren() []*File {
	return f.Children
}
