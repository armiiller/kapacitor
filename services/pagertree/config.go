package pagertree

import (
	"net/url"

	"github.com/pkg/errors"
)

const DefaultPagerTreeAPIURL = "https://api.pagertree.com/incident"

type Config struct {
	// Whether PagerTree integration is enabled.
	Enabled bool `toml:"enabled" override:"enabled"`
	// The PagerTree API URL, should not need to be changed.
	URL string `toml:"url" override:"url"`
	// The PagerTree service key.
	ServiceKey string `toml:"service-key" override:"service-key,redact"`
	// Whether every alert should automatically go to PagerDuty
	Global bool `toml:"global" override:"global"`
}

func NewConfig() Config {
	return Config{
		URL: DefaultPagerTreeAPIURL,
	}
}

func (c Config) Validate() error {
	if c.URL == "" {
		return errors.New("url cannot be empty")
	}
	if _, err := url.Parse(c.URL); err != nil {
		return errors.Wrapf(err, "invalid URL %q", c.URL)
	}
	return nil
}
