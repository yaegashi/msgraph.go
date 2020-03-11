// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Office365ActiveUserCounts undocumented
type Office365ActiveUserCounts struct {
	// Entity is the base model of Office365ActiveUserCounts
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Office365 undocumented
	Office365 *int `json:"office365,omitempty"`
	// Exchange undocumented
	Exchange *int `json:"exchange,omitempty"`
	// OneDrive undocumented
	OneDrive *int `json:"oneDrive,omitempty"`
	// SharePoint undocumented
	SharePoint *int `json:"sharePoint,omitempty"`
	// SkypeForBusiness undocumented
	SkypeForBusiness *int `json:"skypeForBusiness,omitempty"`
	// Yammer undocumented
	Yammer *int `json:"yammer,omitempty"`
	// Teams undocumented
	Teams *int `json:"teams,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
