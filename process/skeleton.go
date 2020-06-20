package process

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var _app = new(App)

type Module interface {
	Init()
	Run()
	Close(chan struct{})
}

type App struct {
	queue   chan func()
	modules []Module
	close   []chan struct{}
	stop    chan struct{}
	state   int
}

func (a *App) Init() {
	for _, v := range a.modules {
		v.Init()
	}
}

func (a *App) Send() {

}

func (a *App) Run() {
	// 主线程
	go a.run()
	// 模块线程
	for _, v := range a.modules {
		go v.Run()
	}
}

func (a *App) run() {
	for do := range a.queue {
		safeDone(do)
	}
}

func safeDone(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("do error", err)
		}
	}()
	f()
}

func (a *App) Close() {
	for k, v := range a.modules {
		go v.Close(a.close[k])
	}
	for _, v := range a.close {
		<-v
	}
	// 关闭主线程
	a.stop <- struct{}{}
}

func Start() {
	app := new(App)
	app.Init()
	app.Run()
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign
	app.Close()
}

func RegisterModule(m Module) {
	_app.modules = append(_app.modules, m)
	_app.close = append(_app.close, make(chan struct{}, 1))
}
