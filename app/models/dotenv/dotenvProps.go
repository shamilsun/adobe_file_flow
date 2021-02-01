package dotenv

type IDotEnv struct {
	DbUser     string
	DbPassword string
	Database   string
	DbAddr     string

	ClientId     string
	ClientSecret string
	FromName     string
	FromEmail    string

	QuickEmailVerification string

	Mode string
}
