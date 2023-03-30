package imp

import (
	"cmdb/apps/host"
	"context"
	"github.com/infraboard/mcube/sqlbuilder"
)




func (i *HostServiceImp) CreateHost(ctx context.Context ,ins *host.Host)(*host.Host , error){

    //i.log.Debug("ccreate log")
	//i.log.Debugf("create host %s " , ins.Name)
	//i.log.With(logger.NewAny("request_id" , "reqo1")).Debug("cereate host with kv")

	//校验数据的合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	ins.InjectDeault()
	if err := i.save(ctx , ins);err!=nil{
		i.l.Debug("ccreate log")
		 return nil,err
	}
	return  ins,nil

}


//查单个数据
func(i *HostServiceImp) DescibeHost(ctx context.Context, request *host.DescribeHostRequest) (*host.Host , error){
	b :=sqlbuilder.NewBuilder(QueryHostSQL)
	b.Where("r.id=?", request.Id)
	querysql , args := b.Build()
	i.l.Debugf("describe sql: %s, args: %v", querysql, args)
	stmt , err := i.db.PrepareContext(ctx,querysql)
	if err!=nil{
		 return nil , err
	}

	defer  stmt.Close()
	ins := host.NewHost()
	//传递切片
	err = stmt.QueryRowContext(ctx , args...).Scan(&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
		&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
		&ins.Account, &ins.PublicIP, &ins.PrivateIP,
		&ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber,)
	if err != nil {
		return nil, err
	}
	return ins, nil

}


func (i *HostServiceImp) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.Host, error) {

	return nil, nil
}

func (i *HostServiceImp)  UpdateHost(ctx context.Context ,request *host.UpdateHostRequest)(*host.Host , error) {
	// 获取已有对象
	return nil, nil
}

func (i *HostServiceImp) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (
	*host.Host, error) {
	return nil, nil
}