package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bombsimon/logrusr/v2"
	"github.com/go-logr/glogr"
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"github.com/go-logr/zapr"
	"github.com/go-logr/zerologr"
	"github.com/iand/logfmtr"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"k8s.io/klog/v2/klogr"
)

func printAllLevels(name string, logger logr.Logger) {
	fmt.Println("#####################")
	fmt.Println(name)
	fmt.Println("#####################")

	fmt.Fprintln(os.Stderr, "#####################")
	fmt.Fprintln(os.Stderr, name)
	fmt.Fprintln(os.Stderr, "#####################")

	fmt.Println("#####################")
	fmt.Println("Logging with No level")
	logger.Info("This is Info no level", "level", "None")
	logger.Error(nil, "This is Error no level", "level", "None")

	fmt.Println("#####################")
	fmt.Println("Logging with Level 1")
	logger.V(1).Info("This is Info Level 1", "level", "None")
	logger.V(1).Error(nil, "This is Error Level 1", "level", "None")

	fmt.Println("#####################")
	fmt.Println("Logging with Level 2")
	logger.V(2).Info("This is Info Level 2", "level", "None")
	logger.V(2).Error(nil, "This is Error Level 2", "level", "None")

	fmt.Println("#####################")
	fmt.Println("Logging with Level 3")
	logger.V(3).Info("This is Info Level 3", "level", "None")
	logger.V(3).Error(nil, "This is Error Level 3", "level", "None")

	fmt.Println("#####################")
	fmt.Println("Logging with Level 4")
	logger.V(4).Info("This is Info Level 4", "level", "None")
	logger.V(4).Error(nil, "This is Error Level 4", "level", "None")

	fmt.Println("#####################")
	fmt.Println("Logging with Level 5")
	logger.V(5).Info("This is Info Level 5", "level", "None")
	logger.V(5).Error(nil, "This is Error Level 5", "level", "None")

	fmt.Println("")
}

// This is just a survey to understand what is the default logging levels for most logging implementations
// Run this command like this `go run . > stdout.log 2> stderr.log`
func main() {
	printAllLevels("discard", logr.Discard())

	flag.Parse()
	printAllLevels("glogr", glogr.New())

	printAllLevels("klogr", klogr.New())

	zl, _ := zap.NewProduction()
	printAllLevels("zapr", zapr.NewLogger(zl))

	printAllLevels("stdr", stdr.New(log.New(os.Stdout, "", 0)))

	printAllLevels("logrusr", logrusr.New(logrus.New()))

	printAllLevels("logfmtr", logfmtr.New())

	zerol := zerolog.New(os.Stdout)
	printAllLevels("zerologr", zerologr.New(&zerol))
}
