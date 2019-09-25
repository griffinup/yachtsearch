package schema

type Yacht struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
	Company	int `json:"companyId"`
	Model	int `json:"yachtModelId"`
}

type YachtFull struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
	CompanyName	string `json:"companyName"`
	ModelName	string `json:"yachtModelName"`
}

type YachtAllResponse struct {
	Status   string `json:"status"`
	Yachts []Yacht `xml:"yachts"`
}

type Company struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
}

type CompaniesAllResponse struct {
	Status   string `json:"status"`
	Companies []Company `json:"companies"`
}

type Model struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
}

type ModelsAllResponse struct {
	Status   string `json:"status"`
	Models []Model `json:"models"`
}

type YachtCompany struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
}

type YachtBuilder struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
}

type YachtModel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
}
