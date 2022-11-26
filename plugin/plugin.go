package plugin

import (
	"fmt"
	"text/template"
)

type FeatureType int

type PostGenerationFeature struct {
	name                   string
	description            string
	post_generation_action func(PostGenerationContext) error
	availability_filter    func() bool
}

type TemplateFeature struct {
	name                string
	description         string
	templates           []template.Template
	get_template_data   func(TemplateContext) (any, error)
	availability_filter func() bool
	entrypoint          func(TemplateContext) EntryPoint
}

type EntryPoint struct {
	IsDir bool
	Path  string
	Line  int
}

type TemplateContext struct {
	ProjectDir string
	GoVersion  string
	Config     Config
}

type PostGenerationContext struct {
	EntryPoint EntryPoint
}

func (t TemplateFeature) String() string {
	return fmt.Sprintf("%s: %s", t.name, t.description)
}

type Plugin struct {
	post  []PostGenerationFeature
	templ []TemplateFeature
}

func (p Plugin) PostGenerationFeatures() []PostGenerationFeature {
	return p.post
}

func (p Plugin) TemplateFeatures() []TemplateFeature {
	return p.templ
}
