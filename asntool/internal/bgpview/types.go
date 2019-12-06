package bgpview

type Asn struct {
	AbuseContacts []string `json:"abuse_contacts"`
	Asn           int      `json:"asn"`
	CountryCode   string   `json:"country_code"`
	Description   string   `json:"description"`
	EmailContacts []string `json:"email_contacts"`
	Name          string   `json:"name"`
	RirName       string   `json:"rir_name"`
}

type IPv4Prefix struct {
	AbuseContacts []string `json:"abuse_contacts"`
	Cidr          int      `json:"cidr"`
	CountryCode   string   `json:"country_code"`
	Description   string   `json:"description"`
	EmailContacts []string `json:"email_contacts"`
	IP            string   `json:"ip"`
	Name          string   `json:"name"`
	ParentCidr    int      `json:"parent_cidr"`
	ParentIP      string   `json:"parent_ip"`
	ParentPrefix  string   `json:"parent_prefix"`
	Prefix        string   `json:"prefix"`
	RirName       string   `json:"rir_name"`
}

type IPv6Prefix struct {
	AbuseContacts []string `json:"abuse_contacts"`
	Cidr          int      `json:"cidr"`
	CountryCode   string   `json:"country_code"`
	Description   string   `json:"description"`
	EmailContacts []string `json:"email_contacts"`
	IP            string   `json:"ip"`
	Name          string   `json:"name"`
	ParentCidr    int      `json:"parent_cidr"`
	ParentIP      string   `json:"parent_ip"`
	ParentPrefix  string   `json:"parent_prefix"`
	Prefix        string   `json:"prefix"`
	RirName       string   `json:"rir_name"`
}

// SearchResults contains search results
type SearchResults struct {
	Data struct {
		Asns              []Asn         `json:"asns"`
		InternetExchanges []interface{} `json:"internet_exchanges"`
		Ipv4Prefixes      []IPv4Prefix  `json:"ipv4_prefixes"`
		IPv6Prefixes      []IPv6Prefix  `json:"ipv6_prefixes"`
	} `json:"data"`
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
}

// ASNPrefixResults contains prefix search results
type ASNPrefixResults struct {
	Data struct {
		IPv4Prefixes []IPv4Prefix `json:"ipv4_prefixes"`
		IPv6Prefixes []IPv6Prefix `json:"ipv6_prefixes"`
	} `json:"data"`
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
}
