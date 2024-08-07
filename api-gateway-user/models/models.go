package models

type RegisterReqSwag struct {
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password
}

type RegisterReq struct {
	ID       string `json:"id"`       // User's unique identifier
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password
	Role     string
}

type LoginReq struct {
	Email    string `json:"email"`    // User's email
	Password string `json:"password"` // User's password
}

type GetProfileReq struct {
	Email string `json:"email"` // Username of the profile to retrieve
}

type GetProfileResp struct {
	ID          string `json:"id"`       // User's unique identifier
	Email       string `json:"email"`    // User's email address
	Password    string `json:"password"` // User's password
	Role        string `json:"role"`
	IsConfirmed bool   `json:"is_confirmed"` // Add IsConfirmed to the model
}

type GetProfileByIdReq struct {
	ID string `json:"id"` // Username of the profile to retrieve
}

type GetProfileByIdResp struct {
	ID    string `json:"id"`    // User's unique identifier
	Email string `json:"email"` // User's email address
	Role  string `json:"role"`
}

type ForgotPasswordReq struct {
	Email string `json:"email"` // User's email address
}

type ResetPasswordReq struct {
	Email       string `json:"email"`        // User's email address
	NewPassword string `json:"new_password"` // User's new password
}

type RecoverPasswordReq struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

type UpdatePasswordReq struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

type ConfirmUserReq struct {
	Email string `json:"email"`
}

type ConfirmRegistrationReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type MongoUser struct {
	ID   string `bson:"user_id"`
	Cart []Cart `bson:"cart"`
}

type Cart struct {
	ProductID string `bson:"product_id"`
	Quantity  int    `bson:"quantity"`
}
