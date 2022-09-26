package populate

import (
	"fmt"

	"github.com/magodo/armid"
	"github.com/magodo/aztft/internal/client"
)

// populateFunc populates the hypothetic azure resource ids that represent the property like resources of the specified resource.
type populateFunc func(*client.ClientBuilder, armid.ResourceId) ([]armid.ResourceId, error)

var populaters = map[string]populateFunc{
	"azurerm_linux_virtual_machine": populateVirtualMachine,
}

func NeedsAPI(rt string) bool {
	_, ok := populaters[rt]
	return ok
}

func Populate(id armid.ResourceId, rt string) ([]armid.ResourceId, error) {
	populater, ok := populaters[rt]
	if !ok {
		return nil, nil
	}

	b, err := client.NewClientBuilder()
	if err != nil {
		return nil, fmt.Errorf("new API client builder: %v", err)
	}

	return populater(b, id)
}
