package types

type APIList struct {
	Total int    `json:"total,omitempty"`
	Data  []*API `json:"data,omitempty"`
}

type API struct {
	CreatedAt              int64    `json:"created_at,omitempty"`
	Hosts                  []string `json:"hosts,omitempty"`
	URIS                   []string `json:"uris,omitempty"`
	Methods                []string `json:"methods,omitempty"`
	HTTPIfTerminated       bool     `json:"http_if_terminated,omitempty"`
	HTTPSOnly              bool     `json:"http_only,omitempty"`
	Id                     string   `json:"id,omitempty"`
	Name                   string   `json:"name,omitempty"`
	PreserveHost           bool     `json:"preserve_host,omitempty"`
	Retries                int      `json:"retries,omitempty"`
	StripURI               bool     `json:"strip_uri,omitempty"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout,omitempty"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout,omitempty"`
	UpstreamURL            string   `json:"upstream_url,omitempty"`
}

type APIListOptions struct {
	Id   string
	Name string
}

type APICreateRequest struct {
	APICreate
	Name string `json:"name,omitempty"`
}

type APIUpdate struct {
	APICreate
	Name string `json:"name,omitempty"`
}

type APICreate struct {
	Hosts                  []string `json:"hosts,omitempty"`
	URIS                   []string `json:"uris,omitempty"`
	Methods                []string `json:"methods,omitempty"`
	HTTPIfTerminated       bool     `json:"http_if_terminated,omitempty"`
	HTTPSOnly              bool     `json:"https_only,omitempty"`
	PreserveHost           bool     `json:"preserve_host,omitempty"`
	Retries                int      `json:"retries,omitempty"`
	StripURI               bool     `json:"strip_uri,omitempty"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout,omitempty"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout,omitempty"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout,omitempty"`
	UpstreamURL            string   `json:"upstream_url,omitempty"`
}

type APIErrorResponse struct {
	Name        string `json:"name,omitempty"`
	UpstreamURL string `json:"upstream_url,omitempty"`
	Message     string `json:"message,omitempty"`
	Hosts       string `json:"hosts,omitempty"`
	Methods     string `json:"methods,omitempty"`
}
