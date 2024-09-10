package search

import (
	"file-search/criteria"
	"file-search/fileSystem"
)

type FileSearchService struct {
	Directories  []fileSystem.Directory
	SearchParams criteria.ICriteria
}

func (fs *FileSearchService) Search() []fileSystem.FileSystem {
	i := 0
	var files []fileSystem.FileSystem
	for i < len(fs.Directories) {
		children := fs.Directories[i].GetChildren()
		for _, child := range children {
			if child.IsDirectory() {
				fs.Directories = append(fs.Directories, *child.(*fileSystem.Directory))
			} else if fs.SearchParams.IsSatifiedBy(*child.(*fileSystem.File)) {
				files = append(files, child)
			}
		}
		i++
	}
	return files
}
