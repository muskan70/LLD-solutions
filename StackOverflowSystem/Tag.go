package main

type Tag struct {
	Id   int
	Name string
}

var tagId = 0
var tagMap = make(map[string]int)
var tagIdMap = make(map[int]string)

func CreateTag(name string) int {
	if tag, ok := tagMap[name]; ok {
		return tag
	}
	tagId++
	tagMap[name] = tagId
	tagIdMap[tagId] = name
	return tagId
}

func GetTagId(name string) int {
	return CreateTag(name)
}

func GetTagName(id int) string {
	if name, ok := tagIdMap[id]; ok {
		return name
	}
	return ""
}
