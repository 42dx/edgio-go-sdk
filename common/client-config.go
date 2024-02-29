package common

type ClientConfig struct {
	Url        string
	ApiVersion string
	Service    string
	Scope      string
	OrgId      string
}

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
