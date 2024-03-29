package uomcache

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"github.com/rameshrajagopalanbayer/uom-api-go/api/models"
	"log"
	"time"
)

type uomCache struct {
	uoms *cache.Cache
}

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

func NewUomCache() *uomCache {
	Cache := cache.New(defaultExpiration, purgeTime)
	return &uomCache{
		uoms: Cache,
	}
}

func (c *uomCache) read(id string) (item []byte, ok bool) {
	uom, ok := c.uoms.Get(id)
	if ok {
		log.Println("from uomcache")
		res, err := json.Marshal(uom.(models.Uom))
		if err != nil {
			log.Fatal("Error")
		}
		return res, true
	}
	return nil, false
}

func (c *uomCache) update(id string, uom models.Uom) {
	c.uoms.Set(id, uom, cache.DefaultExpiration)
}

func (c *uomCache) updateAll(uoms []models.Uom) {
	for _, uom := range uoms {
		c.uoms.Set(uom.Code, uom, cache.DefaultExpiration)
	}
}

func (c *uomCache) getAll() []models.Uom {
	var uoms []models.Uom
	for _, value := range c.uoms.Items() {
		uoms = append(uoms, value.Object.(models.Uom))
	}
	return uoms
}
