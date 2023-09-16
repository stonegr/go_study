package server_back

import (
	"my-web/init_process"
	"os"
	"path"
	"sort"

	log "github.com/sirupsen/logrus"
)

// 验证文件或者文件夹是否存在
func PathExists(file_path string) (exsits bool, err error) {
	_, _err := os.Stat(file_path)
	if _err != nil {
		if os.IsExist(_err) {
			exsits = true
			err = nil
		} else {
			exsits = false
			err = _err
		}
	} else {
		exsits = true
	}
	return
}

// 文件处理
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
func SortFile(path string, max_file int) (files ByModTime) {
	f, err := os.Open(path)
	if err != nil {
		log.Error(err)
	}
	fis, err := f.Readdir(-1)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()

	files = make(ByModTime, len(fis))
	j := 0
	for _, v := range fis {
		files[j] = v
		j++
	}
	files = files[:j]

	sort.Sort(ByModTime(files))
	return
}

// 返回当下时间的文件，并删除大于 2 个的文件，删除最早的，如果目录下没有文件，就自动创建
func DealWithFiles(dir_path string, max_file int) (deleted_files []os.FileInfo) {
	// timestamp := time.Now().Format("20060102.150405")
	files := SortFile(dir_path, max_file)
	if len(files) > init_process.Myconfig.MaxFile {
		for _, v := range files[:len(files)-init_process.Myconfig.MaxFile] {
			err := os.Remove(path.Join(dir_path, v.Name()))
			if err != nil {
				log.Error(err)
			}
			deleted_files = append(deleted_files, v)
		}
	}
	return
}
