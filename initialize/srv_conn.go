package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop_api/global"
	"mxshop_api/proto"
)

func InitSrvConn() {
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	zap.S().Infof("端22口： %d", global.ServerConfig.ConsulInfo.Host)
	zap.S().Infof("端22口： %d", global.ServerConfig.ConsulInfo.Port)
	if err != nil {
		zap.S().Infof(fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.UserSrvInfo.Name))
		zap.S().Infof("端22口： %d", global.ServerConfig.ConsulInfo.Port)
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	newGoodsClient := proto.NewGoodsClient(conn)
	global.GoodsSrvClient = newGoodsClient
}
