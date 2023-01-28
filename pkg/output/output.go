package output

import (
	"encoding/csv"
	"os"

	"github.com/kenjoe41/netdomain/pkg/config"
)

// WriteToFile writes the list of domains to a file specified in the config and also handles whether to print the domains to the console or not based on the Silent field in the config struct.
func WriteToFile(domains []string, cfg *config.Config) error {
	if cfg.OutputFile != "" {
		file, err := os.Create(cfg.OutputFile)
		if err != nil {
			return err
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for _, domain := range domains {
			err := writer.Write([]string{domain})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
