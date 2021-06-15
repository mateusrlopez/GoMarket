package settings

type settings struct {
	Server struct {
		Port   string
		Prefix string
		Secret string
	}

	Database struct {
		UserName string
		Password string
		Host     string
		Port     string
		Name     string
	}
}

var Settings settings
