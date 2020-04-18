package data

import (
	"github.com/wzije/covid19-collection/domains"
	"time"
)

type ProvinceCasePlain struct {
	Cases       []domains.ProvinceCase
	Confirm     int
	Deaths      int
	Recovered   int
	LastUpdated time.Time
}

type TemanggungCasePlain struct {
	Cases       []domains.TemanggengCase
	Confirm     int
	Deaths      int
	Recovered   int
	ODP         int
	PDP         int
	LastUpdated time.Time
}
