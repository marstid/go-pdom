package gopingdom

type ChecksResult struct {
	Checks Checks `json:"checks"`
	Counts struct {
		Total    int `json:"total"`
		Limited  int `json:"limited"`
		Filtered int `json:"filtered"`
	} `json:"counts"`
}

type CheckResult struct {
	Checks Check `json:"check"`
}

type Check struct {
	Type struct {
		HTTP struct {
			Username         string `json:"username"`
			Password         string `json:"password"`
			URL              string `json:"url"`
			Encryption       bool   `json:"encryption"`
			Port             int    `json:"port"`
			Shouldcontain    string `json:"shouldcontain"`
			Shouldnotcontain string `json:"shouldnotcontain"`
			Postdata         string `json:"postdata"`
			Requestheaders   struct {
			} `json:"requestheaders"`
		} `json:"http"`
	} `json:"type"`
	Sendnotificationwhendown int      `json:"sendnotificationwhendown"`
	Notifyagainevery         int      `json:"notifyagainevery"`
	Notifywhenbackup         bool     `json:"notifywhenbackup"`
	ResponsetimeThreshold    int      `json:"responsetime_threshold"`
	CustomMessage            string   `json:"custom_message"`
	Integrationids           []int    `json:"integrationids"`
	ID                       int      `json:"id"`
	Name                     string   `json:"name"`
	Lasterrortime            int      `json:"lasterrortime"`
	Lasttesttime             int      `json:"lasttesttime"`
	Lastresponsetime         int      `json:"lastresponsetime"`
	Status                   string   `json:"status"`
	Resolution               int      `json:"resolution"`
	Hostname                 string   `json:"hostname"`
	Created                  int      `json:"created"`
	Tags                     []Tag    `json:"tags"`
	ProbeFilters             []string `json:"probe_filters"`
	Ipv6                     bool     `json:"ipv6"`
	VerifyCertificate        bool     `json:"verify_certificate"`
	SslDownDaysBefore        int      `json:"ssl_down_days_before"`
}

// Make check slice sortable
type Checks []Check

func (e Checks) Len() int {
	return len(e)
}

func (e Checks) Less(i, j int) bool {
	return e[i].Name > e[j].Name
}

func (e Checks) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Check1 struct {
	ID                int      `json:"id"`
	Created           int      `json:"created"`
	Name              string   `json:"name"`
	Hostname          string   `json:"hostname"`
	Resolution        int      `json:"resolution"`
	Type              string   `json:"type"`
	Ipv6              bool     `json:"ipv6"`
	VerifyCertificate bool     `json:"verify_certificate"`
	Lasterrortime     int      `json:"lasterrortime,omitempty"`
	Lasttesttime      int      `json:"lasttesttime"`
	Lastresponsetime  int      `json:"lastresponsetime"`
	Status            string   `json:"status"`
	Tags              []Tag    `json:"tags"`
	Maintenanceids    []string `json:"maintenanceids,omitempty"`
}

type Tag struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type ProbesResult struct {
	Probes []Probe `json:"probes"`
}

type Probe struct {
	ID         int    `json:"id"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Name       string `json:"name"`
	Active     bool   `json:"active"`
	Hostname   string `json:"hostname"`
	IP         string `json:"ip"`
	Countryiso string `json:"countryiso"`
	Region     string `json:"region"`
	Ipv6       string `json:"ipv6,omitempty"`
}

type WebHookPost struct {
	CheckID     int    `json:"check_id"`
	CheckName   string `json:"check_name"`
	CheckType   string `json:"check_type"`
	CheckParams struct {
		BasicAuth  bool   `json:"basic_auth"`
		Encryption bool   `json:"encryption"`
		FullURL    string `json:"full_url"`
		Header     string `json:"header"`
		Hostname   string `json:"hostname"`
		Ipv6       bool   `json:"ipv6"`
		Port       int    `json:"port"`
		URL        string `json:"url"`
	} `json:"check_params"`
	Tags                  []string `json:"tags"`
	PreviousState         string   `json:"previous_state"`
	CurrentState          string   `json:"current_state"`
	ImportanceLevel       string   `json:"importance_level"`
	StateChangedTimestamp int      `json:"state_changed_timestamp"`
	StateChangedUtcTime   string   `json:"state_changed_utc_time"`
	LongDescription       string   `json:"long_description"`
	Description           string   `json:"description"`
	FirstProbe            struct {
		IP       string `json:"ip"`
		Ipv6     string `json:"ipv6"`
		Location string `json:"location"`
	} `json:"first_probe"`
	SecondProbe struct {
		IP       string `json:"ip"`
		Ipv6     string `json:"ipv6"`
		Location string `json:"location"`
		Version  int    `json:"version"`
	} `json:"second_probe"`
}
