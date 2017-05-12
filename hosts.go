package goforeman

import (
	"fmt"

	"github.com/amine7536/goforeman/context"
)

const hostBasePath = "api/hosts"

type HostsService interface {
	List(context.Context) ([]Host, *Response, error)
	Get(context.Context, string) (*Host, *Response, error)
}

type Host struct {
	IP                  string      `json:"ip"`
	IP6                 interface{} `json:"ip6"`
	EnvironmentID       interface{} `json:"environment_id"`
	EnvironmentName     interface{} `json:"environment_name"`
	LastReport          interface{} `json:"last_report"`
	Mac                 interface{} `json:"mac"`
	RealmID             interface{} `json:"realm_id"`
	RealmName           interface{} `json:"realm_name"`
	SpMac               interface{} `json:"sp_mac"`
	SpIP                interface{} `json:"sp_ip"`
	SpName              interface{} `json:"sp_name"`
	DomainID            interface{} `json:"domain_id"`
	DomainName          interface{} `json:"domain_name"`
	ArchitectureID      interface{} `json:"architecture_id"`
	ArchitectureName    interface{} `json:"architecture_name"`
	OperatingsystemID   interface{} `json:"operatingsystem_id"`
	OperatingsystemName interface{} `json:"operatingsystem_name"`
	SubnetID            interface{} `json:"subnet_id"`
	SubnetName          interface{} `json:"subnet_name"`
	Subnet6ID           interface{} `json:"subnet6_id"`
	Subnet6Name         interface{} `json:"subnet6_name"`
	SpSubnetID          interface{} `json:"sp_subnet_id"`
	PtableID            interface{} `json:"ptable_id"`
	PtableName          interface{} `json:"ptable_name"`
	MediumID            interface{} `json:"medium_id"`
	MediumName          interface{} `json:"medium_name"`
	Build               bool        `json:"build"`
	Comment             interface{} `json:"comment"`
	Disk                interface{} `json:"disk"`
	InstalledAt         interface{} `json:"installed_at"`
	ModelID             interface{} `json:"model_id"`
	HostgroupID         interface{} `json:"hostgroup_id"`
	OwnerID             interface{} `json:"owner_id"`
	OwnerType           interface{} `json:"owner_type"`
	Enabled             bool        `json:"enabled"`
	Managed             bool        `json:"managed"`
	UseImage            interface{} `json:"use_image"`
	ImageFile           string      `json:"image_file"`
	UUID                interface{} `json:"uuid"`
	ComputeResourceID   interface{} `json:"compute_resource_id"`
	ComputeResourceName interface{} `json:"compute_resource_name"`
	ComputeProfileID    interface{} `json:"compute_profile_id"`
	ComputeProfileName  interface{} `json:"compute_profile_name"`
	Capabilities        []string    `json:"capabilities"`
	ProvisionMethod     string      `json:"provision_method"`
	Certname            string      `json:"certname"`
	ImageID             interface{} `json:"image_id"`
	ImageName           interface{} `json:"image_name"`
	CreatedAt           string      `json:"created_at"`
	UpdatedAt           string      `json:"updated_at"`
	LastCompile         interface{} `json:"last_compile"`
	GlobalStatus        int         `json:"global_status"`
	GlobalStatusLabel   string      `json:"global_status_label"`
	OrganizationID      interface{} `json:"organization_id"`
	OrganizationName    interface{} `json:"organization_name"`
	LocationID          interface{} `json:"location_id"`
	LocationName        interface{} `json:"location_name"`
	PuppetStatus        int         `json:"puppet_status"`
	ModelName           interface{} `json:"model_name"`
	Name                string      `json:"name"`
	ID                  int         `json:"id"`
	PuppetProxyID       interface{} `json:"puppet_proxy_id"`
	PuppetProxyName     interface{} `json:"puppet_proxy_name"`
	PuppetCaProxyID     interface{} `json:"puppet_ca_proxy_id"`
	PuppetCaProxyName   interface{} `json:"puppet_ca_proxy_name"`
	PuppetProxy         interface{} `json:"puppet_proxy"`
	PuppetCaProxy       interface{} `json:"puppet_ca_proxy"`
	HostgroupName       interface{} `json:"hostgroup_name"`
	HostgroupTitle      interface{} `json:"hostgroup_title"`
}

type HostsServiceOp struct {
	client *Client
}

var _ HostsService = &HostsServiceOp{}

type hostsRoot struct {
	Hosts []Host `json:"results"`
}

// Convert Droplet to a string
func (h Host) String() string {
	return Stringify(h)
}

func (s *HostsServiceOp) List(ctx context.Context) ([]Host, *Response, error) {
	req, err := s.client.NewRequest(ctx, "GET", hostBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(hostsRoot)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, nil, err
	}

	return v.Hosts, resp, nil
}

func (s *HostsServiceOp) Get(ctx context.Context, hostname string) (*Host, *Response, error) {

	path := fmt.Sprintf("%s/%s", hostBasePath, hostname)

	req, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Host)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
