package mvac

type UserContext struct {
	Iss         string   `json:"iss"`
	Sub         string   `json:"sub"`
	Aud         []string `json:"aud"`
	Iat         int      `json:"iat"`
	Exp         int      `json:"exp"`
	Scope       string   `json:"scope"`
	Azp         string   `json:"azp"`
	Permissions []string `json:"permissions"`
}
