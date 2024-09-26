package constants

type SubscriptionsType string

const (
	// Videos
	Netflix        SubscriptionsType = "netflix"
	AmazonPrime    SubscriptionsType = "amazon_prime"
	CanalPlus      SubscriptionsType = "canal_plus"
	DisneyPlus     SubscriptionsType = "disney_plus"
	Hulu           SubscriptionsType = "hulu"
	HBO            SubscriptionsType = "hbo"
	MAX            SubscriptionsType = "max"
	YouTubePremium SubscriptionsType = "youtube_premium"
	AppleTV        SubscriptionsType = "apple_tv"
	ParamountPlus  SubscriptionsType = "paramount_plus"

	// Musics
	Deezer       SubscriptionsType = "deezer"
	Spotify      SubscriptionsType = "spotify"
	AppleMusic   SubscriptionsType = "apple_music"
	Tidal        SubscriptionsType = "tidal"
	AmazonMusic  SubscriptionsType = "amazon_music"
	YouTubeMusic SubscriptionsType = "youtube_music"

	// Software & Tools
	Office365   SubscriptionsType = "office_365"
	AdobeCC     SubscriptionsType = "adobe_cc"
	Grammarly   SubscriptionsType = "grammarly"
	Dropbox     SubscriptionsType = "dropbox"
	GoogleDrive SubscriptionsType = "google_drive"
	GitHub      SubscriptionsType = "github"
	GitLab      SubscriptionsType = "gitlab"

	// Fitness & Health
	Fitbit    SubscriptionsType = "fitbit"
	Peloton   SubscriptionsType = "peloton"
	Strava    SubscriptionsType = "strava"
	Headspace SubscriptionsType = "headspace"
	Calm      SubscriptionsType = "calm"

	// News & Magazines
	NewYorkTimes   SubscriptionsType = "new_york_times"
	WashingtonPost SubscriptionsType = "washington_post"
	TheGuardian    SubscriptionsType = "the_guardian"
	LeMonde        SubscriptionsType = "le_monde"
	TimeMagazine   SubscriptionsType = "time_magazine"

	// Gaming
	PlayStationPlus      SubscriptionsType = "playstation_plus"
	XboxGamePass         SubscriptionsType = "xbox_game_pass"
	NintendoSwitchOnline SubscriptionsType = "nintendo_switch_online"
	EAPlay               SubscriptionsType = "ea_play"
	Steam                SubscriptionsType = "steam"

	// Cloud & Hosting
	AmazonAWS    SubscriptionsType = "amazon_aws"
	GoogleCloud  SubscriptionsType = "google_cloud"
	Azure        SubscriptionsType = "azure"
	DigitalOcean SubscriptionsType = "digital_ocean"

	// Miscellaneous
	Patreon  SubscriptionsType = "patreon"
	OnlyFans SubscriptionsType = "onlyfans"
)

func (st SubscriptionsType) String() string {
	return string(st)
}
