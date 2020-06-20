package main

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	f7()
}

func f1() {
	file, err := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(file)
	log.Info("hello")
}

func f2() {
	pathMap := lfshook.PathMap{
		log.InfoLevel:  "./info.log",
		log.ErrorLevel: "./error.log",
	}

	Log := log.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&log.JSONFormatter{},
	))
	Log.Info("info")
	Log.Error("error")
}

func f3() {
	Log := log.New()
	path := "./go.log"

	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	Log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			log.InfoLevel:  writer,
			log.ErrorLevel: writer,
		},
		&log.JSONFormatter{},
	))

	Log.Info("info")
	Log.Error("error")
}

func f4() {
	path := "./test.log"
	rl, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),            // 最新日志的快捷方式
		rotatelogs.WithRotationTime(time.Minute), // 日志分割间隔
		rotatelogs.WithMaxAge(time.Minute*2),     // 日志文件最大保存时间
	)
	log.SetOutput(rl)
	for {
		time.Sleep(time.Second * 5)
		log.Println("-->", time.Now())
	}
}

func f5() {
	rl, err := rotatelogs.New(
		"./test.log.%Y%m%d%H%M",
		rotatelogs.WithRotationTime(time.Minute),
		rotatelogs.WithHandler(rotatelogs.HandlerFunc(func(e rotatelogs.Event) {
			if e.Type() != rotatelogs.FileRotatedEventType {
				return
			}
			fmt.Println("rotate...")
		})),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(rl)
	for {
		time.Sleep(time.Second * 5)
		log.Println("-->", time.Now())
	}
}

func f6() {
	rl, err := rotatelogs.New(
		"./test.log.%Y%m%d%H%M",
		rotatelogs.WithRotationTime(time.Minute),
		rotatelogs.ForceNewFile(), // 如果有同名文件存在，必须创建新文件；没有设置次配置，如果有同名文件存在，则往已经存在的文件中写入日志
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(rl)
	for {
		time.Sleep(time.Second * 5)
		log.Println("-->", time.Now())
	}
}

func f7() {
	rl, err := rotatelogs.New(
		"./test.log.%Y%m%d%H%M",
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.AfterFunc(time.Second*10, func() {
		rl.Rotate()
	})
	log.SetOutput(rl)
	for {
		time.Sleep(time.Second * 5)
		log.Println("-->", time.Now())
	}
}
