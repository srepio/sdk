package types

type MFAType string

const (
	TOTP = MFAType("totp")
)

type MFAData struct {
	Secret        string    `json:"secret"`
	URL           string    `json:"url"`
	RecoveryCodes [6]string `json:"recovery_codes"`
}
