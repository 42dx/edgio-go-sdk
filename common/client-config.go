package common

type ClientConfig struct {
	Url        string
	ApiVersion string
	Service    string
	Scope      string
	OrgId      string
}

// Merge Merges the `other` ClientConfig into the default one,
// overwriting the current values with the other's if they are not empty.
// MAndatory params:
// other common.ClientConfig
// Even though `other` is mandatory, no keys are mandatory in it. If no keys are provided,
// the default value is returned
// Returns the merged ClientConfig.
func (c ClientConfig) Merge(other ClientConfig) ClientConfig {
	if other.Url != "" {
		c.Url = other.Url
	}

	if other.ApiVersion != "" {
		c.ApiVersion = other.ApiVersion
	}

	if other.Service != "" {
		c.Service = other.Service
	}

	if other.Scope != "" {
		c.Scope = other.Scope
	}

	if other.OrgId != "" {
		c.OrgId = other.OrgId
	}

	return c
}
