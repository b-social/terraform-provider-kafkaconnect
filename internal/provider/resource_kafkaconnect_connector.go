package provider

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-kafka/connect"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gopkg.in/matryer/try.v1"
)

func newConnector() *schema.Resource {
	return &schema.Resource{
		CreateContext: createConnector,
		ReadContext:   readConnector,
		UpdateContext: updateConnector,
		DeleteContext: deleteConnector,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"maximum_tasks": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"key_converter_class": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_converter_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"value_converter_class": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"value_converter_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func createConnector(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	connector := buildConnector(data)
	client := meta.(*connect.Client)

	if _, err := client.CreateConnector(connector); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(connector.Name)

	return readConnector(ctx, data, meta)
}

func readConnector(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	exists, err := checkIfConnectorExists(data, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	if !exists {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Connector does not exist",
			Detail:   "Connector does not exist",
		})
		return diags
	}

	client := meta.(*connect.Client)
	connector, _, err := client.GetConnector(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	data.Set("name", connector.Name)
	data.Set("class", connector.Config["connector.class"])

	maximumTasks, err := strconv.Atoi(connector.Config["tasks.max"])
	if err != nil {
		return diag.FromErr(err)
	}
	data.Set("maximum_tasks", maximumTasks)

	keyConverterClass := connector.Config["key.converter"]
	if keyConverterClass != "" {
		data.Set("key_converter_class", keyConverterClass)
	}

	valueConverterClass := connector.Config["value.converter"]
	if valueConverterClass != "" {
		data.Set("value_converter_class", valueConverterClass)
	}

	keyConverterConfiguration := make(map[string]string)
	keyConverterRegexp, _ := regexp.Compile("^key\\.converter")

	valueConverterConfiguration := make(map[string]string)
	valueConverterRegexp, _ := regexp.Compile("^value\\.converter")

	configuration := make(map[string]string)

	for key, value := range connector.Config {
		if key == "name" || key == "connector.class" || key == "tasks.max" {
			continue
		}

		if strings.HasPrefix(key, "key.converter.") {
			strippedKey := keyConverterRegexp.ReplaceAllString(
				"key.converter.", "")
			keyConverterConfiguration[strippedKey] = value
			continue
		}

		if strings.HasPrefix(key, "value.converter.") {
			strippedKey := valueConverterRegexp.ReplaceAllString(
				"value.converter.", "")
			valueConverterConfiguration[strippedKey] = value
			continue
		}

		configuration[key] = value
	}

	if len(keyConverterConfiguration) != 0 {
		data.Set(
			"key_converter_configuration",
			keyConverterConfiguration)
	}

	if len(valueConverterConfiguration) != 0 {
		data.Set(
			"value_converter_configuration",
			valueConverterConfiguration)
	}

	if len(configuration) != 0 {
		data.Set("configuration", configuration)
	}

	return diags
}

func updateConnector(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	connector := buildConnector(data)
	client := meta.(*connect.Client)

	_, _, err := client.UpdateConnectorConfig(data.Id(), connector.Config)
	if err != nil {
		return diag.FromErr(err)
	}

	return readConnector(ctx, data, meta)
}

func deleteConnector(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*connect.Client)

	var diags diag.Diagnostics

	_, err := client.DeleteConnector(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func checkIfConnectorExists(
	data *schema.ResourceData,
	meta interface{}) (bool, error) {
	client := meta.(*connect.Client)

	var attempts = 60
	var delay = 5 * time.Second
	var exists bool
	retryError := try.Do(func(attempt int) (bool, error) {
		_, _, err := client.GetConnectorStatus(data.Id())
		if err != nil {
			if apiError, ok := err.(connect.APIError); ok {
				if apiError.Code == 404 {
					if attempt < attempts {
						time.Sleep(delay)
						return true, nil
					}
					return false, nil
				}
			}
			if attempt < attempts {
				time.Sleep(delay)
				return true, err
			}
			return false, err
		}
		exists = true
		return false, nil
	})

	if retryError != nil {
		return false, retryError
	}

	return exists, nil
}

func buildConnector(d *schema.ResourceData) *connect.Connector {
	connectorConfig := connect.ConnectorConfig{
		"name":            d.Get("name").(string),
		"connector.class": d.Get("class").(string),
		"tasks.max":       strconv.Itoa(d.Get("maximum_tasks").(int)),
	}

	if keyConverterClass, ok := d.GetOk("key_converter_class"); ok {
		connectorConfig["key.converter"] = keyConverterClass.(string)

		if keyConverterConfiguration, ok :=
			d.GetOk("key_converter_configuration"); ok {
			for key, val := range keyConverterConfiguration.(map[string]interface{}) {
				connectorConfig["key.converter."+key] = val.(string)
			}
		}
	}

	if valueConverterClass, ok := d.GetOk("value_converter_class"); ok {
		connectorConfig["value.converter"] = valueConverterClass.(string)

		if valueConverterConfiguration, ok :=
			d.GetOk("value_converter_configuration"); ok {
			for key, val := range valueConverterConfiguration.(map[string]interface{}) {
				connectorConfig["value.converter."+key] = val.(string)
			}
		}
	}

	if configuration, ok := d.GetOk("configuration"); ok {
		for key, val := range configuration.(map[string]interface{}) {
			connectorConfig[key] = val.(string)
		}
	}

	connector := &connect.Connector{
		Name:   d.Get("name").(string),
		Config: connectorConfig,
	}

	return connector
}
