package main

import (
	"code.cloudfoundry.org/lager/v3"
	"os"
)

func main() {

	logger := lager.NewLogger("my-app")
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.INFO))
	logger.Info("doing-stuff", lager.Data{
		"informative": true,
	})

}
