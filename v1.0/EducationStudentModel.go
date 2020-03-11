// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationStudent undocumented
type EducationStudent struct {
	// Object is the base model of EducationStudent
	Object
	// GraduationYear undocumented
	GraduationYear *string `json:"graduationYear,omitempty"`
	// Grade undocumented
	Grade *string `json:"grade,omitempty"`
	// BirthDate undocumented
	BirthDate *Date `json:"birthDate,omitempty"`
	// Gender undocumented
	Gender *EducationGender `json:"gender,omitempty"`
	// StudentNumber undocumented
	StudentNumber *string `json:"studentNumber,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
}
