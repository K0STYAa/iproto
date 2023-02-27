package app

import (
	"net"
	"net/rpc"

	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/K0STYAa/vk_iproto/internal/iproto_server"
	"github.com/K0STYAa/vk_iproto/internal/storage"
	"github.com/K0STYAa/vk_iproto/internal/usecase"
	"github.com/K0STYAa/vk_iproto/pkg/iproto"
	"github.com/sirupsen/logrus"
)

type MyService struct {
	iprotoserver *iprotoserver.IprotoServer
}

func (ms *MyService) MainHandler(req iproto.Request, reply *iproto.Response) error {
	*reply = ms.iprotoserver.MainHandler(req)
	return nil //nolint: nlreturn
}

func Run() {
	myStorage := new(internal.BaseStorage)
	repos := storage.NewRepository(myStorage)
	usecase := usecase.NewUsecase(repos)
	iprotoserver := iprotoserver.NewIprotoServer(*usecase)
	myService := &MyService{iprotoserver: iprotoserver}

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
