package models

const (
	// DbSettingsDefaultName default database name
	DbSettingsDefaultName = "test_backend"

	// DbSettingsDefaultAddress default address of database
	DbSettingsDefaultAddress = "127.0.0.1"

	// DbSettingsDefaultPort default port of database
	DbSettingsDefaultPort = "3306"
)

// DBSettings could be defined in configs/config.json
type DBSettings struct {
	Name     string
	User     string
	Password string
	Address  string
	Port     string
}

// CorsSettings ...
type CorsSettings struct {
	AllowOrigins []string
}

// Config contains all the other configs
type Config struct {
	CorsSettings CorsSettings
	Environment  string
	DBSettings   DBSettings
}

// SetDefaults could set default value in the Config struct
func (o *Config) SetDefaults() {
	if o.DBSettings.Name == "" {
		o.DBSettings.Name = DbSettingsDefaultName
	}
	if o.DBSettings.Address == "" {
		o.DBSettings.Address = DbSettingsDefaultAddress
	}
	if o.DBSettings.Port == "" {
		o.DBSettings.Port = DbSettingsDefaultPort
	}
	if o.Environment == "" {
		o.Environment = "production"
	}
}
