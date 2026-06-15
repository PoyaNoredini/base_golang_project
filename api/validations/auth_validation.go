package validations


type LoginWithPasswordRequest struct {
    Phone    string `json:"phone"    binding:"required,len=11"`
    Password string `json:"password" binding:"required,min=6"`
}

type LoginWithOtpRequest struct {
    Phone string `json:"phone" binding:"required,len=11"`
    Code  string `json:"code"  binding:"required,len=6"`
}

// type RegisterRequest struct {
//     FirstName   string `json:"first_name"   binding:"required,min=2,max=50"`
//     LastName    string `json:"last_name"    binding:"required,min=2,max=50"`
//     Phone       string `json:"phone"        binding:"required,len=11"`
//     NationalID  string `json:"national_id"  binding:"required,len=10"`
//     Password    string `json:"password"     binding:"required,min=6"`
//     Email       string `json:"email"        binding:"omitempty,email"` // optional
// }