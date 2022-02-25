package telegraph

// CreateAccountOpts is the set of fields for 'Telegraph.CreateAccount'
type CreateAccountOpts struct {
	// Required. Account name, helps users with several accounts remember which they are currently using. Displayed to the user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
	ShortName string `json:"short_name"`
	// Optional. Default author name used when creating new articles.
	AuthorName string `json:"author_name"`
	// Optional. Default profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url"`
}

// CreateAccount using Telegraph API. Returns Account object
func CreateAccount(opts CreateAccountOpts) (acc *Account, err error) {
	r, e := Get("createAccount", opts)
	return &Account{
		ShortName:   r.ShortName,
		AuthorName:  r.AuthorName,
		AuthorURL:   r.AuthorURL,
		AccessToken: r.AccessToken,
		AuthURL:     r.AuthURL,
		PageCount:   r.PageCount,
	}, e
}

// EditAccountInfoOpts is the set of fields for 'Telegraph.EditAccountInfo'
type EditAccountInfoOpts struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// Optional. New account name.
	ShortName string `json:"short_name"`
	// Optional. New default author name used when creating new articles.
	AuthorName string `json:"author_name"`
	// Optional. New default profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url"`
}

// EditAccountInfo using Telegraph API. Returns Account object
func EditAccountInfo(opts EditAccountInfoOpts) (acc *Account, err error) {
	r, e := Get("editAccountInfo", opts)
	return &Account{
		ShortName:   r.ShortName,
		AuthorName:  r.AuthorName,
		AuthorURL:   r.AuthorURL,
		AccessToken: r.AccessToken,
		AuthURL:     r.AuthURL,
		PageCount:   r.PageCount,
	}, e
}

// GetAccountInfoOpts is the set of fields for 'Telegraph.GetAccountInfo'
type GetAccountInfoOpts struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// Optional. List of account fields to return. Available fields: short_name, author_name, author_url, auth_url, page_count.
	Fields []string `json:"fields"`
}

// GetAccountInfo using Telegraph API. Returns Account object
func GetAccountInfo(opts GetAccountInfoOpts) (acc *Account, err error) {
	r, e := Get("getAccountInfo", opts)
	return &Account{
		ShortName:   r.ShortName,
		AuthorName:  r.AuthorName,
		AuthorURL:   r.AuthorURL,
		AccessToken: r.AccessToken,
		AuthURL:     r.AuthURL,
		PageCount:   r.PageCount,
	}, e
}

// RevokeAccessTokenOpts is the set of fields for 'Telegraph.RevokeAccessToken'
type RevokeAccessTokenOpts struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
}

// RevokeAccessToken using Telegraph API. Returns Account object
func RevokeAccessToken(opts RevokeAccessTokenOpts) (acc *Account, err error) {
	r, e := Get("revokeAccessToken", opts)
	return &Account{
		ShortName:   r.ShortName,
		AuthorName:  r.AuthorName,
		AuthorURL:   r.AuthorURL,
		AccessToken: r.AccessToken,
		AuthURL:     r.AuthURL,
		PageCount:   r.PageCount,
	}, e
}

// CreatePageOpts is the set of fields for 'Telegraph.CreatePage'
type CreatePageOpts struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// Required. Page title.
	Title string `json:"title"`
	// Optional. Author name, displayed below the article's title.
	AuthorName string `json:"author_name"`
	// Optional. Profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url"`
	// Required if Content is empty. Content of the page as HTML string.
	HTMLContent string `json:"html_content"`
	// Required. Content of the page.
	Content []Node `json:"content"`
	// Optional. If true, a content field will be returned in the Page object
	ReturnContent bool `json:"return_content"`
}

// CreatePage using Telegraph API. Returns Page object
func CreatePage(opts CreatePageOpts) (page *Page, err error) {
	if opts.Content == nil && opts.HTMLContent != "" {
		opts.Content = HTMLToNode(opts.HTMLContent)
		opts.HTMLContent = ""
	}
	r, e := Post("createPage", opts) // Post request instead of Get
	return &Page{
		Path:        r.Path,
		URL:         r.URL,
		Title:       r.Title,
		Description: r.Description,
		AuthorName:  r.AuthorName,
		AuthorURL:   r.AuthorURL,
		ImageURL:    r.ImageURL,
		Content:     r.Content,
		Views:       r.Views,
		CanEdit:     r.CanEdit,
	}, e
}

