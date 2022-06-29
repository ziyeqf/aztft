package resolve

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
	"github.com/magodo/armid"
	"github.com/magodo/aztft/internal/client"
)

func resolveStreamAnalyticsInputs(b *client.ClientBuilder, id armid.ResourceId) (string, error) {
	resourceGroupId := id.RootScope().(*armid.ResourceGroup)
	client, err := b.NewStreamAnalyticsInputsClient(resourceGroupId.SubscriptionId)
	if err != nil {
		return "", err
	}
	resp, err := client.Get(context.Background(), resourceGroupId.Name, id.Names()[0], id.Names()[1], nil)
	if err != nil {
		return "", fmt.Errorf("retrieving %q: %v", id, err)
	}
	props := resp.Input.Properties
	if props == nil {
		return "", fmt.Errorf("unexpected nil property in response")
	}
	switch props := props.(type) {
	case *armstreamanalytics.StreamInputProperties:
		ds := props.Datasource
		if ds == nil {
			return "", fmt.Errorf("unexpected nil properties.datasource in response")
		}
		switch ds.(type) {
		case *armstreamanalytics.EventHubStreamInputDataSource:
			return "azurerm_stream_analytics_stream_input_eventhub", nil
		case *armstreamanalytics.BlobStreamInputDataSource:
			return "azurerm_stream_analytics_stream_input_blob", nil
		case *armstreamanalytics.IoTHubStreamInputDataSource:
			return "azurerm_stream_analytics_stream_input_iothub", nil
		default:
			return "", fmt.Errorf("unknown input property data source type: %T", ds)
		}
	case *armstreamanalytics.ReferenceInputProperties:
		ds := props.Datasource
		if ds == nil {
			return "", fmt.Errorf("unexpected nil properties.datasource in response")
		}
		switch ds.(type) {
		case *armstreamanalytics.AzureSQLReferenceInputDataSource:
			return "azurerm_stream_analytics_reference_input_mssql", nil
		case *armstreamanalytics.BlobReferenceInputDataSource:
			return "azurerm_stream_analytics_reference_input_blob", nil
		default:
			return "", fmt.Errorf("unknown input property data source type: %T", ds)
		}

	default:
		return "", fmt.Errorf("unknown input property type: %T", props)
	}
}