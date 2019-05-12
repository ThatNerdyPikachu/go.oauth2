package oauth2

type Provider struct {
	ID           string
	AuthorizeURL string
	TokenURL     string
	MeURL        string
	Scope        string
	NameProp     string
	NamePrefix   string
}

var (
	ProviderDiscord = Provider{
		"discord",
		"https://discordapp.com/api/oauth2/authorize",
		"https://discordapp.com/api/oauth2/token",
		"https://discordapp.com/api/users/@me",
		"identify",
		"username",
		"@",
	}
	ProviderGitHub = Provider{
		"github",
		"https://github.com/login/oauth/authorize",
		"https://github.com/login/oauth/access_token",
		"https://api.github.com/user",
		"read:user",
		"login",
		"@",
	}
	ProviderGoogle = Provider{
		"google",
		"https://accounts.google.com/o/oauth2/v2/auth",
		"https://www.googleapis.com/oauth2/v4/token",
		"https://www.googleapis.com/oauth2/v1/userinfo?alt=json",
		"profile",
		"name",
		"",
	}
	ProviderReddit = Provider{
		"reddit",
		"https://old.reddit.com/api/v1/authorize",
		"https://old.reddit.com/api/v1/access_token",
		"https://oauth.reddit.com/api/v1/me",
		"identity",
		"name",
		"u/",
	}
)
