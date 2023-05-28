package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	instance.UnimplementedServiceServer

	app service.MetaServiceServer
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	i.app = app.GetGrpcApp(service.AppName).(service.MetaServiceServer)
	return nil
}

func (i *impl) Name() string {
	return instance.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	instance.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
