// @author:llychao<lychao_vip@163.com>
// @contributor: Junyi<me@junyi.pw>
// @date:2020-02-18
// @功能:golang m3u8 video Downloader
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/levigross/grequests"
)

const (
	// HEAD_TIMEOUT 请求头超时时间
	HEAD_TIMEOUT = 10 * time.Second
	// PROGRESS_WIDTH 进度条长度
	PROGRESS_WIDTH = 20
	// TS_NAME_TEMPLATE ts视频片段命名规则
	TS_NAME_TEMPLATE = "%05d.ts"
)

var (
	// 命令行参数
	urlFlag = flag.String("u", "", "m3u8下载地址(http(s)://url/xx/xx/index.m3u8)")
	nFlag   = flag.Int("n", 16, "下载线程数(默认16)")
	htFlag  = flag.String("ht", "apiv1", "设置getHost的方式(apiv1: `http(s):// + url.Host + filepath.Dir(url.Path)`; apiv2: `http(s)://+ u.Host`")
	oFlag   = flag.String("o", "movie", "自定义文件名(默认为movie)不带后缀")
	cFlag   = flag.String("c", "", "自定义请求cookie")
	sFlag   = flag.Int("s", 0, "是否允许不安全的请求(默认0)")
	spFlag  = flag.String("sp", "", "文件保存的绝对路径(默认为当前路径,建议默认值)")
	pFlag   = flag.String("p", "", "网络http代理")
	fFlag   = flag.String("f", "", "指定m3u8文件")
	stFlag  = flag.Bool("st", false, "保留下载的ts文件")
	newfile = flag.Bool("nm", false, "是否生成新的m3u8文件,必须在st为true的情况下，否则会删除")

	logger *log.Logger
	ro     = &grequests.RequestOptions{
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36",
		RequestTimeout: HEAD_TIMEOUT,
		Headers: map[string]string{
			"Connection":      "keep-alive",
			"Accept":          "*/*",
			"Accept-Encoding": "*",
			"Accept-Language": "zh-CN,zh;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5",
		},
	}
)

// TsInfo 用于保存 ts 文件的下载地址和文件名
type TsInfo struct {
	Name string
	Url  string
}

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Run()
}

func Run() {
	msgTpl := "[功能]:多线程下载直播流m3u8视屏\n[提醒]:下载失败，请使用 -ht=apiv2 \n[提醒]:下载失败，m3u8 地址可能存在嵌套\n[提醒]:进度条中途下载失败，可重复执行"
	fmt.Println(msgTpl)
	runtime.GOMAXPROCS(runtime.NumCPU())
	now := time.Now()

	// 1、解析命令行参数
	flag.Parse()
	m3u8Url := *urlFlag
	maxGoroutines := *nFlag
	hostType := *htFlag
	movieName := *oFlag
	cookie := *cFlag
	insecure := *sFlag
	savePath := *spFlag
	proxy := *pFlag
	m3u8_file_path := *fFlag
	is_save_ts := *stFlag
	newfile := *newfile

	proxyURL, e := url.Parse(proxy)
	checkErr(e)

	if proxy != "" {
		ro.Proxies = map[string]*url.URL{
			proxyURL.Scheme: proxyURL,
			"https":         proxyURL,
		}
	}

	ro.Headers["Referer"] = getHost(m3u8Url, "apiv2")
	if insecure != 0 {
		ro.InsecureSkipVerify = true
	}
	// http 自定义 cookie
	if cookie != "" {
		ro.Headers["Cookie"] = cookie
	}
	if m3u8Url == "" && m3u8_file_path == "" {
		flag.Usage()
		return
	}
	var download_dir string
	pwd, _ := os.Getwd()
	if savePath != "" {
		pwd, _ = filepath.Abs(savePath)
		fmt.Printf("程序运行目录为: %s\n", pwd)
	}
	// 初始化下载ts的目录，后面所有的ts文件会保存在这里
	download_dir = filepath.Join(pwd, movieName)
	if isExist, _ := pathExists(download_dir); !isExist {
		os.MkdirAll(download_dir, os.ModePerm)
	}

	// 2、解析m3u8
	m3u8Host := getHost(m3u8Url, hostType)
	var m3u8Body string
	if m3u8_file_path != "" {
		m3u8Body = getFromFile(filepath.Join(pwd, m3u8_file_path))
	} else {
		m3u8Body = getM3u8Body(m3u8Url)
	}
	ts_key := getM3u8Key(m3u8Host, m3u8Body)
	if ts_key != "" {
		fmt.Printf("待解密 ts 文件 key : %s \n", ts_key)
	}
	ts_list := getTsList(m3u8Host, m3u8Body)
	if newfile {
		creatNewMF(filepath.Join(download_dir, movieName+".m3u8"), ts_list, get_d(m3u8Body))
	}
	fmt.Println("待下载 ts 文件数量:", len(ts_list))

	// 3、下载ts文件到download_dir
	downloader(ts_list, maxGoroutines, download_dir, ts_key)
	if ok := checkTsDownDir(download_dir); !ok {
		fmt.Printf("\n[Failed] 请检查url地址有效性 \n")
		return
	}

	if !newfile {
		// 4、合并ts切割文件成mp4文件
		mv := mergeTs(download_dir, is_save_ts, movieName, pwd)
		//5、输出下载视频信息
		DrawProgressBar("Merging", float32(1), PROGRESS_WIDTH, mv)
		fmt.Printf("\n[Success] 下载保存路径：%s | 共耗时: %6.2fs\n", mv, time.Now().Sub(now).Seconds())
	} else {
		fmt.Println("")
	}

}

