package main

import (
	_ "cmdb/alls"
	"cmdb/cmd"
	"fmt"
)


func main(){

	//
	//err := conf.LoadTomlToConfig("./etc/config.toml")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if err := (); err != nil {
	//	return err
	//}

	//return nil
	//apps.InitMpl()
	//gin.SetMode(gin.ReleaseMode)
	//g := gin.Default()
	//apps.InitGin(g)
	//g.Run(conf.Conf().App.HttpAddr())
	//svc := cmd.NewManager()
	//ch := make(chan os.Signal , 1)
	//signal.Notify(ch , syscall.SIGTERM ,syscall.SIGQUIT ,syscall.SIGHUP , syscall.SIGINT )
	//go svc.WaitStop(ch)
	//err =svc.Start()
	//fmt.Println(err)




	//conf.LoadTomlToConfig("./etc/config.toml")
	//apps.InitMpl()
	//gin.SetMode(gin.ReleaseMode)
	//g := gin.Default()
	//apps.InitGin(g)
	//g.Run(conf.Conf().App.HttpAddr())


	if  err :=cmd.RootCmd.Execute();err!=nil{
		fmt.Println(err)
	}


	//err := conf.LoadTomlToConfig("./etc/config.toml")
	//if err !=nil{
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//zap.DevelopmentSetup()
	//service :=imp.NewHostServiceImp()
	//api:= http.NewHostHttpHander(service)
	//gin.SetMode(gin.DebugMode)
	//g := gin.Default()
	//api.Register(g)
	//fmt.Println(conf.Conf().App.HttpAddr())
	//
	//
	// g.Run(conf.Conf().App.HttpAddr())
}
