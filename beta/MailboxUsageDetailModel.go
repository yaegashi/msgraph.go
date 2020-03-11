// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MailboxUsageDetail undocumented
type MailboxUsageDetail struct {
	// Entity is the base model of MailboxUsageDetail
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// DeletedDate undocumented
	DeletedDate *Date `json:"deletedDate,omitempty"`
	// CreatedDate undocumented
	CreatedDate *Date `json:"createdDate,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// ItemCount undocumented
	ItemCount *int `json:"itemCount,omitempty"`
	// StorageUsedInBytes undocumented
	StorageUsedInBytes *int `json:"storageUsedInBytes,omitempty"`
	// DeletedItemCount undocumented
	DeletedItemCount *int `json:"deletedItemCount,omitempty"`
	// DeletedItemSizeInBytes undocumented
	DeletedItemSizeInBytes *int `json:"deletedItemSizeInBytes,omitempty"`
	// IssueWarningQuotaInBytes undocumented
	IssueWarningQuotaInBytes *int `json:"issueWarningQuotaInBytes,omitempty"`
	// ProhibitSendQuotaInBytes undocumented
	ProhibitSendQuotaInBytes *int `json:"prohibitSendQuotaInBytes,omitempty"`
	// ProhibitSendReceiveQuotaInBytes undocumented
	ProhibitSendReceiveQuotaInBytes *int `json:"prohibitSendReceiveQuotaInBytes,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