// 获取m3u8地址的host
func getHost(Url, ht string) (host string) {
	u, err := url.Parse(Url)
	checkErr(err)
	switch ht {
	case "apiv1":
		host = u.Scheme + "://" + u.Host + filepath.Dir(u.EscapedPath())
	case "apiv2":
		host = u.Scheme + "://" + u.Host
	}
	return
}

// 获取m3u8地址的内容体
func getM3u8Body(Url string) string {
	r, err := grequests.Get(Url, ro)
	checkErr(err)
	return r.String()
}

// 获取m3u8加密的密钥
func getM3u8Key(host, html string) (key string) {
	lines := strings.Split(html, "\n")
	key = ""
	for _, line := range lines {
		if strings.Contains(line, "#EXT-X-KEY") {
			uri_pos := strings.Index(line, "URI")
			quotation_mark_pos := strings.LastIndex(line, "\"")
			key_url := strings.Split(line[uri_pos:quotation_mark_pos], "\"")[1]
			if !strings.Contains(line, "http") {
				key_url, _ = url.JoinPath(host, key_url)
				// key_url = fmt.Sprintf("%s/%s", host, key_url)
			}
			fmt.Printf("提取到的key_url: %s\n", key_url)
			res, err := grequests.Get(key_url, ro)
			checkErr(err)
			if res.StatusCode == 200 {
				key = res.String()
			}
		}
	}
	return
}

func get_d(html string) (key string) {
	lines := strings.Split(html, "\n")
	key = ""
	for _, line := range lines {
		if strings.Contains(line, "#EXTINF") {
			line_cap := strings.Split(line, ":")
			if len(line_cap) > 1 {
				key = line_cap[len(line_cap)-1]
				break
			}
		}
	}
	return
}

func getTsList(host, body string) (tsList []TsInfo) {
	lines := strings.Split(body, "\n")
	index := 0
	var ts TsInfo
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") && line != "" {
			//有可能出现的二级嵌套格式的m3u8,请自行转换！
			index++
			if strings.HasPrefix(line, "http") {
				ts = TsInfo{
					Name: fmt.Sprintf(TS_NAME_TEMPLATE, index),
					Url:  line,
				}
				tsList = append(tsList, ts)
			} else {
				ts = TsInfo{
					Name: fmt.Sprintf(TS_NAME_TEMPLATE, index),
					Url:  fmt.Sprintf("%s/%s", host, line),
				}
				tsList = append(tsList, ts)
			}
		}
	}
	return
}

