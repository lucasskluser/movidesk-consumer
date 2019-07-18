package models

type Ticket struct {
	ID                 int    `json:"id"`
	Type               int    `json:"type"`
	Subject            string `json:"subject"`
	Category           string `json:"category"`
	Urgency            string `json:"urgency"`
	Status             string `json:"status"`
	BaseStatus         string `json:"baseStatus"`
	Justification      string `json:"justification"`
	Origin             int    `json:"origin"`
	CreatedDate        string `json:"createdDate"`
	IsDeleted          bool   `json:"isDeleted"`
	OriginEmailAccount string `json:"originEmailAccount"`
	Owner              People `json:"owner"`
	OwnerTeam          string `json:"ownerTeam"`
	CreatedBy          People `json:"createdBy"`
}

type People struct {
	ID           string `json:"id"`
	PersonType   int    `json:"personType"`
	ProfileType  int    `json:"profileType"`
	BusinessName string `json:"businessName"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Complement   string `json:"complement"`
	Cep          string `json:"cep"`
	City         string `json:"city"`
	Bairro       string `json:"bairro"`
	Number       string `json:"number"`
	Reference    string `json:"reference"`
}

