package mvac

import (
	"context"
	"strings"
)

type UserContext struct {
	Iss         string   `json:"iss"`
	Sub         string   `json:"sub"`
	Aud         []string `json:"aud"`
	Iat         int      `json:"iat"`
	Exp         int      `json:"exp"`
	Scope       string   `json:"scope"`
	Azp         string   `json:"azp"`
	Permissions []string `json:"permissions"`
	Raw         string   `json:"raw"`
}

func WithAuth(ctx context.Context) *UserContext {
	user := ctx.Value("user").(*UserContext)
	user.Raw = ctx.Value("RawToken").(string)
	return user
}

func (u *UserContext) HasPermission(permission string) bool {
	for _, p := range u.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

func (u *UserContext) HasScope(scope string) bool {
	return u.Scope == scope
}

func (u *UserContext) UserId() string {
	splits := strings.Split(u.Sub, "|")[1]
	return splits[:len(splits)-1]
}
