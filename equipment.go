package scomportal

// BaremetalEquipment ..
type BaremetalEquipment struct {
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

// GetBaremetalEquipment ..
func (api *API) GetBaremetalEquipment() (*BaremetalEquipment, error) {
	var equipment BaremetalEquipment
	if err := api.getRequest("/stats", &equipment); err != nil {
		return nil, err
	}
	return &equipment, nil
}
