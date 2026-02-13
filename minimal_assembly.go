package main

import (
	"fmt"
	"github.com/avanha/pmaas-core"
	"github.com/avanha/pmaas-core/config"
)

func main() {
	fmt.Printf("Starting minimal PMAAS assembly\n")
        conf := config.NewConfig()
	conf.HttpPort = 8090
        var pmaas = core.NewPMAAS(conf)
	err := pmaas.Run()

	if err != nil {
		fmt.Printf("pmaas.Run completed with error: %s\n", err)
	}

        fmt.Printf("Minimal PMAAS assembly terminstated\n")
}
