package marathon

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv("MARATHON_URL"); v != "" {
						return v, nil
					}
					return nil, nil
				},
				Description: "Marathon's Base HTTP URL",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"marathon_app": resourceMarathonApp(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Url: d.Get("url").(string),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return config.client, nil
}
