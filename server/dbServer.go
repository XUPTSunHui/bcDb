package main

// + build etcd
import (
	"database/sql"
	"flag"
	"log"
	"time"

	"github.com/XUPTSunHui/bcDb/dbhandle"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	dbaddr   = flag.String("dbaddr", "192.168.208.12:8972", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_db", "prefix path")
)

func main() {
	flag.Parse()

	dbi := new(dbhandle.Dbop)
	dbi.DataBase = "bctest"
	dbi.UserName = "root"
	dbi.Passwd = "root"

	// dataSource := dbi.UserName + ":" + dbi.Passwd + "@/" + dbi.DataBase + "?charset=utf8"
	dataSource := dbi.UserName + ":" + dbi.Passwd + "@tcp(192.168.208.12:3306)/" + dbi.DataBase + "?charset=utf8"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("打开mysql失败: %v", err)
	}
	dbi.Db = db

	log.Println("打开mysql数据库成功")

	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)

	err = db.Ping()
	if err != nil {
		log.Fatalf("测试mysql连接失败: %v", err)
	}
	log.Println("测试mysql连接成功")

	s := server.NewServer()
	log.Println("启动server @" + *dbaddr + ", 服务名称为 Dbop")

	//注册到etcd
	// addRegistryPlugin(s)

	err = s.RegisterName("Dbop", dbi, "")
	if err != nil {
		log.Fatalf("注册Dbop失败: %v", err)
	}

	err = s.Serve("tcp", *dbaddr)
	if err != nil {
		log.Fatalf("启动Dbop失败: %v", err)
	}

}

func addRegistryPlugin(s *server.Server) {
	log.Println("*dbaddr: ", *dbaddr)

	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *dbaddr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		//Services:       []string{"Dbop"},
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
