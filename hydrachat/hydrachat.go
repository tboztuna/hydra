package hydrachat

import (
	"fmt"
	"hydra/hlogger"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = hlogger.GetInstance()

func Run(connection string) error {
	l, err := net.Listen("tcp", connection)

	if err != nil {
		logger.Println("Error connecting to chat client", err)
		return err
	}

	r := CreateRoom("HydraChat")

	go func() {
		// Handle SIGINT and SIGTERM
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		l.Close()
		fmt.Println("Closing TCP connection")
		close(r.Quit)

		if r.ClientCount() > 0 {
			<-r.Msgch
		}
		os.Exit(0)
	}()

}
