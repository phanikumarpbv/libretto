package flavors

import (
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud"
	os "github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud/openstack/db/v1/flavors"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud/pagination"
)

// List will list all available flavors.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return os.List(client)
}

// Get retrieves the details for a particular flavor.
func Get(client *gophercloud.ServiceClient, flavorID string) os.GetResult {
	return os.Get(client, flavorID)
}