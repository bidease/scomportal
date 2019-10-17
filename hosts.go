package scomportal

// GetBaremetalHosts ..
type GetBaremetalHosts struct {
	Data []struct {
		ID                 uint64
		Title              string
		Conf               string
		HasDRAC            bool   `json:"has_drac"`
		DRACIsEnable       string `json:"temporary_drac_access"`
		ScheduledReleaseAt string `json:"scheduled_release_at"`
		Location           struct {
			Name     string
			Timezone string
		}
		Networks []struct {
			ID       uint64
			HostIP   string `json:"host_ip"`
			PoolType string `json:"pool_type"` // drac, private, public
			Netmask  string
		}
		OSReinstallation bool `json:"os_reinstallation"`
	}
	NumFound uint64 `json:"num_found"`
}

// GetBaremetalHosts ..
func (api *API) GetBaremetalHosts() (*GetBaremetalHosts, error) {
	var hosts GetBaremetalHosts
	if err := api.getRequest("/hosts", &hosts); err != nil {
		return nil, err
	}
	return &hosts, nil
}
