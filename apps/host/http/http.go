package http

import (
	"cmdb/apps/host"
	"cmdb/apps/oss"
	"github.com/gin-gonic/gin"
	"os"
)

func NewHostHttpHander()*Hander{
	return &Hander{
	}
}

func( h *Hander) Config(){
	 if apps.HostService == nil{
		 os.Exit(10)
	 }
	 h.svc = apps.HostService
}


func (h *Hander) Name() string {
	return host.APPNAME
}

type Hander struct {
	svc host.Service
}


func (h *Hander)Register(r gin.IRouter){
	 r.POST("/hosts" , h.createHost)
}

func init(){
	apps.GinRegister(NewHostHttpHander())

}