package common

type ClientConfig struct {
	URL         string
	APIVersion  string
	Service     string
	Scope       string
	OrgID       string
	AccessToken string
}

// Merge Merges the `other` ClientConfig into the default one,
// overwriting the current values with the other's if they are not empty.
// MAndatory params:
// other common.ClientConfig
// Even though `other` is mandatory, no keys are mandatory in it. If no keys are provided,
// the default value is returned
// Returns the merged ClientConfig.
func (c ClientConfig) Merge(other ClientConfig) ClientConfig {
	if other.URL != "" {
		c.URL = other.URL
	}

	if other.APIVersion != "" {
		c.APIVersion = other.APIVersion
	}

	if other.Service != "" {
		c.Service = other.Service
	}

	if other.Scope != "" {
		c.Scope = other.Scope
	}

	if other.OrgID != "" {
		c.OrgID = other.OrgID
	}

	if other.OrgID != "" {
		c.AccessToken = other.AccessToken
	}

	return c
}
