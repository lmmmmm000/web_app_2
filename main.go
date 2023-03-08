package main

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/controllers"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/pkg/snowflake"
	"web_app/routers"
	"web_app/settings"
)

func main() {
	//1. 加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("settings.Init failed ,err:%v\n", err)

	}
	//2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("logger.Init failed ,err:%v\n", err)

	}
	defer zap.L().Sync() //缓冲区日志追加到日志文件中
	zap.L().Debug("logger init success")
	//3. 初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("mysql.Init failed ,err:%v\n", err)

	}
	defer mysql.Close()
	//4. 初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis.Init failed ,err:%v\n", err)

	}
	defer redis.Close()

	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId); err != nil {
		fmt.Printf("snowflake.Init failed ,err:%v\n", err)
		return
	}

	//初始化gin框架内置的校验器使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("controllers.Init failed ,err:%v\n", err)
		return
	}
	//5. 注册路由
	r := routers.SetupRouter()
	//6. 启动服务(优雅关机)
	fmt.Println(settings.Conf.Port)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
