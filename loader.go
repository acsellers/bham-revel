package bhamrevel

import (
	"github.com/acsellers/bham"
	"github.com/acsellers/revel"
)

func init() {
	revel.AddParser(
		[]string{"bham", "bhaml", "haml"},
		bham.Parse,
		map[string]interface{}{},
	)
}
