package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/levonliu/micro-cart/common"
	"github.com/levonliu/micro-cart/domain/repository"
	service2 "github.com/levonliu/micro-cart/domain/service"
	"github.com/levonliu/micro-cart/handler"
	"github.com/levonliu/micro-cart/proto/cart"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var QPS = 100

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		logger.Error(err)
	}

	//注册中心
	consulRegistry := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"127.0.0.1:8500",
			}
		},
	)

	//链路追踪
	tracer, io, err := common.NewTracer("github.com/levonliu/micro-cart", "localhost:6831")
	if err != nil {
		logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	//获取 mysql 配置, 路径中不带前缀
	mysqlInfo := common.GetMysqlConfig(consulConfig, "mysql")
	//创建数据库连接
	//db, err := gorm.Open("mysql", "root:root@/go_micro_demo?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//禁止复表
	db.SingularTable(true)

	//初始化表
	//err = repository.NewCartRepository(db).InitTable()
	//if err != nil {
	//	logger.Error(err)
	//}

	// Create service
	service := micro.NewService(
		micro.Name("micro.service.cart"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8087"),
		//添加 consulRegistry 作为注册中心
		micro.Registry(consulRegistry),
		//绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	//init service
	service.Init()

	// Register handler
	cartDataService := service2.NewCartDataService(repository.NewCartRepository(db))
	cart.RegisterCartHandler(service.Server(), &handler.Cart{CartDataService: cartDataService})

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
