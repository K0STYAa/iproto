package app

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/K0STYAa/iproto/internal/iproto_server"
	"github.com/K0STYAa/iproto/internal/storage"
	"github.com/K0STYAa/iproto/pkg/iproto"
	"github.com/K0STYAa/iproto/pkg/prometheus"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

type MyService struct {
	iprotoserver *iprotoserver.IprotoServer
	rateLimiter  *rate.Limiter
}

const (
	rpsLimit       = 10000
	burstLimit     = 10000
	maxConnections = 100
	errTemplate    = "%w"
	connTimeout    = 15 * time.Second
)

func (ms *MyService) MainHandler(req iproto.Request, reply *iproto.Response) error {
	if err := ms.rateLimiter.Wait(context.Background()); err != nil {
		return fmt.Errorf(errTemplate, err)
	}

	start := time.Now()
	defer prometheus.Latency.Observe(time.Since(start).Seconds())

	*reply = ms.iprotoserver.MainHandler(req)

	return nil
}

func Run() {
	rateLimiter := rate.NewLimiter(rpsLimit, burstLimit)
	// Set up a counting semaphore to limit the number of connections to 100.
	semaphore := make(chan struct{}, maxConnections)
	// Set up a wait group to keep track of active connections.
	var waitGroup sync.WaitGroup

	go prometheus.InitPrometheus()

	myStorage := new(storage.BaseStorage)
	repos := storage.NewStorage(myStorage)
	iprotoserver := iprotoserver.NewIprotoServer(*repos)
	myService := &MyService{iprotoserver: iprotoserver, rateLimiter: rateLimiter}

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

		prometheus.ConnectionsCount.Inc()
		semaphore <- struct{}{}

		waitGroup.Add(1)
		// Serve the connection in a separate goroutine.
		go func() {
			defer func() {
				// Release the slot in the semaphore and mark the connection as done.
				<-semaphore
				waitGroup.Done()
				prometheus.ConnectionsCount.Dec()
			}()
			// Add connection timeout
			if err := conn.SetReadDeadline(time.Now().Add(connTimeout)); err != nil {
				logrus.Error("Connection closed: ", err)
			}

			rpc.ServeConn(conn)
		}()
	}
}
