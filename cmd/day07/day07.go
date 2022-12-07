package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

type entry interface {
	getName() string
}

type file struct {
	name   string
	size   int
	parent *directory
}

func (f *file) getName() string {
	return f.name
}

func newFile(name string, size int, parent *directory) *file {
	return &file{name: name, size: size, parent: parent}
}

type directory struct {
	name       string
	parent     *directory
	entries    []entry
	cachedSize int
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{name: name, parent: parent}
}

func (d *directory) getName() string {
	return d.name
}

func (d *directory) addEntry(e entry) {
	/* when sorting
	d.entries = util.InsertSortedF(d.entries, e, func(left, right entry) bool {
		return left.getName() < right.getName()
	})
	*/
	d.entries = append(d.entries, e)
	d.cachedSize = -1
}

func (d *directory) addFile(name string, size int) *file {
	f := newFile(name, size, d)
	d.addEntry(f)
	return f
}

func (d *directory) getEntry(name string) entry {
	for _, e := range d.entries {
		if e.getName() == name {
			return e
		}
	}

	return nil
}

func (d *directory) getDirectory(name string) *directory {
	e := d.getEntry(name)
	if d, ok := e.(*directory); ok {
		return d
	}

	return nil
}

func (d *directory) addDirectory(name string) *directory {
	subDir := newDirectory(name, d)
	d.addEntry(subDir)
	return subDir
}

func (d *directory) getOrAddDirectory(name string) *directory {
	subDir := d.getDirectory(name)
	if subDir != nil {
		return subDir
	}

	return d.addDirectory(name)
}

func (d *directory) size() int {
	if d.cachedSize != -1 {
		return d.cachedSize
	}

	s := 0
	for _, e := range d.entries {
		switch v := e.(type) {
		case *directory:
			s += v.size()
		case *file:
			s += v.size
		default:
			panic("invalid type")
		}
	}

	d.cachedSize = s
	return s
}

func (d *directory) forEach(callback func(*directory)) {
	callback(d)
	for _, e := range d.entries {
		switch v := e.(type) {
		case *directory:
			v.forEach(callback)
		default:
		}
	}
}

func (d *directory) String() string {
	var sb strings.Builder
	addIndent := func(indent int) {
		for i := 0; i < indent; i++ {
			sb.WriteRune(' ')
		}
	}

	var core func(*directory, int)
	core = func(current *directory, indent int) {
		addIndent(indent)
		fmt.Fprintf(&sb, "- %s (dir)\n", current.name)

		indent += 2
		for _, e := range current.entries {
			switch v := e.(type) {
			case *directory:
				core(v, indent)
			case *file:
				addIndent(indent)
				fmt.Fprintf(&sb, "- %s (file, size=%d)\n", v.name, v.size)
			default:
				panic("invalid type")
			}
		}
	}

	core(d, 0)
	return sb.String()
}

func parseInput(f embed.FS, name string) (*directory, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return nil, err
	}

	root := newDirectory("/", nil)
	if lines[0] != "$ cd /" {
		return nil, errors.New("bad input")
	}

	lines = lines[1:]
	current := root
	index := 0

	handleChangeDir := func(name string) {
		switch name {
		case "/":
			current = root
		case "..":
			current = current.parent
		default:
			current = current.getOrAddDirectory(name)
		}
	}

	handleList := func() error {
		for index < len(lines) {
			line := lines[index]
			if util.StartsWith(line, '$') {
				break
			}

			parts := strings.Split(line, " ")
			if len(parts) != 2 {
				return errors.New("bad entry")
			}

			if parts[0] == "dir" {
				_ = current.getOrAddDirectory(parts[1])
			} else {
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					return fmt.Errorf("can't parse size: %w", err)
				}

				_ = current.addFile(parts[1], size)
			}

			index++
		}

		return nil
	}

	for index < len(lines) {
		line := lines[index]
		parts := strings.Split(line, " ")
		if parts[0] != "$" {
			return nil, errors.New("expected command")
		}

		index++
		switch parts[1] {
		case "cd":
			handleChangeDir(parts[2])
		case "ls":
			err := handleList()
			if err != nil {
				return nil, err
			}
		}
	}

	return root, nil
}

func part1() {
	d, err := parseInput(f, "input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(d.String())

	sum := 0
	d.forEach(func(current *directory) {
		size := current.size()
		if size < 100000 {
			sum += size
		}
	})

	fmt.Printf("Sum is %d\n", sum)
}

func findSmallestToFree(d *directory) *directory {
	total := 70000000
	free := total - d.size()
	require := 30000000
	smallestSize := total
	var smallest *directory

	d.forEach(func(current *directory) {
		size := current.size()
		if free+size >= require && size < smallestSize {
			smallestSize = size
			smallest = current
		}
	})

	return smallest
}

func part2() {
	d, err := parseInput(f, "input.txt")
	if err != nil {
		panic(err)
	}

	found := findSmallestToFree(d)
	fmt.Printf("%s %d\n", found.name, found.size())
}

func main() {
	p1 := flag.Bool("part1", false, "run part1")
	if *p1 {
		part1()
	} else {
		part2()
	}
}
