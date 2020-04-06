package presenter

type (
	PrivateClaims struct {
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Role      string `json:"role"`
		Timestamp string `json:"timestamp"`
	}

	RegisterRequestStruct struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Role  string `json:"role"`
	}

	RegisterResponseStruct struct {
		Password string `json:"password"`
	}
)
