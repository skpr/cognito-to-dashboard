package dashboard

import "fmt"

// GetURI returns the URI for a CloudWatch dashboard.
func GetURI(region, name string) string {
	return fmt.Sprintf("https://%s.console.aws.amazon.com/cloudwatch/home?region=%s#dashboards/dashboard/%s", region, region, name)
}
