package conf

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}




func NewDefaultConfig()*Config{
	 return &Config{
		  App: NewDefaultApp(),
		  Log: NewDefaultLog(),
		  MySQL: NewDefaultMysql(),
	 }
}



func NewDefaultApp()*App{
	 return &App{
		 Name:"demo",
		 Host: "127.0.0.1",
		 Port: "8088",
	 }
}


func(a *App)HttpAddr()string{
	 return fmt.Sprintf("%s:%s" , a.Host ,a.Port)
}

type App struct {
	Name      string `toml:"name" env:"APP_NAME"`
	Host      string `toml:"host" env:"APP_HOST"`
	Port      string `toml:"port" env:"APP_PORT"`
}


func NewDefaultLog()*Log{
	return &Log{
		Level: "info",
		Format: "json",
		To:ToStdout,
	}
}

// Log todo
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

var db *sql.DB

func (m *MySQL)GetDB()(*sql.DB){
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil{
		 conn ,err := m.getConn()
		 if err!=nil{
			return nil
		 }
		 db = conn
	}
	return  db
}


func (m *MySQL)getConn()(*sql.DB ,error){
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	ctx , cancle := context.WithTimeout(context.Background() , 5*time.Second) // 设置了超时取消
	defer  cancle()
	if  err := db.PingContext(ctx);err!=nil{
		 return nil , err
	}
	return  db ,nil
}


func NewDefaultMysql()*MySQL{
	 return &MySQL{
		  Host: "127.0.0.1",
		  Port: "3306",
		  UserName:"root",
		  Password: "123456",
		  Database: "test",
		  MaxOpenConn: 10,
		  MaxIdleConn: 5,
	 }
}

// MySQL todo
type MySQL struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	UserName    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	//mysql 连接池
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	//控制mysql mysql 复用
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	//连接声明周期
	//MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	////idle 连接 最多允许存活多久
	//MaxIdleTime int    `toml:"max_idle_time" env:"D_MYSQL_MAX_idle_TIME"`
	lock   sync.Mutex
}
