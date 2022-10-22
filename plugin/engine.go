package plugin

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type Engine struct {
	plugins []Plugin
}

func NewEngine(plugins ...Plugin) *Engine {
	return &Engine{plugins: plugins}
}

func (e *Engine) Run(dir string) error {

	//templates
	available_tf := make([]TemplateFeature, 0)

	for _, p := range e.plugins {
		for _, tf := range p.TemplateFeatures() {
			if tf.availability_filter() {
				available_tf = append(available_tf, tf)
			}
		}
	}

	prompt := promptui.Select{
		Label: "Template",
		Items: available_tf,
	}
	idx, _, err := prompt.Run()

	if err != nil {
		return err
	}

	tf := available_tf[idx]

	data, err := tf.get_template_data()
	if err != nil {
		return err
	}

	err = os.Mkdir(dir, 0775)

	if err != nil {
		return err
	}

	err = os.Chdir(dir)

	if err != nil {
		return err
	}

	for _, templ := range tf.templates {

		if _, err := os.Stat(templ.Name()); err == nil {
			return fmt.Errorf("file %s already exists", templ.Name())
		}
		f, err := os.Create(templ.Name())

		if err != nil {
			return err
		}
		err = templ.Execute(f, data)
		if err != nil {
			return err
		}

	}

	// post generation
	for _, p := range e.plugins {
		for _, pgf := range p.PostGenerationFeatures() {
			if pgf.availability_filter() {
				if err := pgf.post_generation_action(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
