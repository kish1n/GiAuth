package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type EmailRaw struct {
	Password string `fig:"password,required"`
	Email    string `fig:"email,required"`
}

type Email struct {
	Password string
	Address  string
}

func (c *config) Email() *Email {
	return c.email.Do(func() interface{} {
		cfgRaw := EmailRaw{}
		err := figure.
			Out(&cfgRaw).
			From(kv.MustGetStringMap(c.getter, "email")).
			Please()
		if err != nil {
			panic(errors.WithMessage(err, "failed to figure email"))
		}

		return &Email{
			Password: cfgRaw.Password,
			Address:  cfgRaw.Email,
		}
	}).(*Email)
}
