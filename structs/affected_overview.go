package structs

type AffectedOverview struct {
	UpdateTimestamp int64          `json:"updateTimestamp"`
	Date            int64          `json:"date"`
	AffectedAreas   []AffectedArea `json:"affectedAreas"`
}
