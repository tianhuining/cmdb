package protocol

import (
	apps "cmdb/apps/oss"
	"cmdb/conf"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"time"
)

type HttpServerce struct {
	server *http.Server
	l  logger.Logger
	r gin.IRouter
}

func NewHttpService()*HttpServerce{
	r:=gin.Default()
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.Conf().App.HttpAddr(),
		Handler:           r,
	}

	return &HttpServerce{
		server: server,
		l:zap.L().Named("http server"),
		r:r,
	}
}

func(s  *HttpServerce)Start()error{
	apps.InitGin(s.r)
	//加载apps 日志
	apps := apps.LoadGinApps()
	s.l.Infof( "load gin apps %s ",apps)
	//监听端口
	if err :=s.server.ListenAndServe();err!=nil{
		if err == http.ErrServerClosed{
			 //正常关闭
			 s.l.Infof("close  %s" , err)
			 return  nil
		}
		return err
	}
	return nil
}

func(s  *HttpServerce)Stop(){
	//优雅关闭

	ctx,cancle := context.WithTimeout(context.Background(), 30*time.Second)
	defer  cancle()

	if err := s.server.Shutdown(ctx);err!=nil{
		 s.l.Warnf("shut down http server  error,%s" , err)
	}

}
