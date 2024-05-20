package cmd

import (
	"fmt"
	"github.com/cossim/coss-cli/pkg/apisix"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

func uploadSSL(context *cli.Context) error {
	certPath := context.String("cert")
	keyPath := context.String("private_key")
	domain := context.String("domain")
	apiKey := context.String("key")
	host := context.String("host")
	num := context.Int("num")

	cert, err := ioutil.ReadFile(certPath)
	if err != nil {
		return fmt.Errorf("failed to read certificate file: %v", err)
	}

	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("failed to read key file: %v", err)
	}

	if domain == "" {
		return fmt.Errorf("domain name is required")
	}
	client := apisix.NewApiClient(apiKey, host)
	if err := client.UpdateSSL(cert, key, []string{domain}, num); err != nil {
		fmt.Printf("Error updating SSL: %v\n", err)
	} else {
		fmt.Printf("SSL updated successfully\n")
	}
	return nil
}
