// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SkypeForBusinessDeviceUsageDistributionUserCounts undocumented
type SkypeForBusinessDeviceUsageDistributionUserCounts struct {
	// Entity is the base model of SkypeForBusinessDeviceUsageDistributionUserCounts
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Windows undocumented
	Windows *int `json:"windows,omitempty"`
	// WindowsPhone undocumented
	WindowsPhone *int `json:"windowsPhone,omitempty"`
	// AndroidPhone undocumented
	AndroidPhone *int `json:"androidPhone,omitempty"`
	// IPhone undocumented
	IPhone *int `json:"iPhone,omitempty"`
	// IPad undocumented
	IPad *int `json:"iPad,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
