package telegraph

import (
	"net/url"
	"strconv"
	"strings"
)

// CreateAccountOpts is the set of fields for 'Telegraph.CreateAccount'
type CreateAccountOpts struct {
	// Required. Account name, helps users with several accounts remember which they are currently using. Displayed to the user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
	ShortName string
	// Optional. Default author name used when creating new articles.
	AuthorName string
	// Optional. Default profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorURL string
}

// CreateAccount using Telegraph API. Returns Account object
func CreateAccount(opts *CreateAccountOpts) (acc *Account, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.ShortName) {
		v.Add("short_name", opts.ShortName)
	}
	if !isZeroOfType(opts.AuthorName) {
		v.Add("author_name", opts.AuthorName)
	}
	if !isZeroOfType(opts.AuthorURL) {
		v.Add("author_url", opts.AuthorURL)
	}
	r, e := callAPI("createAccount", v.Encode())
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
	AccessToken string
	// Optional. New account name.
	ShortName string
	// Optional. New default author name used when creating new articles.
	AuthorName string
	// Optional. New default profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorURL string
}

// EditAccountInfo using Telegraph API. Returns Account object
func EditAccountInfo(opts *EditAccountInfoOpts) (acc *Account, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.AccessToken) {
		v.Add("access_token", opts.AccessToken)
	}
	if !isZeroOfType(opts.ShortName) {
		v.Add("short_name", opts.ShortName)
	}
	if !isZeroOfType(opts.AuthorName) {
		v.Add("author_name", opts.AuthorName)
	}
	if !isZeroOfType(opts.AuthorURL) {
		v.Add("author_url", opts.AuthorURL)
	}
	r, e := callAPI("editAccountInfo", v.Encode())
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
	AccessToken string
	// Optional. List of account fields to return. Available fields: short_name, author_name, author_url, auth_url, page_count.
	Fields []string
}

// GetAccountInfo using Telegraph API. Returns Account object
func GetAccountInfo(opts *GetAccountInfoOpts) (acc *Account, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.AccessToken) {
		v.Add("access_token", opts.AccessToken)
	}
	if !isZeroOfType(opts.Fields) {
		v.Add("fields", strings.Join(opts.Fields, ","))
	}
	r, e := callAPI("getAccountInfo", v.Encode())
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
	AccessToken string
}

// RevokeAccessToken using Telegraph API. Returns Account object
func RevokeAccessToken(opts *RevokeAccessTokenOpts) (acc *Account, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.AccessToken) {
		v.Add("access_token", opts.AccessToken)
	}
	r, e := callAPI("revokeAccessToken", v.Encode())
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
	AccessToken string
	// Required. Page title.
	Title string
	// Optional. Author name, displayed below the article's title.
	AuthorName string
	// Optional. Profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorURL string
	// Required. Content of the page.
	Content string
	// Optional. If true, a content field will be returned in the Page object (see: Content format).
	ReturnContent bool
}

// CreatePage using Telegraph API. Returns Page object
func CreatePage(opts *CreatePageOpts) (page *Page, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.AccessToken) {
		v.Add("access_token", opts.AccessToken)
	}
	if !isZeroOfType(opts.Title) {
		v.Add("title", opts.Title)
	}
	if !isZeroOfType(opts.AuthorName) {
		v.Add("author_name", opts.AuthorName)
	}
	if !isZeroOfType(opts.AuthorURL) {
		v.Add("author_url", opts.AuthorURL)
	}
	if !isZeroOfType(opts.Content) {
		v.Add("content", HTMLToNodeString(opts.Content))
	}
	if !isZeroOfType(opts.ReturnContent) {
		v.Add("return_content", strconv.FormatBool(opts.ReturnContent))
	}
	r, e := callAPI("createPage", v.Encode())
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
	AccessToken string
	// Required. Path to the page.
	Path string
	// Required. Page title.
	Title string
	// Required. Content of the page.
	Content string
	// Optional. Author name, displayed below the article's title.
	AuthorName string
	// Optional. Profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorURL string
	// Optional. If�true, a content field will be returned in the Page object.
	ReturnContent bool
}

