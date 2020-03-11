// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// StandardTimeZoneOffset undocumented
type StandardTimeZoneOffset struct {
	// Object is the base model of StandardTimeZoneOffset
	Object
	// Time undocumented
	Time *TimeOfDay `json:"time,omitempty"`
	// DayOccurrence undocumented
	DayOccurrence *int `json:"dayOccurrence,omitempty"`
	// DayOfWeek undocumented
	DayOfWeek *DayOfWeek `json:"dayOfWeek,omitempty"`
	// Month undocumented
	Month *int `json:"month,omitempty"`
	// Year undocumented
	Year *int `json:"year,omitempty"`
}
