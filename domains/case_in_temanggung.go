package domains

import "time"

//model/data
type CaseInTemanggung struct {
	ID        int       `json:"id,omitempty"`
	Area      string    `json:"area"`
	ODP       int       `json:"id,omitempty"`
	PDP       int       `json:"id,omitempty"`
	Confirmed int       `json:"confirm"`
	Recovered int       `json:"recovered"`
	Deaths    int       `json:"death"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CaseInTemanggungInfo struct {
	LastUpdate time.Time
}

func (ci CaseInTemanggung) CreatedAtID() time.Time {
	unixDate := ci.CreatedAt.Unix()
	return time.Unix(unixDate, 0)
}

func (ci CaseInTemanggung) UpdatedAtID() time.Time {
	unixDate := ci.UpdatedAt.Unix()
	return time.Unix(unixDate, 0)
}
