// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SkypeForBusinessPeerToPeerActivityCounts undocumented
type SkypeForBusinessPeerToPeerActivityCounts struct {
	// Entity is the base model of SkypeForBusinessPeerToPeerActivityCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// Audio undocumented
	Audio *int `json:"audio,omitempty"`
	// Video undocumented
	Video *int `json:"video,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// FileTransfer undocumented
	FileTransfer *int `json:"fileTransfer,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
