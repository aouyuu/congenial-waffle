package structs

type AffectedArea struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Affected int64  `json:"affected"`
}
