package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里写日志相关代码

type FileLogger struct {
	level       LogLevel
	filePath    string //日志文件保存路径
	fileName    string //日志文件保存名称
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 //最大文件大小
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	lineNo    int
}

//NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, 50000),
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}

	return fl
}

//initFile 初始化文件
func (f *FileLogger) initFile() error {
	fullfileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullfileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullfileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	//日志文件打开成功
	f.fileObj = fileObj
	f.errFileObj = errFileObj

	//开启后台goroutine往文件写日志，要访问文件句柄多个会出问题
	// for i := 0; i < 5; i++ {
	go f.writeLogBackground()
	// }
	return nil
}

// Close 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// enable 判断是否需要记录该日志
func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.level
}

// checkSize 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

//splitFile 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogStr := fmt.Sprintf("%s.bak%s", logName, nowStr)
	newLogName := path.Join(f.filePath, newLogStr)

	file.Close()
	//备份现有文件

	os.Rename(logName, newLogName)
	//打开新文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return nil, err
	}
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTemp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTemp.timestamp, getLogString(logTemp.level), logTemp.fileName, logTemp.funcName, logTemp.lineNo, logTemp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTemp.level >= ERROR {
				//ERROR级别再记录到.err文件中
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			//取不到日志休息500ms
			time.Sleep(time.Millisecond * 500)
		}

	}
}

//log 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		//先把日志发送到通道中
		logTemp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			lineNo:    lineNo,
		}
		select {
		case f.logChan <- logTemp:
		default:
			//通道满了就丢掉，保证不阻塞
		}
	}
}

// Debug 输出Debug
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info 输出Info
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning 输出Warning
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error 输出Error
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal 输出Fatal
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
