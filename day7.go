package main

import (
	"fmt"
	"strings"
)

type filesystem struct {
	root    *dir
	current *dir
}

type dir struct {
	name     string
	files    int // total size of files
	parent   *dir
	children []*dir
}

func initFilesystem() *filesystem {
	root := &dir{
		name:     "/",
		files:    0,
		parent:   nil,
		children: []*dir{},
	}
	fs := &filesystem{
		root:    root,
		current: root,
	}
	return fs
}

func (f *filesystem) newDir(name string) *dir {
	d := &dir{
		name:     name,
		files:    0,
		parent:   f.current,
		children: []*dir{},
	}
	return d
}

func (f *filesystem) changeDir(name string) {
	if name == ".." {
		f.current = f.current.parent
	} else if name == "/" {
		f.current = f.root
	} else {
		for _, val := range f.current.children {
			if val.name == name {
				f.current = val
				return
			}
		}
	}
}

func (f *filesystem) insertDir(name string) {
	children := f.current.children
	children = append(children, f.newDir(name))
	f.current.children = children
}

func (f *filesystem) insertFile(size int) {
	f.current.files += size
}

func day7(filename string) {
	data := getLines(filename)
	part1, part2 := makeSpaceForUpdate(data)

	fmt.Println("Day 7")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func makeSpaceForUpdate(data []string) (int, int) {
	fs := initFilesystem()

	for _, line := range data[1:] {
		if strings.HasPrefix(line, "$") {
			command := strings.Split(line, " ")
			if command[1] == "ls" {

				continue
			}
			fs.changeDir(command[2])
			continue
		}
		if strings.HasPrefix(line, "dir") {
			name := strings.Split(line, " ")[1]
			fs.insertDir(name)
			continue
		}
		fileSize := strings.Split(line, " ")[0]
		fs.insertFile(readInt(fileSize))
	}

	part1, part2 := iterateFs(fs.root, dirSize(fs.root))

	return part1, part2
}

func iterateFs(dir *dir, usedSpace int) (int, int) {
	var (
		sum      = 0  // part 1
		toDelete = -1 // part 2

		size      = dirSize(dir)
		total     = 70_000_000
		need      = 30_000_000
		freeSpace = total - usedSpace + size
	)

	if size <= 100_000 {
		sum += size
	}

	if freeSpace >= need {
		toDelete = size
	}

	for _, val := range dir.children {
		valSum, valToDelete := iterateFs(val, usedSpace)

		sum += valSum

		if valToDelete > 0 && valToDelete < toDelete {
			toDelete = valToDelete
		}
	}

	return sum, toDelete
}

func dirSize(dir *dir) int {
	size := dir.files
	for _, d := range dir.children {
		size += dirSize(d)
	}
	return size
}
