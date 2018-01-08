/*
sort documenataion: https://golang.org/pkg/sort/
ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir
os.FileInfo: https://golang.org/pkg/os/#FileInfo
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

/*
ioutil.ReadDir returns []os.FileInfo
os.fileInfo do not implement sort.Interface
so we have to implement it by our own
*/
type BySize []os.FileInfo

/* sort requires three method
len -> knows how many elements
swap -> swap element a with element b
less -> if one element is less than other
*/
func (a BySize) Len() int      { return len(a) }
func (a BySize) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

/* we need to sort by size. FileInfo is defined as
	type FileInfo interface {
			Name() string       // base name of the file
			Size() int64        // length in bytes for regular files; system-dependent for others
			Mode() FileMode     // file mode bits
			ModTime() time.Time // modification time
			IsDir() bool        // abbreviation for Mode().IsDir()
			Sys() interface{}   // underlying data source (can return nil)
	}
so when we build the Less method, we call element.Size()
*/
func (a BySize) Less(i, j int) bool { return a[i].Size() < a[j].Size() }

func main() {
	files, err := ioutil.ReadDir("./")
	isdir := "file"
	if err != nil {
		fmt.Println(err)
	}
	sort.Sort(BySize(files))
	for _, f := range files {
		if f.IsDir() == true {
			isdir = "dir "
		} else {
			isdir = "file"
		}
		fmt.Printf("%7d %s %s %s\n", f.Size(), f.Mode(), isdir, f.Name())
	}
}
