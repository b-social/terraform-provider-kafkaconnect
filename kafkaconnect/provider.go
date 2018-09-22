package kafkaconnect

import (
	"net/url"

	"github.com/go-kafka/connect"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc(
					"KAFKA_CONNECT_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"kafkaconnect_connector": newConnector(),
		},
		ConfigureFunc: configure,
	}
}

func configure(data *schema.ResourceData) (interface{}, error) {
	rawUrl := data.Get("url").(string)
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	return connect.NewClient(parsedUrl.String()), nil
}
