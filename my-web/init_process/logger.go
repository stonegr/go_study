package init_process

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

func DeferCs() (r int) {
	defer func() {
		r++
	}()
	return 0
}

type MyFormatter struct {
}

func (m *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	newLog := fmt.Sprintf("[%s]-[%s]-[%s:%d]-%s\n", timestamp, entry.Level, entry.Caller.File, entry.Caller.Line, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func InitLogger(logdir string) *lumberjack.Logger {
	// f, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	f := &lumberjack.Logger{
		// 日志输出文件路径
		Filename: path.Join(logdir, "output.log"),
		// 日志文件最大 size, 单位是 MB
		MaxSize: 1, // megabytes
		// 最大过期日志保留的个数
		MaxBackups: 3,
		// 保留过期文件的最大时间间隔,单位是天
		MaxAge: 1, //days
		// 是否需要压缩滚动日志, 使用的 gzip 压缩
		Compress: false, // disabled by default
	}
	writers := []io.Writer{
		f,
		// os.Stdout,
	}
	if Myconfig.Debug {
		log.SetOutput(os.Stdout)
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdoutWriter)

	log.SetLevel(log.TraceLevel) // 在测试环境中设置低等级级别，
	//logrus.SetLevel(logrus.InfoLevel)    // 在生产环境中需要考虑性能，关注关键信息，level 设置高一点
	log.SetReportCaller(true) // 调用者文件名与位置
	// log.SetFormatter(&log.TextFormatter{
	// 	ForceQuote:      true,
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
	log.SetFormatter(&MyFormatter{})

	// log.Traceln("trace 日志")
	// log.Debugln("debug 日志")
	// log.Infoln("Info 日志")
	// log.Warnln("warn 日志")
	// log.Errorln("error msg")
	return f
}
