package settings

type settings struct {
	Server struct {
		Port          string
		Prefix        string
		AccessSecret  string
		RefreshSecret string
	}

	Database struct {
		UserName string
		Password string
		Host     string
		Port     int
		Name     string
	}

	Redis struct {
		Host     string
		Port     int
		Database int
	}

	Stripe struct {
		Token string
	}
}

var Settings settings
