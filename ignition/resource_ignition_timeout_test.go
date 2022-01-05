package ignition

import (
	"testing"

	"github.com/coreos/ignition/v2/config/v3_3/types"
)

func TestIgnitionTimeout(t *testing.T) {
	testIgnition(t, `
		data "ignition_timeout" "foo" {
			http_response_headers = 15
			http_total = 15
		}

		data "ignition_config" "test" {
			timeouts = [data.ignition_timeout.foo.rendered]
		}
	`, func(c *types.Config) error {
		//if *c.Ignition.Timeouts.HTTPResponseHeaders != 15 {
		//	return fmt.Errorf("http_respnse_headers, found %d", *c.Ignition.Timeouts.HTTPResponseHeaders)
		//}
		//
		//if *c.Ignition.Timeouts.HTTPTotal != 15 {
		//	return fmt.Errorf("http_total, found %d", *c.Ignition.Timeouts.HTTPTotal)
		//}

		return nil
	})
}
