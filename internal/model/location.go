package model

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next,omitempty"`
	Previous string     `json:"previous,omitempty"`
	Results  []location `json:"results"`
}

func (l *LocationResponse) GetLocationNames() []string {
	names := make([]string, 0, len(l.Results))
	for _, loc := range l.Results {
		names = append(names, loc.Name+"-area")
	}
	return names

}
