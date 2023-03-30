package imp_test

import (
	"cmdb/apps/host"
	"cmdb/apps/host/imp"
	"cmdb/conf"
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)
//
//func TestCreateHost(t *testing.T)  {
//	ins := host.NewHost()
//	ins.Name = "test"
//	service.CreateHost(context.Background() , ins)
//}

var  (
	 service host.Service
)

func TestCreate(t *testing.T) {

	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-01"
	ins.Name = "test"
	ins.Region = "cn-hangzhou"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func init(){

	err := conf.LoadTomlToConfig("../../../etc/config.toml")
	if err !=nil{
		 fmt.Println(err)
		 os.Exit(1)
	}
	zap.DevelopmentSetup()
     imp.NewHostServiceImp()
}
