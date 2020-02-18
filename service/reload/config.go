package reload

import "github.com/spiral/roadrunner/service"

// Config is a Reload configuration point.
type Config struct {
	// Enable or disable Reload extension, default disable.
	Enabled bool

	// Watch is general pattern of files to watch. It will be applied to every directory in project
	Watch []string

	// Services is set of services which would be reloaded in case of FS changes
	Services map[string]ServiceConfig
}

type ServiceConfig struct {
	// Watch is per-service specific files to watch
	Watch []string
	// Dirs is per-service specific dirs which will be combined with Watch
	Dirs  []string
	// Ignore is set of files which would not be watched
	Ignore []string
}


// Hydrate must populate Config values using given Config source. Must return error if Config is not valid.
func (c *Config) Hydrate(cfg service.Config) error {
	if err := cfg.Unmarshal(c); err != nil {
		return err
	}
	return nil
}

// InitDefaults sets missing values to their default values.
func (c *Config) InitDefaults() error {
	//c.Interval = time.Second

	return nil
}
