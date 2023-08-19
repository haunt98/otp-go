package vault

const (
	// Only support totp for now
	EntryTypeTOTP = "totp"

	AlgoSHA1   = "SHA1"
	AlgoSHA256 = "SHA256"
	AlgoSHA512 = "SHA512"
)

type EntryData struct {
	OTPData   any    `json:"otp_data,omitempty"`
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Issuer    string `json:"issuer,omitempty"`
	EntryType string `json:"entry_type,omitempty"`
}

type TOTPData struct {
	Secret string `json:"secret,omitempty"`
	// See Algo const
	Algo string `json:"algo,omitempty"`
	// 6-8
	Digits        int `json:"digits,omitempty"`
	PeriodSeconds int `json:"period,omitempty"`
}

// ID -> Name
type OTPAll map[string]string
