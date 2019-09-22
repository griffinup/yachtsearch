package schema

type Yacht struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
	Company	int `json:"companyId"`
	Model	int `json:"yachtModelId"`
}

type YachtAllResponse struct {
	Status   string `json:"status"`
	Yachts []Yacht `xml:"yachts"`
}

type YachtCompany struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
}

type YachtBuilder struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
}

type YachtModel struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
}
