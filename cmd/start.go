package cmd

import (

	//_"cmdb/alls"
	//	"cmdb/conf"
	_ "cmdb/alls"
	apps "cmdb/apps/oss"
	"cmdb/conf"
	"cmdb/protocol"
	"fmt"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端API",
	Long:  "启动 demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadTomlToConfig("./etc/config.toml")
		if err != nil {
			return err
		}
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		//return nil
		apps.InitMpl()
		//gin.SetMode(gin.ReleaseMode)
		//g := gin.Default()
		//apps.InitGin(g)
		//g.Run(conf.Conf().App.HttpAddr())
		svc := NewManager()
		ch := make(chan os.Signal , 1)
		signal.Notify(ch , syscall.SIGTERM ,syscall.SIGQUIT ,syscall.SIGHUP , syscall.SIGINT )
        go svc.WaitStop(ch)
		return  svc.Start()
	},
}

func NewManager()*Manager{
	 return &Manager{
		 http: protocol.NewHttpService(),
		 l:zap.L().Named("cli "),


	 }
}

func (m *Manager)Start()error{
	return  m.http.Start()
}


//处理来自外部的中断信号
func (m *Manager)WaitStop(ch <-chan os.Signal){
	 for v:=range ch{
		 switch v {
		 default:
			 m.l.Infof("recieve %v" , v)
			 m.http.Stop()
		 }

	}

}



// 问题:
// 1. http API, Grpc API 需要启动, 消息总线也需要监听, 比如负责注册于配置,  这些模块都是独立
//    都需要在程序启动时，进行启动, 都写在start start膨胀到不易维护
// 2. 服务的优雅关闭怎么办? 外部都会发送一个Terminal(中断)信号给程序, 程序时需要处理这个信号
//    需要实现程序优雅关闭的逻辑的处理: 由先后顺序 (从外到内完成资源的释放逻辑处理)
//    1. api 层的关闭 (HTTP, GRPC)
//    2. 消息总线关闭
//    3. 关闭数据库链接
//    4. 如果 使用了注册中心, 完成注册中心的注销操作
//    5. 退出完毕

// 还没有初始化Logger实例
// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 更加Config里面的日志配置，来配置全局Logger对象
	lc := conf.Conf().Log

	// 解析日志Level配置
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	// 使用默认配置初始化Logger的全局配置
	zapConfig := zap.DefaultConfig()

	// 配置日志的Level基本
	zapConfig.Level = level

	// 程序没启动一次, 不必都生成一个新日志文件
	zapConfig.Files.RotateOnStartup = false

	// 配置日志的输出方式
	switch lc.To {
	case conf.ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 并没在把日志输入输出到文件
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}

	// 配置日志的输出格式:
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}

	// 把配置应用到全局Logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

type Manager struct {
	 http *protocol.HttpServerce
	 l    logger.Logger
}


func init() {
	//StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/config.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}