package scomportal

import "fmt"

// BaremetalHostNetwork ..
type BaremetalHostNetwork struct {
	ID       uint64
	HostIP   string `json:"host_ip"`
	PoolType string `json:"pool_type"` // drac, private, public
	Netmask  string
}

// BaremetalHostLocation ..
type BaremetalHostLocation struct {
	Name     string
	Timezone string
}

// BaremetalHosts ..
type BaremetalHosts struct {
	Data []struct {
		ID          uint64
		Title       string
		Conf        string
		ServiceType uint64 `json:"service_type"`
		Networks    []BaremetalHostNetwork
		Location    BaremetalHostLocation
	}
}

// GetBaremetalHosts ..
func (api *API) GetBaremetalHosts() (*BaremetalHosts, error) {
	var hosts BaremetalHosts
	if err := api.getRequest("/hosts", &hosts); err != nil {
		return nil, err
	}
	return &hosts, nil
}

// BaremetalHost ..
type BaremetalHost struct {
	Data struct {
		ID                 uint64
		Title              string
		Conf               string
		Networks           []BaremetalHostNetwork
		Location           BaremetalHostLocation
		ScheduledReleaseAt string `json:"scheduled_release_at"`
		RackID             uint64 `json:"rack_id"`
		OsID               uint64 `json:"os_id"`
		OS                 struct {
			Arch    string
			Name    string
			Version string
		}
		Domain              string
		OSReinstallation    bool   `json:"os_reinstallation"`
		RescueMode          bool   `json:"rescue_mode"`
		HostClass           string `json:"host_class"`
		PublicIP            string `json:"public_ip"`
		DRACIP              string `json:"drac_ip"`
		HasDRAC             bool   `json:"has_drac"`
		TemporaryDracAccess string `json:"temporary_drac_access"`
	}
}

// GetBaremetalHost ..
func (api *API) GetBaremetalHost(id uint64) (*BaremetalHost, error) {
	var host BaremetalHost
	if err := api.getRequest(fmt.Sprintf("/hosts/%d", id), &host); err != nil {
		return nil, err
	}
	return &host, nil
}

// BaremetalTraficHost ..
type BaremetalTraficHost struct {
	Data struct {
		Type   string
		Commit struct {
			TotalCommitValue            uint    `json:"total_commit_value"`
			CommitValueForBillingPeriod float64 `json:"commit_value_for_billing_period"`
			UsageQuantity               string  `json:"usage_quantity"`
			DateStart                   string  `json:"date_start"`
			DateEnd                     string  `json:"date_end"`
		}
		Overuse struct {
			UsageQuantity string `json:"usage_quantity"`
			DateStart     string `json:"date_start"`
			DateEnd       string `json:"date_end"`
		}
	}
}

// GetBaremetalHostTraffic ..
func (api *API) GetBaremetalHostTraffic(id uint64) (*BaremetalTraficHost, error) {
	var traffic BaremetalTraficHost
	if err := api.getRequest(fmt.Sprintf("/hosts/%d/traffic_summary", id), &traffic); err != nil {
		return nil, err
	}
	return &traffic, nil
}

// BaremetalServices ..
type BaremetalServices struct {
	Data []struct {
		ID               uint64
		Type             uint64
		ParentID         uint64 `json:"parent_id"`
		Description      string
		Label            string
		Comment          string
		DateStart        string  `json:"date_start"`
		DateEnd          string  `json:"date_end"`
		OriginalCurrency string  `json:"original_currency"`
		OriginalPrice    float64 `json:"original_price"`
		Currency         string
		Price            float64
	}
}

// GetServices ..
func (api API) GetServices(id uint64) (*BaremetalServices, error) {
	var services BaremetalServices
	if err := api.getRequest(fmt.Sprintf("/hosts/%d/services", id), &services); err != nil {
		return nil, err
	}
	return &services, nil
}
