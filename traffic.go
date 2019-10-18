package scomportal

import "fmt"

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