func getFromFile(p string) string {
	data, _ := ioutil.ReadFile(p)
	return string(data)
}

// 生成新的m3u8文件
func creatNewMF(name string, p []TsInfo, d string) {
	content := []byte(`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-TARGETDURATION:6
#EXT-X-PLAYLIST-TYPE:VOD
#EXT-X-MEDIA-SEQUENCE:0
`)
	for _, t := range p {
		content = append(content, []byte(fmt.Sprintf("%s\n%s\n", fmt.Sprintf("#EXTINF:%s", d), t.Name))...)
	}
	content = append(content, []byte("#EXT-X-ENDLIST")...)
	err := ioutil.WriteFile(name, content, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write new m3u8 successful")
}

// 下载ts文件
// @modify: 2020-08-13 修复ts格式SyncByte合并不能播放问题
func downloadTsFile(ts TsInfo, download_dir, key string, retries int) {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("网络不稳定，正在进行断点持续下载")
			downloadTsFile(ts, download_dir, key, retries-1)
		}
	}()
	curr_path_file := fmt.Sprintf("%s/%s", download_dir, ts.Name)
	if isExist, _ := pathExists(curr_path_file); isExist {
		//logger.Println("[warn] File: " + ts.Name + "already exist")
		return
	}
	res, err := grequests.Get(ts.Url, ro)
	if err != nil || !res.Ok {
		if retries > 0 {
			downloadTsFile(ts, download_dir, key, retries-1)
			return
		} else {
			//logger.Printf("[warn] File :%s", ts.Url)
			return
		}
	}
	// 校验长度是否合法
	var origData []byte
	origData = res.Bytes()
	contentLen := 0
	contentLenStr := res.Header.Get("Content-Length")
	if contentLenStr != "" {
		contentLen, _ = strconv.Atoi(contentLenStr)
	}
	if len(origData) == 0 || (contentLen > 0 && len(origData) < contentLen) || res.Error != nil {
		//logger.Println("[warn] File: " + ts.Name + "res origData invalid or err：", res.Error)
		downloadTsFile(ts, download_dir, key, retries-1)
		return
	}
	// 解密出视频 ts 源文件
	if key != "" {
		//解密 ts 文件，算法：aes 128 cbc pack5
		origData, err = AesDecrypt(origData, []byte(key))
		if err != nil {
			downloadTsFile(ts, download_dir, key, retries-1)
			return
		}
	}
	// https://en.wikipedia.org/wiki/MPEG_transport_stream
	// Some TS files do not start with SyncByte 0x47, they can not be played after merging,
	// Need to remove the bytes before the SyncByte 0x47(71).
	syncByte := uint8(71) //0x47
	bLen := len(origData)
	for j := 0; j < bLen; j++ {
		if origData[j] == syncByte {
			origData = origData[j:]
			break
		}
	}
	ioutil.WriteFile(curr_path_file, origData, 0666)
}

// downloader m3u8 下载器
func downloader(tsList []TsInfo, maxGoroutines int, downloadDir string, key string) {
	retry := 5 //单个 ts 下载重试次数
	var wg sync.WaitGroup
	limiter := make(chan struct{}, maxGoroutines) //chan struct 内存占用 0 bool 占用 1
	tsLen := len(tsList)
	downloadCount := 0
	for _, ts := range tsList {
		wg.Add(1)
		limiter <- struct{}{}
		go func(ts TsInfo, downloadDir, key string, retryies int) {
			defer func() {
				wg.Done()
				<-limiter
			}()
			downloadTsFile(ts, downloadDir, key, retryies)
			downloadCount++
			DrawProgressBar("Downloading", float32(downloadCount)/float32(tsLen), PROGRESS_WIDTH, ts.Name)
			return
		}(ts, downloadDir, key, retry)
	}
	wg.Wait()
}

