package kafkaconnect

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-kafka/connect"
	"github.com/hashicorp/terraform/helper/schema"
)

func newConnector() *schema.Resource {
	return &schema.Resource{
		Create: createConnector,
		Read:   readConnector,
		Update: updateConnector,
		Delete: deleteConnector,
		Exists: checkIfConnectorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
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

func createConnector(data *schema.ResourceData, context interface{}) error {
	connector := buildConnector(data)
	client := context.(*connect.Client)

	if _, err := client.CreateConnector(connector); err != nil {
		return err
	}

	data.SetId(connector.Name)

	return readConnector(data, context)
}

func readConnector(data *schema.ResourceData, context interface{}) error {
	client := context.(*connect.Client)
	connector, _, err := client.GetConnector(data.Id())
	if err != nil {
		return err
	}

	data.Set("name", connector.Name)
	data.Set("class", connector.Config["connector.class"])

	maximumTasks, err := strconv.Atoi(connector.Config["tasks.max"])
	if err != nil {
		return err
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

	return nil
}

func updateConnector(data *schema.ResourceData, context interface{}) error {
	connector := buildConnector(data)
	client := context.(*connect.Client)

	_, _, err := client.UpdateConnectorConfig(data.Id(), connector.Config)
	if err != nil {
		return err
	}

	return readConnector(data, context)
}

func deleteConnector(data *schema.ResourceData, context interface{}) error {
	client := context.(*connect.Client)

	_, err := client.DeleteConnector(data.Id())
	if err != nil {
		return err
	}

	return nil
}

func checkIfConnectorExists(
	data *schema.ResourceData,
	context interface{}) (bool, error) {
	client := context.(*connect.Client)

	_, _, err := client.GetConnectorStatus(data.Id())
	if err != nil {
		if apiError, ok := err.(connect.APIError); ok  {
			if apiError.Code == 404 {
				return false, nil
			}
		}

		return false, err
	}

	return true, nil
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
