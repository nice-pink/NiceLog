package main

import (
	"sync"
	"time"

	"github.com/nice-pink/nicelog/log"
)

// func main() {
// 	slog.Debug("debug log")
// 	slog.Info("info log")
// 	slog.Warn("warn log")
// 	slog.Error("error log")
// }

func main() {
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
