package zabbix

import (
	"fmt"
)

// jTemplate is a private map for the Zabbix API Template object.
// See: https://www.zabbix.com/documentation/3.4/manual/api/reference/template/object
type jTemplate struct {
	TemplateID  string `json:"templateid"`
	Name        string `json:"host"`
	Description string `json:"description,omitempty"`
	DisplayName string `json:"name,omitempty"`
}

// Template returns a native Go Template struct mapped from the given JSON Template data.
func (c *jTemplate) Template() (*Template, error) {
	//var err error

	template := &Template{}
	template.TemplateID = c.TemplateID
	template.Name = c.Name
	template.DisplayName = c.DisplayName
	template.Description   = c.Description

	return template, nil
}

// jTemplates is a slice of jTemplate structs.
type jTemplates []jTemplate

// Templates returns a native Go slice of Templates mapped from the given JSON Templates
// data.
func (c jTemplates) Templates() ([]Template, error) {
	if c != nil {
		templates := make([]Template, len(c))
		for i, jtemplate := range c {
			template, err := jtemplate.Template()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling Template %d in JSON data: %v", i, err)
			}

			templates[i] = *template
		}

		return templates, nil
	}

	return nil, nil
}
