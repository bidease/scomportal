package scomportal

// Balance ..
type Balance struct {
	Data struct {
		Balance          string
		EstimatedBalance string `json:"estimated_balance"`
	}
}

// GetBaremetalBalance ..
func (api *API) GetBaremetalBalance() (*Balance, error) {
	var balances Balance
	if err := api.getRequest("/statement/balance", &balances); err != nil {
		return nil, err
	}
	return &balances, nil
}
