package company

type NewCompany struct {
	CompanyName string `json:"companyName"`
	CompanySize string `json:"companySize"`

	RootUserEmail   string `json:"rootUserEmail"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
