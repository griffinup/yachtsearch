package schema

type Yacht struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
	Company	int `json:"companyId"`
	Model	int `json:"yachtModelId"`
}

type InfoResult struct {
	ID	int		`json:"id"`
	Name	string	`json:"name"`
	ModelName	string	`json:"modelname"`
	BuilderName	string	`json:"buildername"`
	CompanyName	string	`json:"companyname"`
	Avail	bool	`json:"avail"`
}

type LiveResult struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
	Type	string `json:"type"`
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
	Builder      int    `json:"yachtBuilderId"`
}

type ModelsAllResponse struct {
	Status   string `json:"status"`
	Models []Model `json:"models"`
}

type Builder struct {
	ID        int    `json:"id"`
	Name      string    `json:"name"`
}

type BuildersAllResponse struct {
	Status   string `json:"status"`
	Builders []Builder `json:"builders"`
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
