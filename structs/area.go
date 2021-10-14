package structs

type Area struct {
	ID              int64      `json:"id"`
	UpdateTimestamp int64      `json:"updateTimestamp"`
	Name            string     `json:"name"`
	District        []District `json:"district"`
}

type District struct {
	ID              int64         `json:"id"`
	UpdateTimestamp int64         `json:"updateTimestamp"`
	Name            string        `json:"name"`
	Subdistrict     []SubDistrict `json:"subdistrict"`
}
type SubDistrict struct {
	ID              int64  `json:"id"`
	UpdateTimestamp int64  `json:"updateTimestamp"`
	Name            string `json:"name"`
	Affected        int64  `json:"affected"`
}
