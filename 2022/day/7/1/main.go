package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

const maxDirSize = 100_000

type file struct {
	name string
	size int64
}

type dirModel struct {
	dir
	total int64
}

type dir []file

func (dir dir) Size() int64 {
	var v int64
	for _, f := range dir {
		v += f.size
	}
	return v
}

func main() {
	args := os.Args[1:]
	filePath := args[0]

	paths := make(map[string]dirModel)
	currentPath := ""
	currentPathWithFile := ""
	err := reader.ReadFile(filePath, func(l string) error {
		ch := l[0]
		if ch == '$' {
			// command

			command := l[2:4]
			if command == "cd" {

				//fmt.Println("->command:", l)
				//fmt.Println("current_path+before", currentPath)

				if l[5] == '/' {
					currentPath = path.Base("/")
				} else if len(l) == 7 && l[5:6] == ".." {
					if currentPath != "/" {
						currentPath = path.Dir(currentPath)
					}
				} else {
					dir := l[5:]
					currentPath = path.Join(currentPath, dir)
				}

				//fmt.Println("current_path+after", currentPath)

			} else if command == "ls" {
				currentPathWithFile = currentPath
			}
		} else if ch == 'd' {

			//fmt.Println("->directory:", l)

			// directory
			withDir := path.Join(currentPathWithFile, l[4:])
			_, ok := paths[withDir]
			if !ok {
				paths[withDir] = dirModel{
					dir:   nil,
					total: 0,
				}
			}

			//fmt.Println("currentPathWithFile", currentPathWithFile)

		} else if ch >= '0' && ch <= '9' {
			// file

			//fmt.Println("currentPathWithFile->", currentPathWithFile, "file", l)

			s := strings.Split(l, " ")
			nn, err := strconv.Atoi(s[0])
			if err != nil {
				return err
			}
			n := int64(nn)

			m := paths[currentPathWithFile]
			exists := false
			for _, ff := range m.dir {
				if ff.name == s[1] {
					exists = true

					//fmt.Println("exists file", ff.name)

					break
				}
			}
			if !exists {
				m.dir = append(m.dir, file{name: s[1], size: n})
				m.total += n
				paths[currentPathWithFile] = m
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	//return

	//fmt.Println(paths)

	pathsArr := make([]string, 0, len(paths))
	for k := range paths {
		pathsArr = append(pathsArr, k)
	}
	sort.Strings(pathsArr)

	//fmt.Println(pathsArr)

	for i := 0; i < len(pathsArr); i++ {
		v1 := pathsArr[i]
		for j := i + 1; j < len(pathsArr); j++ {
			v2 := pathsArr[j]

			//fmt.Println("v1", v1, "v2", v2)

			if strings.HasPrefix(v2, v1) {
				vv := paths[v1]
				vv.total += paths[v2].total
				paths[v1] = vv

				//fmt.Println("vv", vv)
			}
		}
	}

	fmt.Println(paths)

	var total int64
	for _, v := range paths {
		if v.total < maxDirSize {
			total += v.total
		}
	}

	fmt.Println(total)
}