func checkTsDownDir(dir string) bool {
	if isExist, _ := pathExists(filepath.Join(dir, fmt.Sprintf(TS_NAME_TEMPLATE, 0))); !isExist {
		return true
	}
	return false
}

// 合并ts文件
func mergeTs(downloadDir string, is_save_ts bool, movieName string, pwd string) string {
	mvName := downloadDir + ".mp4"
	// outMv, _ := os.Create(mvName)
	// defer outMv.Close()
	// writer := bufio.NewWriter(outMv)
	// err := filepath.Walk(downloadDir, func(path string, f os.FileInfo, err error) error {
	// 	if f == nil {
	// 		return err
	// 	}
	// 	if f.IsDir() || filepath.Ext(path) != ".ts" {
	// 		return nil
	// 	}
	// 	bytes, _ := ioutil.ReadFile(path)
	// 	_, err = writer.Write(bytes)
	// 	return err
	// })
	// checkErr(err)
	// _ = writer.Flush()
	execUnixShell(fmt.Sprintf("cd %s && rm -rf hb.txt %s.mp4 && ls -l %s | tail -n +2 | awk '{print $NF}' | sed \"s/^/file %s\\//g\" > hb.txt", pwd, movieName, movieName, movieName))
	execUnixShell(fmt.Sprintf("cd %s && ffmpeg -f concat -safe 0 -i hb.txt -async 1000 -c copy %s.mp4", pwd, movieName))
	execUnixShell(fmt.Sprintf("cd %s && rm -rf hb.txt", pwd))
	if !is_save_ts {
		os.RemoveAll(downloadDir)
	}
	return mvName
}

// 进度条
func DrawProgressBar(prefix string, proportion float32, width int, suffix ...string) {
	pos := int(proportion * float32(width))
	s := fmt.Sprintf("[%s] %s%*s %6.2f%% \t%s",
		prefix, strings.Repeat("■", pos), width-pos, "", proportion*100, strings.Join(suffix, ""))
	fmt.Print("\r" + s)
}

// ============================== shell相关 ==============================
// 判断文件是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 执行 shell
func execShell(s string) {
	cmd := exec.Command("bash", "-c", "rm -rf hb.txt && ls -l movie | tail -n +2 | awk '{print $NF}' | sed 's/^/file movie\\//g' > hb.txt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s", out.String())
}

func execUnixShell(s string) {
	cmd := exec.Command("bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", out.String())
}

func execWinShell(s string) error {
	cmd := exec.Command("cmd", "/C", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("%s", out.String())
	return nil
}

// windows 合并文件
func win_merge_file(path string) {
	pwd, _ := os.Getwd()
	os.Chdir(path)
	execWinShell("copy /b *.ts merge.tmp")
	execWinShell("del /Q *.ts")
	os.Rename("merge.tmp", "merge.mp4")
	os.Chdir(pwd)
}

// unix 合并文件
func unix_merge_file(path string) {
	pwd, _ := os.Getwd()
	os.Chdir(path)
	//cmd := `ls  *.ts |sort -t "\." -k 1 -n |awk '{print $0}' |xargs -n 1 -I {} bash -c "cat {} >> new.tmp"`
	cmd := `cat *.ts >> merge.tmp`
	execUnixShell(cmd)
	execUnixShell("rm -rf *.ts")
	os.Rename("merge.tmp", "merge.mp4")
	os.Chdir(pwd)
}

// ============================== 加解密相关 ==============================

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte, ivs ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	var iv []byte
	if len(ivs) == 0 {
		iv = key
	} else {
		iv = ivs[0]
	}
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte, ivs ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	var iv []byte
	if len(ivs) == 0 {
		iv = key
	} else {
		iv = ivs[0]
	}
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func checkErr(e error) {
	if e != nil {
		logger.Panic(e)
	}
}
