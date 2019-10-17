package scomportal

// Hosts ..
type Hosts struct {
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

// GetHosts ..
func (api *API) GetHosts() (*Hosts, error) {
	var hosts Hosts
	if err := api.getRequest("/hosts", &hosts); err != nil {
		return nil, err
	}
	return &hosts, nil
}

// Equipment ..
type Equipment struct {
	Data struct {
		Equipment struct {
			Hosts uint64
			Racks uint64
		}
		IPS struct {
			Use     uint64
			Overuse uint64
		}
		Licenses uint64
		Contacts uint64
	}
}

// GetEquipment ..
func (api *API) GetEquipment() (*Equipment, error) {
	var equipment Equipment
	if err := api.getRequest("/stats", &equipment); err != nil {
		return nil, err
	}
	return &equipment, nil
}