// EditPageOpts is the set of fields for 'Telegraph.EditPage'
type EditPageOpts struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// Required. Path to the page.
	Path string `json:"path"`
	// Required. Page title.
	Title string `json:"title"`
	// Required if Content is empty. Content of the page as HTML string.
	HTMLContent string `json:"html_content"`
	// Required if HTMLContent is empty. Content of the page.
	Content []Node `json:"content"`
	// Optional. Author name, displayed below the article's title.
	AuthorName string `json:"author_name"`
	// Optional. Profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url"`
	// Optional. If true, a content field will be returned in the Page object.
	ReturnContent bool `json:"return_content"`
}

// EditPage using Telegraph API. Returns Page object
func EditPage(opts EditPageOpts) (page *Page, err error) {
	if opts.Content == nil && opts.HTMLContent != "" {
		opts.Content = HTMLToNode(opts.HTMLContent)
		opts.HTMLContent = ""
	}
	r, e := Post("editPage", opts) // Post request instead of Get
	return &Page{
		Path:        r.Path,
		URL:         r.URL,
		Title:       r.Title,
		Description: r.Description,
		AuthorName:  r.AuthorName,
		AuthorURL:   r.AuthorURL,
		ImageURL:    r.ImageURL,
		Content:     r.Content,
		Views:       r.Views,
		CanEdit:     r.CanEdit,
	}, e
}

// GetPageOpts is the set of fields for 'Telegraph.GetPage'
type GetPageOpts struct {
	// Required. Path to the Telegraph page (in the format Title-12-31, i.e. everything that comes after http://telegra.ph/).
	Path string `json:"path"`
	// Optional. If true, content field will be returned in Page object.
	ReturnContent bool `json:"return_content"`
}

// GetPage using Telegraph API. Returns Page object
func GetPage(opts GetPageOpts) (page *Page, err error) {
	r, e := Get("getPage", opts)
	return &Page{
		Path:        r.Path,
		URL:         r.URL,
		Title:       r.Title,
		Description: r.Description,
		AuthorName:  r.AuthorName,
		AuthorURL:   r.AuthorURL,
		ImageURL:    r.ImageURL,
		Content:     r.Content,
		Views:       r.Views,
		CanEdit:     r.CanEdit,
	}, e
}

// GetPageListOpts is the set of fields for 'Telegraph.GetPageList'
type GetPageListOpts struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// Optional. Sequential number of the first page to be returned.
	Offset int64 `json:"offset"`
	// Optional. Limits the number of pages to be retrieved.
	Limit int64 `json:"limit"`
}

// GetPageList using Telegraph API. Returns PageList object
func GetPageList(opts GetPageListOpts) (pl *PageList, err error) {
	r, e := Get("getPageList", opts)
	return &PageList{
		TotalCount: r.TotalCount,
		Pages:      r.Pages,
	}, e
}

// GetViewsOpts is the set of fields for 'Telegraph.GetViews'
type GetViewsOpts struct {
	// Required. Path to the Telegraph page (in the format Title-12-31, where 12 is the month and 31 the day the article was first published).
	Path string `json:"path"`
	// Required if month is passed. If passed, the number of page views for the requested year will be returned.
	Year int64 `json:"year"`
	// Required if day is passed. If passed, the number of page views for the requested month will be returned.
	Month int64 `json:"month"`
	// Required if hour is passed. If passed, the number of page views for the requested day will be returned.
	Day int64 `json:"day"`
	// Optional. If passed, the number of page views for the requested hour will be returned.
	Hour int64 `json:"hour"`
}

// GetViews using Telegraph API. Returns PageViews object
func GetViews(opts GetViewsOpts) (views *PageViews, err error) {
	r, e := Get("getViews", opts)
	return &PageViews{
		Views: r.Views,
	}, e
}