// EditPage using Telegraph API. Returns Page object
func EditPage(opts *EditPageOpts) (page *Page, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.AccessToken) {
		v.Add("access_token", opts.AccessToken)
	}
	if !isZeroOfType(opts.Path) {
		v.Add("path", opts.Path)
	}
	if !isZeroOfType(opts.Title) {
		v.Add("title", opts.Title)
	}
	if !isZeroOfType(opts.Content) {
		v.Add("content", HTMLToNodeString(opts.Content))
	}
	if !isZeroOfType(opts.AuthorName) {
		v.Add("author_name", opts.AuthorName)
	}
	if !isZeroOfType(opts.AuthorURL) {
		v.Add("author_url", opts.AuthorURL)
	}
	if !isZeroOfType(opts.ReturnContent) {
		v.Add("return_content", strconv.FormatBool(opts.ReturnContent))
	}
	r, e := callAPI("editPage", v.Encode())
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
	Path string
	// Optional. If�true, content field will be returned in Page object.
	ReturnContent bool
}

// GetPage using Telegraph API. Returns Page object
func GetPage(opts *GetPageOpts) (page *Page, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.Path) {
		v.Add("path", opts.Path)
	}
	if !isZeroOfType(opts.ReturnContent) {
		v.Add("return_content", strconv.FormatBool(opts.ReturnContent))
	}
	r, e := callAPI("getPage", v.Encode())
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
	AccessToken string
	// Optional. Sequential number of the first page to be returned.
	Offset int64
	// Optional. Limits the number of pages to be retrieved.
	Limit int64
}

// GetPageList using Telegraph API. Returns PageList object
func GetPageList(opts *GetPageListOpts) (pl *PageList, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.AccessToken) {
		v.Add("access_token", opts.AccessToken)
	}
	if !isZeroOfType(opts.Offset) {
		v.Add("offset", strconv.FormatInt(opts.Offset, 10))
	}
	if !isZeroOfType(opts.Limit) {
		v.Add("limit", strconv.FormatInt(opts.Limit, 10))
	}
	r, e := callAPI("getPageList", v.Encode())
	return &PageList{
		TotalCount: r.TotalCount,
		Pages:      r.Pages,
	}, e
}

// GetViewsOpts is the set of fields for 'Telegraph.GetViews'
type GetViewsOpts struct {
	// Required. Path to the Telegraph page (in the format Title-12-31, where 12 is the month and 31 the day the article was first published).
	Path string
	// Required if month is passed. If passed, the number of page views for the requested year will be returned.
	Year int64
	// Required if day is passed. If passed, the number of page views for the requested month will be returned.
	Month int64
	// Required if hour is passed. If passed, the number of page views for the requested day will be returned.
	Day int64
	// Optional. If passed, the number of page views for the requested hour will be returned.
	Hour int64
}

// GetViews using Telegraph API. Returns PageViews object
func GetViews(opts *GetViewsOpts) (views *PageViews, err error) {
	v := url.Values{}
	if !isZeroOfType(opts.Path) {
		v.Add("path", opts.Path)
	}
	if !isZeroOfType(opts.Year) {
		v.Add("year", strconv.FormatInt(opts.Year, 10))
	}
	if !isZeroOfType(opts.Month) {
		v.Add("month", strconv.FormatInt(opts.Month, 10))
	}
	if !isZeroOfType(opts.Day) {
		v.Add("day", strconv.FormatInt(opts.Day, 10))
	}
	if !isZeroOfType(opts.Hour) {
		v.Add("hour", strconv.FormatInt(opts.Hour, 10))
	}
	r, e := callAPI("getViews", v.Encode())
	return &PageViews{
		Views: r.Views,
	}, e
}
