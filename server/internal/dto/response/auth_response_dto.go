package response

type LoginResponseDTO struct {
	AccessToken string `json:"access_token"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}
