package domains

import "time"

//model/data
type Case struct {
	ID        int       `json:"id,omitempty"`
	Province  string    `json:"province"`
	Confirmed int       `json:"confirm"`
	Recovered int       `json:"recovered"`
	Deaths    int       `json:"death"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (ci Case) CreatedAtID() time.Time {
	unixDate := ci.CreatedAt.Unix()
	return time.Unix(unixDate, 0)
}

func (ci Case) UpdatedAtID() time.Time {
	unixDate := ci.UpdatedAt.Unix()
	return time.Unix(unixDate, 0)
}

type CaseInfo struct {
	LastUpdate time.Time
}

func (ci CaseInfo) LastDateID() time.Time {
	unixDate := ci.LastUpdate.Unix()
	return time.Unix(unixDate, 0)
}
