package structs

type App struct {
	Id                int      `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	IconSrc           string   `json:"iconSrc"`
	Url               string   `json:"url"`
	Rating            int      `json:"rating"`
	RatingsAmount     int      `json:"ratingsAmount"`
	Platforms         []string `json:"platforms"`
	ReleaseDate       int64    `json:"releaseDate"`
	ReleaseDateString string   `json:"releaseDateString"`
	Categories        []string `json:"categories"`
	Subcategories     []int    `json:"subcategories"`
}
