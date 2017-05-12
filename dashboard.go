package goforeman

import (
	"github.com/amine7536/goforeman/context"
)

const dashboardBasePath = "api/dashboard"

type DashboardService interface {
	Get(context.Context) (*Dashboard, *Response, error)
}

type Dashboard struct {
	TotalHosts            int `json:"total_hosts"`
	BadHosts              int `json:"bad_hosts"`
	BadHostsEnabled       int `json:"bad_hosts_enabled"`
	ActiveHosts           int `json:"active_hosts"`
	ActiveHostsOk         int `json:"active_hosts_ok"`
	ActiveHostsOkEnabled  int `json:"active_hosts_ok_enabled"`
	OkHosts               int `json:"ok_hosts"`
	OkHostsEnabled        int `json:"ok_hosts_enabled"`
	OutOfSyncHosts        int `json:"out_of_sync_hosts"`
	OutOfSyncHostsEnabled int `json:"out_of_sync_hosts_enabled"`
	DisabledHosts         int `json:"disabled_hosts"`
	PendingHosts          int `json:"pending_hosts"`
	PendingHostsEnabled   int `json:"pending_hosts_enabled"`
	GoodHosts             int `json:"good_hosts"`
	GoodHostsEnabled      int `json:"good_hosts_enabled"`
	Percentage            int `json:"percentage"`
	ReportsMissing        int `json:"reports_missing"`
	Glossary              struct {
		TotalHosts            string `json:"total_hosts"`
		BadHosts              string `json:"bad_hosts"`
		BadHostsEnabled       string `json:"bad_hosts_enabled"`
		ActiveHosts           string `json:"active_hosts"`
		ActiveHostsOk         string `json:"active_hosts_ok"`
		ActiveHostsOkEnabled  string `json:"active_hosts_ok_enabled"`
		OkHosts               string `json:"ok_hosts"`
		OkHostsEnabled        string `json:"ok_hosts_enabled"`
		OutOfSyncHosts        string `json:"out_of_sync_hosts"`
		OutOfSyncHostsEnabled string `json:"out_of_sync_hosts_enabled"`
		DisabledHosts         string `json:"disabled_hosts"`
		PendingHosts          string `json:"pending_hosts"`
		PendingHostsEnabled   string `json:"pending_hosts_enabled"`
		GoodHosts             string `json:"good_hosts"`
		GoodHostsEnabled      string `json:"good_hosts_enabled"`
		Percentage            string `json:"percentage"`
		ReportsMissing        string `json:"reports_missing"`
	} `json:"glossary"`
}

type DashboardServiceOp struct {
	client *Client
}

var _ DashboardService = &DashboardServiceOp{}

func (d Dashboard) String() string {
	return Stringify(d)
}

func (s *DashboardServiceOp) Get(ctx context.Context) (*Dashboard, *Response, error) {

	req, err := s.client.NewRequest(ctx, "GET", dashboardBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Dashboard)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
