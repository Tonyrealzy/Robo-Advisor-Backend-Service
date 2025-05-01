package models

type SignupResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"User created successfully"`
}

type ErrorResponse struct {
	Status string `json:"status" example:"error"`
	Error  string `json:"error" example:"Something went wrong"`
}

type ServerErrorResponse struct {
	Status string `json:"status" example:"error"`
	Error  string `json:"error" example:"Internal Server Error"`
}

type AuthErrorResponse struct {
	Status string `json:"status" example:"error"`
	Error  string `json:"error" example:"Invalid or expired token"`
}

type LoginResponse struct {
	Status string `json:"status" example:"success"`
	Token  string `json:"token" example:"jwt.token.here"`
}

type LogoutResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Action successful"`
}

type PasswordResetResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Password reset successful"`
}

type ConfirmSignupResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"User status updated successfully"`
}

type PasswordChangeResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Password changed successfully"`
}

type InvestmentAdvice struct {
	InvestmentAdvice string `json:"investmentAdvice" example:"Based on your risk tolerance, we recommend a diversified portfolio of stocks and bonds."`
}

type AIResponse struct {
	Status string           `json:"status" example:"success"`
	Data   InvestmentAdvice `json:"data,omitempty"`
}
