package bootstrap

type Env struct {
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiry  int    `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry int    `mapstructure:"REFRESH_TOKEN_EXPIRY"`
	ContextTimeout     int    `mapstructure:"CONTEXT_TIMEOUT"`
}
