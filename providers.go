package oauth2

var (
	ProviderIDMap = map[string]Provider{
		"battle.net": Provider{
			ID:           "battle.net",
			AuthorizeURL: "https://us.battle.net/oauth/authorize",
			TokenURL:     "https://us.battle.net/oauth/token",
			MeURL:        "https://us.battle.net/oauth/userinfo",
			Scope:        "openid",
			NameProp:     "battletag",
		},
		"discord": Provider{
			ID:           "discord",
			AuthorizeURL: "https://discordapp.com/api/oauth2/authorize",
			TokenURL:     "https://discordapp.com/api/oauth2/token",
			MeURL:        "https://discordapp.com/api/users/@me",
			Scope:        "identify",
			NameProp:     "username",
			NamePrefix:   "@",
		},
		"facebook": Provider{
			ID:           "facebook",
			AuthorizeURL: "https://graph.facebook.com/oauth/authorize",
			TokenURL:     "https://graph.facebook.com/oauth/access_token",
			MeURL:        "https://graph.facebook.com/me",
			NameProp:     "name",
		},
		"github": Provider{
			ID:           "github",
			AuthorizeURL: "https://github.com/login/oauth/authorize",
			TokenURL:     "https://github.com/login/oauth/access_token",
			MeURL:        "https://api.github.com/user",
			Scope:        "read:user",
			NameProp:     "login",
			NamePrefix:   "@",
		},
		"gitlab.com": Provider{
			ID:           "gitlab.com",
			AuthorizeURL: "https://gitlab.com/oauth/authorize",
			TokenURL:     "https://gitlab.com/oauth/token",
			MeURL:        "https://gitlab.com/api/v4/user",
			Scope:        "read_user",
			NameProp:     "username",
			NamePrefix:   "@",
		},
		"google": Provider{
			ID:           "google",
			AuthorizeURL: "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL:     "https://www.googleapis.com/oauth2/v4/token",
			MeURL:        "https://www.googleapis.com/oauth2/v1/userinfo?alt=json",
			Scope:        "profile",
			NameProp:     "name",
		},
		"microsoft": Provider{
			ID:           "microsoft",
			AuthorizeURL: "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
			TokenURL:     "https://login.microsoftonline.com/common/oauth2/v2.0/token",
			MeURL:        "https://graph.microsoft.com/v1.0/me/",
			Scope:        "https://graph.microsoft.com/user.read",
			NameProp:     "displayName",
		},
		"reddit": Provider{
			ID:           "reddit",
			AuthorizeURL: "https://old.reddit.com/api/v1/authorize",
			TokenURL:     "https://old.reddit.com/api/v1/access_token",
			MeURL:        "https://oauth.reddit.com/api/v1/me",
			Scope:        "identity",
			NameProp:     "name",
			NamePrefix:   "u/",
		},
	}
)