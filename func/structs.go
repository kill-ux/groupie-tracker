package groupie

type Concert struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type DataConcertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DataLocations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:dates`
}

type Artist struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Locations        string   `json:"locations"`
	ConcertDates     string   `json:"concertDates"`
	Relations        string   `json:"relations"`
	DataLocations    DataLocations
	DataConcertDates DataConcertDates
	Concerts         Concert
}

type Page struct {
	Code      int
	MsgError  string
	Arts      []Artist
	Art       Artist
	ArtGroups [][]Artist
}
