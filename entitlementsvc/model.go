package entitlementsvc

type Entitlement struct {
	ID       string    `json:"id"`
	Features []Feature `json:"features,omitempty"`
}
type Feature struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	AccessLevel string `json:"level,omitempty"`
}

var entitlements = []Entitlement{
	{"E01", []Feature{
		{"F1", "Feature 1", "Rw"},
		{"F1", "Feature 2", "R"},
	}},
}