package imp

import (
	"cmdb/apps/host"
	apps "cmdb/apps/oss"
	"cmdb/conf"
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)
//接口 实现接口
//var _ host.Service = (*HostServiceImp)(nil)

var impl = &HostServiceImp{}

//调用之前保证 db 已经初始化啦
func NewHostServiceImp()*HostServiceImp{
	 return &HostServiceImp{
		  l:zap.L().Named("Host"),
		  db:conf.Conf().MySQL.GetDB(),
	 }
}

//这块儿保证 config 完成初始化的操做
func(i *HostServiceImp)Config(){
    i.l=zap.L().Named("Host")
	i.db=conf.Conf().MySQL.GetDB()
}

func(i *HostServiceImp)Name() string{
	return host.APPNAME
}

type HostServiceImp struct {
	l  logger.Logger
	db   *sql.DB
}

func init(){
	//对象注册
	 apps.HostService = impl
	 apps.Register(impl)
}
