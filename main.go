package main

import (
	// "flag"
	// "fmt"
	// "log"
	// "os"
	"zjz/library"
)

func main() {
	// var (
	// 	logFileName = flag.String("log", "cServer.log", "Log file name")
	// )
	// logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	// if logErr != nil {
	// 	fmt.Println("Fail to find", *logFile, "cServer start Failed")
	// 	os.Exit(1)
	// }
	// log.SetOutput(logFile)
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Println("---------------------------------------------------------------------")
	// log.Println("---------------------------------------------------------------------")
	// log.Println("程序正常启动")
	library.Start()
}
