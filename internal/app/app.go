package app

import (
	"net"
	"net/rpc"
	"runtime"

    "github.com/sirupsen/logrus"

	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/K0STYAa/vk_iproto/internal/delivery"
	"github.com/K0STYAa/vk_iproto/internal/repository"
	"github.com/K0STYAa/vk_iproto/internal/service"
	"github.com/K0STYAa/vk_iproto/pkg/models"
)

type MyService struct {
	delivery *delivery.Delivery
}
func (ms *MyService) MainHandler(req models.Request, reply *models.Response) error {
	*reply = ms.delivery.MainHandler(req)
	return nil
}

func Run() {
	runtime.GOMAXPROCS(4)
	models.LogStart()
	my_storage := new(internal.BaseStorage)

	repos := repository.NewRepository(my_storage)
	service := service.NewService(repos)
	delivery := delivery.NewDelivery(*service)
	my_service := &MyService{delivery: delivery}
	
	err := rpc.Register(my_service)
	if err != nil {
		logrus.Fatal("Register Service error: ", err)
	}

    // Creating a listener on a local machine on port 8080
	listener, err := net.Listen("tcp", ":8080") //#nosec  G102 -- This is a false positive
	if err != nil {
		logrus.Fatal("ListenTCP error: ", err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			logrus.Fatal("Error closing listener: ", err)
		}
	}()

    logrus.Info("The server is running, listening to port 8080...")

    // Infinite loop for processing incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(conn)
	}
}