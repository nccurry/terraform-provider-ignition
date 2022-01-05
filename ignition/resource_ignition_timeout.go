package ignition

import (
	"encoding/json"
	"github.com/coreos/ignition/v2/config/validate"

	"github.com/coreos/ignition/v2/config/v3_3/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTimeout() *schema.Resource {
	return &schema.Resource{
		Exists: resourceTimeoutExists,
		Read:   resourceTimeoutRead,
		Schema: map[string]*schema.Schema{
			"http_response_headers": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"http_total": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"rendered": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTimeoutRead(d *schema.ResourceData, meta interface{}) error {
	id, err := buildTimeout(d)
	if err != nil {
		return err
	}

	d.SetId(id)
	return nil
}

func resourceTimeoutExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := buildTimeout(d)
	if err != nil {
		return false, err
	}

	return id == d.Id(), nil
}

func buildTimeout(d *schema.ResourceData) (string, error) {
	timeouts := &types.Timeouts{}
	hrh := d.Get("http_response_headers").(int)
	timeouts.HTTPResponseHeaders = &hrh

	ht := d.Get("http_total").(int)
	timeouts.HTTPTotal = &ht

	b, err := json.Marshal(timeouts)
	if err != nil {
		return "", err
	}

	err = d.Set("rendered", string(b))
	if err != nil {
		return "", err
	}

	return hash(string(b)), handleReport(validate.ValidateWithContext(new(*types.Timeouts), b))
}
