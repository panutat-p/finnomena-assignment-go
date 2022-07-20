package fund

import "time"

type Fund struct {
	MstarID          string    `json:"mstar_id"`
	ThailandFundCode string    `json:"thailand_fund_code"`
	NavReturn        float64   `json:"nav_return"`
	Nav              float64   `json:"nav"`
	NavDate          time.Time `json:"nav_date"`
	AvgReturn        float64   `json:"avg_return"`
}

func (a Fund) Compare(b Fund) int {
	if a.NavReturn > b.NavReturn {
		return 1
	}
	if a.NavReturn < b.NavReturn {
		return -1
	}
	return 0
}
