package dto

type SendPaymentProof struct {
	ProofURL string `json:"proof_url" validate:"required"`
}
