package cmd

import (
	"fmt"
	"github.com/cossim/coss-cli/pkg/apisix"
	"github.com/urfave/cli/v2"
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

	for i, r := range route {
		resp, err := client.SendRequest("PUT", fmt.Sprintf("%d", i+1), r)
		if err != nil {
			fmt.Printf("Error sending request for route %d: %v\n", i+1, err)
			continue
		}
		fmt.Printf("Route %d created successfully: %s\n", i+1, resp)
	}
	return nil
}
