// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SkypeForBusinessOrganizerActivityMinuteCounts undocumented
type SkypeForBusinessOrganizerActivityMinuteCounts struct {
	// Entity is the base model of SkypeForBusinessOrganizerActivityMinuteCounts
	Entity
	// AudioVideo undocumented
	AudioVideo *int `json:"audioVideo,omitempty"`
	// DialInMicrosoft undocumented
	DialInMicrosoft *int `json:"dialInMicrosoft,omitempty"`
	// DialOutMicrosoft undocumented
	DialOutMicrosoft *int `json:"dialOutMicrosoft,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
