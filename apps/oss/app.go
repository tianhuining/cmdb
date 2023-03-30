package apps

import (
	"cmdb/apps/host"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

//ioc 容器

var (
	 HostService host.Service
	 svcs=map[string]ImpService{}
	 ginApps = map[string]GinService{}
)


func Register(svc  ImpService){
	if  _,ok:= svcs[svc.Name()];ok{
		fmt.Println("注册失败 ")
		os.Exit(10)
	}
	svcs[svc.Name()] =svc
	if  v,ok :=svc.(host.Service);ok{
		HostService = v
	}
}

type ImpService interface {
	Config()
	Name() string
}


type GinService interface {
	Name()string
	Register(r gin.IRouter)
	Config()
}


func GinRegister(gins  GinService){

	if  _,ok:= ginApps[gins.Name()];ok{
		fmt.Println("注册失败 ")
		os.Exit(10)
	}
	ginApps[gins.Name()] =gins
}

func InitMpl(){
	for  _, v := range svcs{
		 v.Config()
	}

}


func LoadGinApps()(names []string){
	  for k :=range ginApps{
		  names = append(names  , k)
	  }
	  return names
}

func InitGin(r gin.IRouter){
	fmt.Println(ginApps)
	for  _, v := range ginApps{
		v.Config()
	}
	//加载完配置文件
	//开始register路由
	for _,v := range ginApps{
		v.Register(r)
	}
}