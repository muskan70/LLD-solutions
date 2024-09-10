package main

import (
	"file-search/criteria"
	"file-search/fileSystem"
	"file-search/search"
	"fmt"
)

func main() {
	rootDir := fileSystem.NewDirectory("root", 10)
	utl := fileSystem.NewFile("hello.txt", 3)
	utl2 := fileSystem.NewFile("picture.png", 7)
	utl3 := fileSystem.NewFile("response.json", 2)

	rootDir.AddChild(utl)
	rootDir.AddChild(utl2)
	rootDir.AddChild(utl3)

	spec1 := criteria.NewNameCriteria("hello")
	spec2 := criteria.NewExtensionCriteria("txt")
	spec := criteria.NewAndCriteria([]criteria.ICriteria{spec1, spec2})

	fs := search.FileSearchService{Directories: []fileSystem.FileSystem{rootDir}, SearchParams: spec}
	files := fs.Search()

	for i := range files {
		fmt.Println(files[i])
	}

}
