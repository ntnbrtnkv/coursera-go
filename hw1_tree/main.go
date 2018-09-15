package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

type Symbol string

const (
	space       Symbol = "\t"
	curDir      Symbol = "───"
	element     Symbol = "├───"
	element2    Symbol = "│"
	lastElement Symbol = "└───"
)

type Depth uint

type Printable interface {
	printf(io.Writer, string)
}

type Node struct {
	Name   string
	IsFile bool
	Size   string
	Path   string
	Childs []Node
}

func (dir *Node) printf(out io.Writer, prefix string, depth uint) {
	if depth != 0 {
		prefix += string(space)
	}

	depth += 1

	if len(dir.Childs) > 1 {
		for _, child := range dir.Childs[:len(dir.Childs)-1] {
			str := prefix + string(element) + child.Name
			if child.Size != "" {
				str += fmt.Sprintf(" (%s)", child.Size)
			}
			fmt.Fprintf(out, "%s\n", str)
			child.printf(out, prefix+string(element2), depth)
		}
	}

	if len(dir.Childs) > 0 {
		child := dir.Childs[len(dir.Childs)-1]
		str := prefix + string(lastElement) + child.Name
		if child.Size != "" {
			str += fmt.Sprintf(" (%s)", child.Size)
		}
		fmt.Fprintf(out, "%s\n", str)

		child.printf(out, prefix, depth)
	}

}

func (dir *Node) buildTree(includeFiles bool) error {
	path := dir.Path

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		dir.Name = fileInfo.Name()

		files, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}

		for _, f := range files {
			name := f.Name()

			nextDir := Node{
				Path: path + string(os.PathSeparator) + name,
			}

			nextFileInfo, err := os.Stat(nextDir.Path)
			if err != nil {
				return err
			}

			if !includeFiles && !nextFileInfo.IsDir() {
				continue
			}

			err = nextDir.buildTree(includeFiles)
			if err != nil {
				return err
			}

			dir.Childs = append(dir.Childs, nextDir)
		}

		sort.Slice(dir.Childs, func(i, j int) bool {
			return dir.Childs[i].Name < dir.Childs[j].Name
		})
	} else {
		dir.Name = fileInfo.Name()
		size := fileInfo.Size()
		if size > 0 {
			dir.Size = fmt.Sprintf("%db", size)
		} else {
			dir.Size = "empty"
		}
	}

	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	root := Node{
		Path: path,
	}

	err := root.buildTree(printFiles)
	if err != nil {
		return err
	}

	root.printf(out, "", 0)

	return nil
}
