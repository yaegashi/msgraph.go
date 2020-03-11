// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SkypeForBusinessActivityUserDetail undocumented
type SkypeForBusinessActivityUserDetail struct {
	// Entity is the base model of SkypeForBusinessActivityUserDetail
	Entity
	// TotalPeerToPeerSessionCount undocumented
	TotalPeerToPeerSessionCount *int `json:"totalPeerToPeerSessionCount,omitempty"`
	// TotalOrganizedConferenceCount undocumented
	TotalOrganizedConferenceCount *int `json:"totalOrganizedConferenceCount,omitempty"`
	// TotalParticipatedConferenceCount undocumented
	TotalParticipatedConferenceCount *int `json:"totalParticipatedConferenceCount,omitempty"`
	// PeerToPeerLastActivityDate undocumented
	PeerToPeerLastActivityDate *Date `json:"peerToPeerLastActivityDate,omitempty"`
	// OrganizedConferenceLastActivityDate undocumented
	OrganizedConferenceLastActivityDate *Date `json:"organizedConferenceLastActivityDate,omitempty"`
	// ParticipatedConferenceLastActivityDate undocumented
	ParticipatedConferenceLastActivityDate *Date `json:"participatedConferenceLastActivityDate,omitempty"`
	// PeerToPeerIMCount undocumented
	PeerToPeerIMCount *int `json:"peerToPeerIMCount,omitempty"`
	// PeerToPeerAudioCount undocumented
	PeerToPeerAudioCount *int `json:"peerToPeerAudioCount,omitempty"`
	// PeerToPeerAudioMinutes undocumented
	PeerToPeerAudioMinutes *int `json:"peerToPeerAudioMinutes,omitempty"`
	// PeerToPeerVideoCount undocumented
	PeerToPeerVideoCount *int `json:"peerToPeerVideoCount,omitempty"`
	// PeerToPeerVideoMinutes undocumented
	PeerToPeerVideoMinutes *int `json:"peerToPeerVideoMinutes,omitempty"`
	// PeerToPeerAppSharingCount undocumented
	PeerToPeerAppSharingCount *int `json:"peerToPeerAppSharingCount,omitempty"`
	// PeerToPeerFileTransferCount undocumented
	PeerToPeerFileTransferCount *int `json:"peerToPeerFileTransferCount,omitempty"`
	// OrganizedConferenceIMCount undocumented
	OrganizedConferenceIMCount *int `json:"organizedConferenceIMCount,omitempty"`
	// OrganizedConferenceAudioVideoCount undocumented
	OrganizedConferenceAudioVideoCount *int `json:"organizedConferenceAudioVideoCount,omitempty"`
	// OrganizedConferenceAudioVideoMinutes undocumented
	OrganizedConferenceAudioVideoMinutes *int `json:"organizedConferenceAudioVideoMinutes,omitempty"`
	// OrganizedConferenceAppSharingCount undocumented
	OrganizedConferenceAppSharingCount *int `json:"organizedConferenceAppSharingCount,omitempty"`
	// OrganizedConferenceWebCount undocumented
	OrganizedConferenceWebCount *int `json:"organizedConferenceWebCount,omitempty"`
	// OrganizedConferenceDialInOut3rdPartyCount undocumented
	OrganizedConferenceDialInOut3rdPartyCount *int `json:"organizedConferenceDialInOut3rdPartyCount,omitempty"`
	// OrganizedConferenceCloudDialInOutMicrosoftCount undocumented
	OrganizedConferenceCloudDialInOutMicrosoftCount *int `json:"organizedConferenceCloudDialInOutMicrosoftCount,omitempty"`
	// OrganizedConferenceCloudDialInMicrosoftMinutes undocumented
	OrganizedConferenceCloudDialInMicrosoftMinutes *int `json:"organizedConferenceCloudDialInMicrosoftMinutes,omitempty"`
	// OrganizedConferenceCloudDialOutMicrosoftMinutes undocumented
	OrganizedConferenceCloudDialOutMicrosoftMinutes *int `json:"organizedConferenceCloudDialOutMicrosoftMinutes,omitempty"`
	// ParticipatedConferenceIMCount undocumented
	ParticipatedConferenceIMCount *int `json:"participatedConferenceIMCount,omitempty"`
	// ParticipatedConferenceAudioVideoCount undocumented
	ParticipatedConferenceAudioVideoCount *int `json:"participatedConferenceAudioVideoCount,omitempty"`
	// ParticipatedConferenceAudioVideoMinutes undocumented
	ParticipatedConferenceAudioVideoMinutes *int `json:"participatedConferenceAudioVideoMinutes,omitempty"`
	// ParticipatedConferenceAppSharingCount undocumented
	ParticipatedConferenceAppSharingCount *int `json:"participatedConferenceAppSharingCount,omitempty"`
	// ParticipatedConferenceWebCount undocumented
	ParticipatedConferenceWebCount *int `json:"participatedConferenceWebCount,omitempty"`
	// ParticipatedConferenceDialInOut3rdPartyCount undocumented
	ParticipatedConferenceDialInOut3rdPartyCount *int `json:"participatedConferenceDialInOut3rdPartyCount,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// DeletedDate undocumented
	DeletedDate *Date `json:"deletedDate,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// AssignedProducts undocumented
	AssignedProducts []string `json:"assignedProducts,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
