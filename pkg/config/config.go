package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Domain     string
	APIKey     string
	OutputFile string
	Silent     bool
}

func Parse() (*Config, error) {
	var c Config

	flag.StringVar(&c.Domain, "d", "", "domain to search for")
	flag.StringVar(&c.APIKey, "apikey", os.Getenv("NETLAS_APIKEY"), "The Netlas API key")
	flag.StringVar(&c.OutputFile, "of", "", "output file for results")
	flag.BoolVar(&c.Silent, "silent", false, "silent mode, no stdout output")

	flag.Parse()

	if c.Domain == "" {
		return nil, fmt.Errorf("no domain provided")
	}

	if c.APIKey == "" {
		return nil, fmt.Errorf("no Netlas API Key specified")
	}

	return &c, nil
}
