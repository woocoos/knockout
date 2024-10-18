package types

type (
	OrgLogo struct {
		Logo      string `json:"logo,omitempty"`
		ThumbLogo string `json:"thumbLogo,omitempty"`
		Favicon   string `json:"favicon,omitempty"`
	}
)
