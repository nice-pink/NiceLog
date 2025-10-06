package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nice-pink/NiceLog/log"
)

// func main() {
// 	slog.Debug("debug log")
// 	slog.Info("info log")
// 	slog.Warn("warn log")
// 	slog.Error("error log")
// }

func main() {
	// TestNdJson()

	log.Print("print log")
	log.Println("println log")

	log.Debug("debug log")

	log.SetPrefix("test")

	log.Debug("debug log")

	log.SetPrefix("bla")

	log.Debug("debug log")

	log.SetLogLevelInfo()

	log.Debug("should not be logged, because log level is info")

	log.Info("info log")

	log.SetLogTimestamp(true)

	log.Info("info log with timestamp")

	const routines int = 10
	var wg sync.WaitGroup
	wg.Add(routines)
	for range routines {
		test := Test{
			Prefix: "prefix_" + time.Now().Format(time.DateTime),
		}
		go test.Test(&wg)
	}
	wg.Wait()

}

func TestNdJson() {
	var wg sync.WaitGroup
	wg.Add(1)
	func() {
		fmt.Println("TestNdJson")
		log.Connect(log.Connection{
			Address:    "http://localhost:9428/insert/jsonline",
			Protocol:   "ndjson",
			StreamName: "test",
		})
		log.Info("info log")
	}()
	wg.Wait()
}
