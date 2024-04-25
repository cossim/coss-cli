package cmd

import (
	"fmt"
	"github.com/cossim/coss-cli/pkg/apisix"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

func initRoute(context *cli.Context) error {
	apiKey := context.String("key")
	host := context.String("host")
	direct := context.Bool("direct")
	domain := context.String("domain")
	livekitDomain := context.String("livekit")
	routeHost := context.String("route-host")

	baseURL := host + "/apisix/admin/routes/"

	client := apisix.NewApiClient(apiKey, baseURL)

	route := client.GetRoutes(domain, routeHost, livekitDomain, direct)

	for i, route := range route {
		resp, err := client.SendRequest("PUT", fmt.Sprintf("%d", i+1), route)
		if err != nil {
			fmt.Printf("Error sending request for route %d: %v\n", i+1, err)
			continue
		}
		fmt.Printf("Route %d created successfully: %s\n", i+1, resp)
	}
	return nil
}

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
