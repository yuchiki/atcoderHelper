package config

import (
	"errors"
	"fmt"
)

var (
	// GlobalAppConfig is a config held globally in a program.
	GlobalAppConfig AppConfig

	// GlobalTaskConfig is a task config held globally in a program.
	GlobalTaskConfig TaskConfig
)

var (
	// ErrTemplateNotFound is an error raised when template.
	ErrTemplateNotFound = errors.New("template not found")

	// ErrLanguageNotFound is an error
	ErrLanguageNotFound = errors.New("language Not found")
)

// AppConfig is a config of ach command.
type AppConfig struct {
	Username        string
	Languages       []Language
	Templates       []Template
	DefaultTemplate string `mapstructure:"default-template"`
}

// Language is an information of a programming language used when solving a task.
type Language struct {
	Name        string
	AtCoderName string `mapstructure:"atcoder-name"`
	Build       string
	Run         string
}

// Template is a template used when new task is created.
type Template struct {
	Name              string
	Language          string
	TemplateDirectory string `mapstructure:"template-directory"`
	SourceFile        string `mapstructure:"source-file"`
}

// TaskConfig is a configuration for each task.
type TaskConfig struct {
	ContestID string `mapstructure:"contest-id"`
	TaskID    string `mapstructure:"task-id"`
	Template  string
}

// GetLanguage finds a language by name.
func (t *AppConfig) GetLanguage(name string) (Language, error) {
	for _, language := range t.Languages {
		if language.Name == name {
			return language, nil
		}
	}

	return Language{}, fmt.Errorf("language %s not found: %w", name, ErrLanguageNotFound)
}

// GetLanguage finds a language by name.
func GetLanguage() (Language, error) {
	template, err := GetTemplate()
	if err != nil {
		return Language{}, err
	}

	return GlobalAppConfig.GetLanguage(template.Language)
}

// GetTemplate finds a template by name.
func (t *AppConfig) GetTemplate(name string) (Template, error) {
	for _, template := range t.Templates {
		if template.Name == name {
			return template, nil
		}
	}

	return Template{}, fmt.Errorf("template %s not found: %w", name, ErrTemplateNotFound)
}

// GetTemplate finds a template by name.
func GetTemplate() (Template, error) {
	return GlobalAppConfig.GetTemplate(GlobalTaskConfig.Template)
}
