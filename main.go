package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"ottodigital.id/library/logger"
	"ottodigital.id/library/utils"
	"rose-be-go/routers"
	"runtime"
	"strconv"
	"syscall"
	_"rose-be-go/db"
	//_"rose-be-go/kafka/publisher/producer"
	//_ "rose-be-go/redis/redis_cluster"
	//_ "rose-be-go/redis/redis_single"
)


func main()  {
	// init Logger Zap with file
	sugarLogger := logger.GetLogger()
	defer sugarLogger.Sync()

	maxProc, _ := strconv.Atoi(utils.GetEnv("MAXPROCS", "1"))

	runtime.GOMAXPROCS(maxProc)

	var errChan = make(chan error, 1)

	go func() {
		listenAddress := utils.GetEnv("ROSE_BE_GO_PORT", "0.0.0.0:8098")

		fmt.Println("Starting @", listenAddress)

		sugarLogger.Info("Log Start", zap.String("Starting Listen Address: ", listenAddress))

		errChan <- routers.Server(listenAddress)
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		fmt.Println("got an interrupt, exiting...")
		sugarLogger.Error("Got an interrupt, exiting...")
	case err := <-errChan:
		if err != nil {
			fmt.Println("error while running api, exiting...", err)
			sugarLogger.Error("Error while running api, exiting... ", zap.Error(err))
		}
	}




}