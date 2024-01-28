package godrinth

type Project struct {
	Slug                 string        `json:"slug"`
	Title                string        `json:"title"`
	Description          string        `json:"description"`
	Categories           []string      `json:"categories"`
	ClientSide           string        `json:"client_side"`
	ServerSide           string        `json:"server_side"`
	Body                 string        `json:"body"`
	Status               string        `json:"status"`
	RequestedStatus      string        `json:"requested_status"`
	AdditionalCategories []string      `json:"additional_categories"`
	IssuesUrl            string        `json:"issues_url"`
	SourceUrl            string        `json:"source_url"`
	WikiUrl              string        `json:"wiki_url"`
	DiscordUrl           string        `json:"discord_url"`
	DonationUrls         []DonationUrl `json:"donation_urls"`
	ProjectType          string        `json:"project_type"`
	Downloads            int           `json:"downloads"`
	IconUrl              string        `json:"icon_url"`
	Color                int           `json:"color"`
	ThreadId             string        `json:"thread_id"`
	MonetizationStatus   string        `json:"monetization_status"`
	Id                   string        `json:"id"`
	Team                 string        `json:"team"`
	Published            string        `json:"published"`
	Updated              string        `json:"updated"`
	Approved             string        `json:"approved"`
	Queued               string        `json:"queued"`
	Followers            int           `json:"followers"`
	License              Licence       `json:"license"`
	Versions             []string      `json:"versions"`
	GameVersions         []string      `json:"game_versions"`
	Loaders              []string      `json:"loaders"`
	Gallery              []Gallery     `json:"gallery"`
}

type Licence struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type DonationUrl struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Gallery struct {
	Url         string `json:"url"`
	Featured    bool   `json:"featured"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Ordering    int    `json:"ordering"`
}

type Meta struct {
	About         string `json:"about"`
	Documentation string `json:"documentation"`
	Name          string `json:"name"`
	Version       string `json:"version"`
}

type SearchResult struct {
	Slug               string   `json:"slug"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Categories         []string `json:"categories"`
	ClientSide         string   `json:"client_side"`
	ServerSide         string   `json:"server_side"`
	ProjectType        string   `json:"project_type"`
	Downloads          int      `json:"downloads"`
	IconUrl            string   `json:"icon_url"`
	Color              int      `json:"color"`
	ThreadId           string   `json:"thread_id"`
	MonetizationStatus string   `json:"monetization_status"`
	ProjectId          string   `json:"project_id"`
	Author             string   `json:"author"`
	DisplayCategories  []string `json:"display_categories"`
	Versions           []string `json:"versions"`
	Follows            int      `json:"follows"`
	DateCreated        string   `json:"date_created"`
	DateModified       string   `json:"date_modified"`
	LatestVersion      string   `json:"latest_version"`
	License            string   `json:"license"`
	Gallery            []string `json:"gallery"`
	FeaturedGallery    string   `json:"featured_gallery"`
}

type SearchResults struct {
	Hits      []SearchResult `json:"hits"`
	Offset    int            `json:"offset"`
	Limit     int            `json:"limit"`
	TotalHits int            `json:"total_hits"`
}
