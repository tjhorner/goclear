package clear

type Region struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Webname       string   `json:"webname"`
	Abbreviation  string   `json:"abbr"`
	Location      location `json:"location"`
	Timezone      string   `json:"timezone"`
	TimezoneGroup string   `json:"timezone_group"`
	CurrentEvent  Event    `json:"current_event"`
}

type location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Event struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	RegionName       string    `json:"region_name"`
	RegionID         string    `json:"region_id"`
	Webname          string    `json:"webname"`
	Timezone         string    `json:"timezone"`
	Hashtag          string    `json:"hashtag"`
	URLs             eventURLs `json:"urls"`
	StartsAt         int       `json:"starts_at"`
	EndsAt           int       `json:"ends_at"`
	Venue            venue     `json:"venue"`
	Loaners          loaners   `json:"loaners"`
	StripePublicKey  string    `json:"stripe_public_key"`
	Notice           string    `json:"notice"`
	OverflowEvent    *Event    `json:"overflow_event"`
	HasRelatedEvents bool      `json:"has_related_events"`
}

type eventURLs struct {
	Home     string `json:"home"`
	Register string `json:"register"`
}

type venue struct {
	Name        string  `json:"name"`
	Address     address `json:"address"`
	FullAddress string  `json:"full_address"`
}

type address struct {
	Line1   string `json:"line_1"`
	Line2   string `json:"line_2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Postal  string `json:"postal"`
	Country string `json:"country"`
}

type loaners struct {
	Total     int `json:"total"`
	Available int `json:"available"`
}
