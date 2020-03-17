package zabbix

import "fmt"

type Template struct {
	TemplateID  string `json:"templateid,omitempty"`
	Name        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayName string `json:"name,omitempty"`
}

type Templates []Template

// TemplateGetParams represent the parameters for a `template.get` API call.
//
// See: https://www.zabbix.com/documentation/3.4/manual/api/reference/template/get#parameters
type TemplateGetParams struct {
	GetParameters
	// TODO
}

// GetTemplates queries the Zabbix API for Templates matching the given search
// parameters.
//
// ErrNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
//
// https://www.zabbix.com/documentation/3.4/manual/api/reference/template/get
func (c *Session) GetTemplates(params TemplateGetParams) ([]Template, error) {
	templates := make([]jTemplate, 0)
	err := c.Get("template.get", params, &templates)
	if err != nil {
		return nil, err
	}

	if len(templates) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Events to Go Events
	out := make([]Template, len(templates))
	for i, jtemplate := range templates {
		template, err := jtemplate.Template()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Template %d in response: %v", i, err)
		}

		out[i] = *template
	}

	return out, nil
}