package plugin

import "text/template"

type PluginBuilder struct {
	post  []PostGenerationFeature
	templ []TemplateFeature
}
type PostGenerationFeatureBuilder struct {
	parent                 *PluginBuilder
	name                   string
	description            string
	post_generation_action func() error
	availability_filter    func() bool
}

type TemplateFeatureBuilder struct {
	parent              *PluginBuilder
	name                string
	description         string
	templates           []template.Template
	get_template_data   func() (any, error)
	availability_filter func() bool
}

func Builder() *PluginBuilder {
	return &PluginBuilder{}
}

func (p *PluginBuilder) Build() Plugin {
	return Plugin{post: p.post, templ: p.templ}
}

//#region post

func (p *PluginBuilder) PostGenerationFeature() *PostGenerationFeatureBuilder {
	return &PostGenerationFeatureBuilder{parent: p}
}

func (f *PostGenerationFeatureBuilder) WithName(name string) *PostGenerationFeatureBuilder {
	f.name = name
	return f
}

func (f *PostGenerationFeatureBuilder) WithDescription(description string) *PostGenerationFeatureBuilder {
	f.description = description
	return f
}

func (f *PostGenerationFeatureBuilder) WithPostGenerationAction(action func() error) *PostGenerationFeatureBuilder {
	f.post_generation_action = action
	return f
}

func (f *PostGenerationFeatureBuilder) WithAvailabilityFilter(filter func() bool) *PostGenerationFeatureBuilder {
	f.availability_filter = filter
	return f
}

func (f *PostGenerationFeatureBuilder) AddFeature() *PluginBuilder {
	f.parent.post = append(f.parent.post, PostGenerationFeature{
		name:                   f.name,
		description:            f.description,
		post_generation_action: f.post_generation_action,
		availability_filter:    f.availability_filter,
	})
	return f.parent
}

//#endregion

//#region template

func (p *PluginBuilder) TemplateFeature() *TemplateFeatureBuilder {
	return &TemplateFeatureBuilder{parent: p}
}

func (f *TemplateFeatureBuilder) WithName(name string) *TemplateFeatureBuilder {
	f.name = name
	return f
}

func (f *TemplateFeatureBuilder) WithDescription(description string) *TemplateFeatureBuilder {
	f.description = description
	return f
}

func (f *TemplateFeatureBuilder) WithGetTemplateData(getter func() (any, error)) *TemplateFeatureBuilder {
	f.get_template_data = getter
	return f
}

func (f *TemplateFeatureBuilder) WithNewTemplate(templ template.Template) *TemplateFeatureBuilder {
	f.templates = append(f.templates, templ)
	return f
}

func (f *TemplateFeatureBuilder) WithAvailabilityFilter(filter func() bool) *TemplateFeatureBuilder {
	f.availability_filter = filter
	return f
}

func (f *TemplateFeatureBuilder) AddFeature() *PluginBuilder {
	f.parent.templ = append(f.parent.templ, TemplateFeature{
		name:                f.name,
		description:         f.description,
		templates:           f.templates,
		get_template_data:   f.get_template_data,
		availability_filter: f.availability_filter,
	})
	return f.parent
}

//#endregion
