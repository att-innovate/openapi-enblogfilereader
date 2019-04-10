package main

import (
	lr "eNBlogReader/logreader"
	"flag"
	"fmt"
	"log"
	"time"
)

// cmdline arguments
var screen = flag.Int("screen", 0, "Path to log file on FS.")

func main() {
	flag.Parse()
	log.Printf("screen: %v", *screen)
	logFilePath := fmt.Sprintf("/screenlog.%v", *screen)
	log.Printf("New log file location: %s", logFilePath)
	logReaderWrapper(logFilePath)
}

func logReaderWrapper(logFilePath string) {
	log.Println("defer")
	lr.ReadLog(logFilePath)
	defer func() {
		if err := recover(); err != nil {
			log.Println("ERROR in logReaderWrapper execution - trying to retart")
			time.Sleep(10 * time.Second)
			go logReaderWrapper(logFilePath)
		}
	}()
}
