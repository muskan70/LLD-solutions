package main

import (
	"file-search/criteria"
	"file-search/file"
	"file-search/search"
	"fmt"
)

func main() {
	rootDir := file.NewFile("root", 10, true)
	utl := file.NewFile("hello.txt", 3, false)
	utl2 := file.NewFile("picture.png", 7, false)
	utl3 := file.NewFile("response.json", 2, false)

	rootDir.AddChild(utl)
	rootDir.AddChild(utl2)
	rootDir.AddChild(utl3)

	spec1 := criteria.NewNameCriteria("hello")
	spec2 := criteria.NewExtensionCriteria("txt")
	spec := criteria.NewAndCriteria([]criteria.ICriteria{spec1, spec2})

	fs := search.FileSearchService{Directories: []*file.File{rootDir}, SearchParams: spec}
	files := fs.Search()

	for i := range files {
		fmt.Println(files[i])
	}

}
