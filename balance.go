package scomportal

// BaremetalBalance ..
type BaremetalBalance struct {
	Data struct {
		Balance          string
		EstimatedBalance string `json:"estimated_balance"`
	}
}

// GetBaremetalBalance ..
func (api *API) GetBaremetalBalance() (*BaremetalBalance, error) {
	var balances BaremetalBalance
	if err := api.getRequest("/statement/balance", &balances); err != nil {
		return nil, err
	}
	return &balances, nil
}
