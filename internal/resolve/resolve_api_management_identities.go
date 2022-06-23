package resolve

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
	"github.com/magodo/aztft/internal/client"
	"github.com/magodo/aztft/internal/resourceid"
)

func resolveApiManagementIdentities(b *client.ClientBuilder, id resourceid.ResourceId) (string, error) {
	it := id.Names()[1]
	switch strings.ToUpper(it) {
	case strings.ToUpper(string(armapimanagement.IdentityProviderTypeAAD)):
		return "azurerm_api_management_identity_provider_aad", nil
	case strings.ToUpper(string(armapimanagement.IdentityProviderTypeAADB2C)):
		return "azurerm_api_management_identity_provider_aadb2c", nil
	case strings.ToUpper(string(armapimanagement.IdentityProviderTypeFacebook)):
		return "azurerm_api_management_identity_provider_facebook", nil
	case strings.ToUpper(string(armapimanagement.IdentityProviderTypeGoogle)):
		return "azurerm_api_management_identity_provider_google", nil
	case strings.ToUpper(string(armapimanagement.IdentityProviderTypeMicrosoft)):
		return "azurerm_api_management_identity_provider_microsoft", nil
	case strings.ToUpper(string(armapimanagement.IdentityProviderTypeTwitter)):
		return "azurerm_api_management_identity_provider_twitter", nil
	default:
		return "", fmt.Errorf("unknown identity provider type: %s", it)
	}
}