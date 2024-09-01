package search

import (
	"file-search/criteria"
	"file-search/file"
)

type FileSearchService struct {
	Directories  []*file.File
	SearchParams criteria.ICriteria
}

func (fs *FileSearchService) Search() []*file.File {
	i := 0
	var files []*file.File
	for i < len(fs.Directories) {
		children := fs.Directories[i].GetChildren()
		for _, child := range children {
			if child.IsDirectory() {
				fs.Directories = append(fs.Directories, child)
			} else if fs.SearchParams.IsSatifiedBy(child) {
				files = append(files, child)
			}
		}
		i++
	}
	return files
}
