package settings

type settings struct {
	Server struct {
		Port   string
		Prefix string
	}

	Database struct {
		UserName     string
		Password     string
		Host         string
		Port         string
		DatabaseName string
	}
}

var Settings settings
