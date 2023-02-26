package app

import (
	"net"
	"net/rpc"
	"runtime"

	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/K0STYAa/vk_iproto/internal/delivery"
	"github.com/K0STYAa/vk_iproto/internal/repository"
	"github.com/K0STYAa/vk_iproto/internal/service"
	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/sirupsen/logrus"
)

type MyService struct {
	delivery *delivery.Delivery
}

func (ms *MyService) MainHandler(req models.Request, reply *models.Response) error {
	*reply = ms.delivery.MainHandler(req)
	return nil //nolint: nlreturn
}

func Run() {
	runtime.GOMAXPROCS(models.GoMaxProcsLim)
	models.LogStart()

	myStorage := new(internal.BaseStorage)
	repos := repository.NewRepository(myStorage)
	service := service.NewService(repos)
	delivery := delivery.NewDelivery(*service)
	myService := &MyService{delivery: delivery}

	err := rpc.Register(myService)
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
