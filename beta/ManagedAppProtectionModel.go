// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAppProtection Policy used to configure detailed management settings for a specified set of apps
type ManagedAppProtection struct {
	// ManagedAppPolicy is the base model of ManagedAppProtection
	ManagedAppPolicy
	// PeriodOfflineBeforeAccessCheck The period after which access is checked when the device is not connected to the internet.
	PeriodOfflineBeforeAccessCheck *Duration `json:"periodOfflineBeforeAccessCheck,omitempty"`
	// PeriodOnlineBeforeAccessCheck The period after which access is checked when the device is connected to the internet.
	PeriodOnlineBeforeAccessCheck *Duration `json:"periodOnlineBeforeAccessCheck,omitempty"`
	// AllowedInboundDataTransferSources Sources from which data is allowed to be transferred.
	AllowedInboundDataTransferSources *ManagedAppDataTransferLevel `json:"allowedInboundDataTransferSources,omitempty"`
	// AllowedOutboundDataTransferDestinations Destinations to which data is allowed to be transferred.
	AllowedOutboundDataTransferDestinations *ManagedAppDataTransferLevel `json:"allowedOutboundDataTransferDestinations,omitempty"`
	// OrganizationalCredentialsRequired Indicates whether organizational credentials are required for app use.
	OrganizationalCredentialsRequired *bool `json:"organizationalCredentialsRequired,omitempty"`
	// AllowedOutboundClipboardSharingLevel The level to which the clipboard may be shared between apps on the managed device.
	AllowedOutboundClipboardSharingLevel *ManagedAppClipboardSharingLevel `json:"allowedOutboundClipboardSharingLevel,omitempty"`
	// DataBackupBlocked Indicates whether the backup of a managed app's data is blocked.
	DataBackupBlocked *bool `json:"dataBackupBlocked,omitempty"`
	// DeviceComplianceRequired Indicates whether device compliance is required.
	DeviceComplianceRequired *bool `json:"deviceComplianceRequired,omitempty"`
	// ManagedBrowserToOpenLinksRequired Indicates whether internet links should be opened in the managed browser app.
	ManagedBrowserToOpenLinksRequired *bool `json:"managedBrowserToOpenLinksRequired,omitempty"`
	// SaveAsBlocked Indicates whether users may use the "Save As" menu item to save a copy of protected files.
	SaveAsBlocked *bool `json:"saveAsBlocked,omitempty"`
	// PeriodOfflineBeforeWipeIsEnforced The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
	PeriodOfflineBeforeWipeIsEnforced *Duration `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`
	// PinRequired Indicates whether an app-level pin is required.
	PinRequired *bool `json:"pinRequired,omitempty"`
	// MaximumPinRetries Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
	MaximumPinRetries *int `json:"maximumPinRetries,omitempty"`
	// SimplePinBlocked Indicates whether simplePin is blocked.
	SimplePinBlocked *bool `json:"simplePinBlocked,omitempty"`
	// MinimumPinLength Minimum pin length required for an app-level pin if PinRequired is set to True
	MinimumPinLength *int `json:"minimumPinLength,omitempty"`
	// PinCharacterSet Character set which may be used for an app-level pin if PinRequired is set to True.
	PinCharacterSet *ManagedAppPinCharacterSet `json:"pinCharacterSet,omitempty"`
	// PeriodBeforePinReset TimePeriod before the all-level pin must be reset if PinRequired is set to True.
	PeriodBeforePinReset *Duration `json:"periodBeforePinReset,omitempty"`
	// AllowedDataStorageLocations Data storage locations where a user may store managed data.
	AllowedDataStorageLocations []ManagedAppDataStorageLocation `json:"allowedDataStorageLocations,omitempty"`
	// ContactSyncBlocked Indicates whether contacts can be synced to the user's device.
	ContactSyncBlocked *bool `json:"contactSyncBlocked,omitempty"`
	// PrintBlocked Indicates whether printing is allowed from managed apps.
	PrintBlocked *bool `json:"printBlocked,omitempty"`
	// FingerprintBlocked Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
	FingerprintBlocked *bool `json:"fingerprintBlocked,omitempty"`
	// DisableAppPinIfDevicePinIsSet Indicates whether use of the app pin is required if the device pin is set.
	DisableAppPinIfDevicePinIsSet *bool `json:"disableAppPinIfDevicePinIsSet,omitempty"`
	// MinimumRequiredOsVersion Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredOsVersion *string `json:"minimumRequiredOsVersion,omitempty"`
	// MinimumWarningOsVersion Versions less than the specified version will result in warning message on the managed app from accessing company data.
	MinimumWarningOsVersion *string `json:"minimumWarningOsVersion,omitempty"`
	// MinimumRequiredAppVersion Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredAppVersion *string `json:"minimumRequiredAppVersion,omitempty"`
	// MinimumWarningAppVersion Versions less than the specified version will result in warning message on the managed app.
	MinimumWarningAppVersion *string `json:"minimumWarningAppVersion,omitempty"`
	// MinimumWipeOsVersion Versions less than or equal to the specified version will wipe the managed app and the associated company data.
	MinimumWipeOsVersion *string `json:"minimumWipeOsVersion,omitempty"`
	// MinimumWipeAppVersion Versions less than or equal to the specified version will wipe the managed app and the associated company data.
	MinimumWipeAppVersion *string `json:"minimumWipeAppVersion,omitempty"`
	// AppActionIfDeviceComplianceRequired Defines a managed app behavior, either block or wipe, when the device is either rooted or jailbroken, if DeviceComplianceRequired is set to true.
	AppActionIfDeviceComplianceRequired *ManagedAppRemediationAction `json:"appActionIfDeviceComplianceRequired,omitempty"`
	// AppActionIfMaximumPinRetriesExceeded Defines a managed app behavior, either block or wipe, based on maximum number of incorrect pin retry attempts.
	AppActionIfMaximumPinRetriesExceeded *ManagedAppRemediationAction `json:"appActionIfMaximumPinRetriesExceeded,omitempty"`
	// PinRequiredInsteadOfBiometricTimeout Timeout in minutes for an app pin instead of non biometrics passcode
	PinRequiredInsteadOfBiometricTimeout *Duration `json:"pinRequiredInsteadOfBiometricTimeout,omitempty"`
	// AllowedOutboundClipboardSharingExceptionLength Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
	AllowedOutboundClipboardSharingExceptionLength *int `json:"allowedOutboundClipboardSharingExceptionLength,omitempty"`
	// NotificationRestriction Specify app notification restriction
	NotificationRestriction *ManagedAppNotificationRestriction `json:"notificationRestriction,omitempty"`
	// PreviousPinBlockCount Requires a pin to be unique from the number specified in this property.
	PreviousPinBlockCount *int `json:"previousPinBlockCount,omitempty"`
	// ManagedBrowser Indicates in which managed browser(s) that internet links should be opened.
	ManagedBrowser *ManagedBrowserType `json:"managedBrowser,omitempty"`
	// MaximumAllowedDeviceThreatLevel Maximum allowed device threat level, as reported by the MTD app
	MaximumAllowedDeviceThreatLevel *ManagedAppDeviceThreatLevel `json:"maximumAllowedDeviceThreatLevel,omitempty"`
	// MobileThreatDefenseRemediationAction Determines what action to take if the mobile threat defense threat threshold isn't met. Warn isn't a supported value for this property
	MobileThreatDefenseRemediationAction *ManagedAppRemediationAction `json:"mobileThreatDefenseRemediationAction,omitempty"`
}
