package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
)

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().Before(fis[j].ModTime())
}

// 根目录下的文件按时间大小排序，从远到近
func SortFile(path string) (files ByModTime) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fis, err := f.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	files = make(ByModTime, len(fis)+10)
	j := 0
	for _, v := range fis {
		files[j] = v
		j++
	}
	files = files[:j]

	sort.Sort(ByModTime(files))
	// for _, fi := range files {
	// 	fmt.Println(fi.Name())
	// }
	return
}

// 返回当下时间的文件，并删除大于 5 个的文件，删除最早的，如果目录下没有文件，就自动创建
func DealWithFiles(dirpath string, maxfile int) (deleted_files []os.FileInfo) {
	// timestamp := time.Now().Format("20060102.150405")
	deleted_files = []os.FileInfo{}
	files := SortFile(dirpath)
	if len(files) > maxfile {
		for k, v := range files[:len(files)-maxfile] {
			err := os.Remove(path.Join(dirpath, files[k].Name()))
			if err != nil {
				log.Fatal(err)
			}
			deleted_files = append(deleted_files, v)
		}
	}
	return deleted_files
}

func main() {
	path := "server_back/bwg"
	max_file := 2

	deleted_files := DealWithFiles(path, max_file)
	for _, v := range deleted_files {
		fmt.Println(v.Name())
	}
}
