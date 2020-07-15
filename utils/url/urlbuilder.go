package url

// Builder builds a URL
type Builder struct {
	baseURL string
	path    string
	query   map[string]string
}

// New Returns a URLBuilder instance
func New(baseURL string) Builder {
	return Builder{
		baseURL: baseURL,
		path:    "",
		query:   make(map[string]string),
	}
}

// AddPath sets the url path
func (builder Builder) AddPath(path string) Builder {
	builder.path = path
	return builder
}

// AddQueryParam adds a query param to the url
func (builder Builder) AddQueryParam(key string, value string) Builder {
	builder.query[key] = value
	return builder
}

// GetURL returns a string of the full url
func (builder Builder) GetURL() string {
	url := builder.baseURL + builder.path
	if len(builder.query) == 0 {
		return builder.path
	}
	delimiter := "?"
	for key, element := range builder.query {
		url += delimiter + key + "=" + element
		delimiter = "&"
	}
	return url
}
