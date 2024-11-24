package configs

type (
	Config struct {
		Service  Service       `mapstructure:"service"`
		Database Database      `mapstructure:"database"`
		Spotify  SpotifyConfig `mapstructure:"spotifyConfig"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJWT string `mapstructure:"secretJWT"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}

	SpotifyConfig struct {
		ClientID     string `mapstructure:"clientID"`
		ClientSecret string `mapstructure:"clientSecret"`
	}
)
