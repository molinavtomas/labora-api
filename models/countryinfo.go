package models

type CountryInfo struct {
	Name     string
	Timezone string
	Currency string
	Flag     string
}

type CountryResponse []struct {
	Name       Name                `json:"name"`
	Timezones  []string            `json:"timezones"`
	Currencies map[string]struct { //symbol
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Flags Flags `json:"flags"`
}

type Flags struct {
	PNG string `json:"png"`
}

type Name struct {
	Common string `json:"common"`
}
