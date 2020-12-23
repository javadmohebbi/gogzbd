package gobdgz

import "fmt"

type Network struct {
	request Request
}

type GetContainersResponse struct {
	ID      *string               `json:"id"`
	JSONRPC string                `json:"jsonrpc"`
	Result  []GetContainersResult `json:"result"`
}
type GetContainersResult struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetNetworkInventoryItemsResponse struct {
	ID      *string                                `json:"id"`
	JSONRPC string                                 `json:"jsonrpc"`
	Result  GetNetworkInventoryItemsResponseResult `json:"result"`
}
type GetNetworkInventoryItemsResponseResult struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	PagesCount int `json:"pagesPage"`

	Items []GetNetworkInventoryItemsResponseResultItems `json:"items"`
}
type InventoryObjectType int

func (v InventoryObjectType) String() string {
	str := "!!undocumented!!"
	switch v {
	case 4:
		str = "Group"
	case 5:
		str = "Computer"
	case 6:
		str = "Virtual Machine"
	case 8:
		str = "Virtual Host"
	case 9:
		str = "vShield App"
	case 10:
		str = "Virtualization Cluster"
	case 11:
		str = "Virtualization Datacenter"
	case 12:
		str = "Resource Pool"
	case 13:
		str = "Virtualization Pool"
	}

	return fmt.Sprintf("(%d) %v", v, str)
}

type GetNetworkInventoryItemsResponseResultItems struct {
	ID       string              `json:"id"`
	Name     string              `json:"name"`
	Type     InventoryObjectType `json:"type"`
	ParentID string              `json:"parentId"`

	// details available only for type 5 , 6
	Details GetNetworkInventoryItemsResponseResultItemsDetails `json:"details"`
}
type InventoryObjectMachineType int

func (v InventoryObjectMachineType) String() string {
	str := "!!undocumented!!"
	switch v {
	case 1:
		str = "computer"
	case 2:
		str = "virtual machine"
	case 0:
		str = "other"
	}
	return fmt.Sprintf("(%d) %v", v, str)
}

type GetNetworkInventoryItemsResponseResultItemsDetails struct {
	Label                  string                     `json:"label"`
	FQDN                   string                     `json:"fqdn"`
	GroupID                string                     `json:"groupId"`
	IsManaged              bool                       `json:"isManaged"`
	MachineType            InventoryObjectMachineType `json:"machineType"`
	OperatingSystemVersion string                     `json:"operatingSystemVersion"`
	IP                     string                     `json:"ip"`
	MACs                   []string                   `json:"macs"`
}

// prepare request before sening the request
func (n *Network) SetRequest(r Request) {
	n.request = r
}

// This method returns network containers. It will return an empty list if the parentId
// is not a container or does not contain any other container within it.
func (n *Network) GetContainers() (GetContainersResponse, error) {
	var resp GetContainersResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method returns network inventory items.
func (n *Network) GetNetworkInventoryItems() (GetNetworkInventoryItemsResponse, error) {
	var resp GetNetworkInventoryItemsResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}
