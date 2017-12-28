package jira

import "fmt"

// ProjectService handles projects for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project
type FieldService struct {
	client *Client
}

// FieldSchema field schema
type FieldSchema struct {
	Type   string `json:"type" structs:"type,omitempty"`
	System string `json:"system" structs:"system,omitempty"`
}

// Field represents a single Field
type Field struct {
	Self   string      `json:"self" structs:"self,omitempty"`
	ID     string      `json:"id" structs:"id,omitempty"`
	Name   string      `json:"name" structs:"name,omitempty"`
	Custom bool        `json:"custom" structs:"custom,omitempty"`
	Schema FieldSchema `json:"schema" structs:"schema,omitempty"`
}

// Value
type Values struct {
	ID    int    `json:"id" structs:"id,omitempty"`
	Value string `json:"value" structs:"value,omitempty"`
}

// FOption
type FOption struct {
	Self   string   `json:"self" structs:"self,omitempty"`
	ID     string   `json:"id" structs:"id,omitempty"`
	Name   string   `json:"name" structs:"name,omitempty"`
	Values []Values `json:"values" structs:"values,omitempty"`
}

// GetList gets all projects form JIRA
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project-getAllProjects
func (s *FieldService) GetList() ([]Field, error) {
	apiEndpoint := "rest/api/2/field"

	fields := []Field{}
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return fields, err
	}

	resp, err := s.client.Do(req, &fields)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, jerr
	}

	return fields, nil
}

// GetOption get options for field
func (s *FieldService) GetOption(fieldID string) (FOption, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/field/%s/option", fieldID)

	options := FOption{}
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return options, err
	}

	resp, err := s.client.Do(req, &options)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return options, jerr
	}

	return options, nil
}
