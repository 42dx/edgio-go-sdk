package common

import "time"

type Searchable interface {
	GetKey() string
	GetName() string
	GetSlug() string
}

type Org struct {
	Searchable

	ID   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

type Variable struct {
	Searchable

	ID        string
	Key       string
	Value     string
	Secret    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (v Variable) GetKey() string {
	return v.Key
}

func (v Variable) GetName() string {
	return ""
}

func (v Variable) GetSlug() string {
	return ""
}

type Property struct {
	Searchable

	ID        string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Property) GetKey() string {
	return ""
}

func (p Property) GetName() string {
	return ""
}

func (p Property) GetSlug() string {
	return p.Slug
}

type Env struct {
	Searchable

	ID                       string
	Name                     string
	LegacyAccNumber          string
	DefaultDomainName        string
	DNSDomainName            string
	CanMembersDeploy         bool
	OnlyMaintainersCanDeploy bool
	HTTPRequestLogging       bool
	PciCompliance            bool
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

func (e Env) GetKey() string {
	return ""
}

func (e Env) GetName() string {
	return e.Name
}

func (e Env) GetSlug() string {
	return ""
}
