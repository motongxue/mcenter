package rest

import (
	"github.com/infraboard/mcube/client/rest"
)

func NewClient(conf *Config) *ClientSet {
	c := rest.NewRESTClient()
	c.SetBearerTokenAuth(conf.Token)
	c.SetBaseURL(conf.Address + conf.PathPrefix)
	return &ClientSet{
		c: c,
	}
}

type ClientSet struct {
	c *rest.RESTClient
}

func (c *ClientSet) Application() MetaService {
	return &appImpl{client: c.c}
}

func (c *ClientSet) Instance() InstanceService {
	return &insImpl{client: c.c}
}
