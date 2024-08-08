package payload

type (
	LoginResponseDTO struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    struct {
			AccessToken string `json:"access_token"`
		}
	}
)
