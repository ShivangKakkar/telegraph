package telegraph 

type Telegraph struct {
	Account `json:"account"`
	PageList `json:"pageList"`
	Page `json:"page"`
	PageViews `json:"pageViews"`
}

// Account implements the Account type of Telegraph API
type Account struct {
	ShortName string `json:"short_name"`
	AuthorName string `json:"author_name"`
	AuthorURL string `json:"author_url"`
	AccessToken string `json:"access_token"`
	AuthURL string `json:"auth_url"`
	PageCount int64 `json:"page_count"`

}

// PageList implements the PageList type of Telegraph API
type PageList struct {
	TotalCount int64 `json:"total_count"`
	Pages []Page `json:"pages"`

}

// Page implements the Page type of Telegraph API
type Page struct {
	Path string `json:"path"`
	URL string `json:"url"`
	Title string `json:"title"`
	Description string `json:"description"`
	AuthorName string `json:"author_name"`
	AuthorURL string `json:"author_url"`
	ImageURL string `json:"image_url"`
	Content string `json:"content"`
	Views int64 `json:"views"`
	CanEdit bool `json:"can_edit"`

}

// PageViews implements the PageViews type of Telegraph API
type PageViews struct {
	Views int64 `json:"views"`

}
// A way to set defaults to their zero types instead of nil 
// Know a better way to do this?
type AllValueTypes struct {
	ShortName string `json:"short_name"`
	AuthorName string `json:"author_name"`
	AuthorURL string `json:"author_url"`
	AccessToken string `json:"access_token"`
	AuthURL string `json:"auth_url"`
	PageCount int64 `json:"page_count"`
	TotalCount int64 `json:"total_count"`
	Pages []Page `json:"pages"`
	Path string `json:"path"`
	URL string `json:"url"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
	Content string `json:"content"`
	Views int64 `json:"views"`
	CanEdit bool `json:"can_edit"`
}