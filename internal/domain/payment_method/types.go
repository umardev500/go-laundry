package paymentmethod

type QRISMetadata struct {
	ImageURL     string `json:"image_url"`
	MerchantName string `json:"merchant_name"`
}

type BankStransferMetadata struct {
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
}

type CashMetadata struct {
	Note *string `json:"note"`
}
