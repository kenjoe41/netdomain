package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/kenjoe41/netdomain/pkg/config"
	"github.com/kenjoe41/netdomain/pkg/netlas"
	"github.com/kenjoe41/netdomain/pkg/output"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	netClient := netlas.NewClient(cfg.APIKey)

	var subdomains []string
	outputCh := make(chan string)

	var outputWG sync.WaitGroup
	outputWG.Add(1)

	go func() {
		defer outputWG.Done()
		for subdomain := range outputCh {
			subdomains = append(subdomains, subdomain)
			if cfg.OutputFile != "" && cfg.Silent {
				continue
			}
			fmt.Println(subdomain)
		}
	}()

	err = netlas.GetAllSubdomains(netClient, cfg.Domain, &outputCh)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfg.OutputFile != "" {
		err = output.WriteToFile(subdomains, cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
