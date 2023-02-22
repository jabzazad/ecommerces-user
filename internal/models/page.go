package models

const (
	// DefaultSize default size
	DefaultSize = 20

	// DefaultPage default page
	DefaultPage = 1
)

// PageInformation page information
type PageInformation struct {
	Page     int   `json:"page,omitempty"`
	Size     int   `json:"size,omitempty"`
	Count    int64 `json:"count,omitempty"`
	LastPage int   `json:"last_page,omitempty"`
}

// Page page model
type Page struct {
	PageInformation *PageInformation `json:"page_information,omitempty"`
	Entities        interface{}      `json:"entities,omitempty"`
	Message
}

// NewPage new page
func NewPage(pageInfo *PageInformation, entities interface{}) *Page {
	return &Page{
		PageInformation: &PageInformation{
			Page:     pageInfo.Page,
			Size:     pageInfo.Size,
			Count:    pageInfo.Count,
			LastPage: pageInfo.LastPage,
		},
		Entities: entities,
	}
}

// GetEntities get entities
func (p *Page) GetEntities() interface{} {
	return p.Entities
}

// PageForm page form
type PageForm struct {
	Page  int    `json:"page,omitempty" form:"page" query:"page" url:"page"`
	Size  int    `json:"size,omitempty" form:"size" query:"size" url:"size"`
	Query string `json:"query,omitempty" form:"query" query:"query" url:"query"`
}

// GetPage get page
func (f *PageForm) GetPage() int {
	if f.Page < 1 {
		f.Page = DefaultPage
	}
	return f.Page
}

// GetSize get size
func (f *PageForm) GetSize() int {
	if f.Size < 1 {
		f.Size = DefaultSize
	}
	return f.Size
}

// GetQuery get query
func (f *PageForm) GetQuery() string {
	return f.Query
}
