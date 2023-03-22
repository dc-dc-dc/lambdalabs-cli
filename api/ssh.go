package api

type SSHAddRequest struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key,omitempty"`
}
