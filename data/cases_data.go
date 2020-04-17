package data

import (
	"github.com/wzije/covid19-collection/domains"
	"time"
)

type CaseProvincePlain struct {
	Cases       []domains.Case
	Confirm     int
	Deaths      int
	Recovered   int
	LastUpdated time.Time
}

type CaseTemanggungPlain struct {
	Cases       []domains.CaseInTemanggung
	Confirm     int
	Deaths      int
	Recovered   int
	ODP         int
	PDP         int
	LastUpdated time.Time
}
