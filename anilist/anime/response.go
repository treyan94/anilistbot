package anime

type SearchResponse struct {
	Data `json:"data"`
}

type Data struct {
	Anime `json:"anime"`
}

type Anime struct {
	Results `json:"results"`
}

type Results []Result

type Result struct {
	Description string `json:"description"`
	Title       `json:"title"`
	CoverImage  `json:"coverImage"`
	SiteUrl     string `json:"siteUrl"`
}

type Title struct {
	English       string `json:"english"`
	Native        string `json:"native"`
	Romaji        string `json:"romaji"`
	UserPreferred string `json:"userPreferred"`
}

type CoverImage struct {
	Medium     string `json:"medium"`
	Large      string `json:"large"`
	ExtraLarge string `json:"extraLarge"`
}
