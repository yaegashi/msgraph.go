// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EmailAppUsageVersionsUserCounts undocumented
type EmailAppUsageVersionsUserCounts struct {
	// Entity is the base model of EmailAppUsageVersionsUserCounts
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Outlook2016 undocumented
	Outlook2016 *int `json:"outlook2016,omitempty"`
	// Outlook2013 undocumented
	Outlook2013 *int `json:"outlook2013,omitempty"`
	// Outlook2010 undocumented
	Outlook2010 *int `json:"outlook2010,omitempty"`
	// Outlook2007 undocumented
	Outlook2007 *int `json:"outlook2007,omitempty"`
	// Undetermined undocumented
	Undetermined *int `json:"undetermined,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
