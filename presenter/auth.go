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

	LoginRequestStruct struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	LoginResponseStruct struct {
		Token string `json:"token"`
	}

	AuthHeaderStruct struct {
		Authorization string `header:"Authorization"`
	}
)
