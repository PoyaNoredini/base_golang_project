package validations

type UpdateRequest struct{
    FirstName   string `json:"first_name"   binding:"omitempty,min=2,max=50"`
    LastName    string `json:"last_name"    binding:"omitempty,min=2,max=50"`
    Phone       string `json:"phone"        binding:"omitempty,len=11"`
    NationalID  string `json:"national_id"  binding:"omitempty,len=10"`
    Email       string `json:"email"        binding:"omitempty,email"` // optional
}

type AdminUpdateRequest struct{
	FirstName   string `json:"first_name"   binding:"omitempty,min=2,max=50"`
    LastName    string `json:"last_name"    binding:"omitempty,min=2,max=50"`
    Phone       string `json:"phone"        binding:"omitempty,len=11"`
    NationalID  string `json:"national_id"  binding:"omitempty,len=10"`
    Email       string `json:"email"        binding:"omitempty,email"` // optional
}