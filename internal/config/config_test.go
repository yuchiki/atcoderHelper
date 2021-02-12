package config

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAppConfig_GetLanguage(t *testing.T) {
	config := newSampleAppConfig()

	testcases := []struct {
		name     string
		input    string
		expected Language
		err      error
	}{
		{
			name:     "OK",
			input:    "csharp",
			expected: config.Languages[0],
		},
		{
			name:  "returns error when not found",
			input: "not-existing-language",
			err:   ErrLanguageNotFound,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			config := newSampleAppConfig()
			actual, err := config.GetLanguage(testcase.input)

			if diff := cmp.Diff(actual, testcase.expected); diff != "" {
				t.Errorf("expected %v, but actual %v", testcase.expected, actual)
			}

			if !errors.Is(err, testcase.err) {
				t.Errorf("expected %v, but actual %v", testcase.err, err)
			}
		})
	}
}

func TestAppConfig_GetTemplate(t *testing.T) {
	config := newSampleAppConfig()

	testcases := []struct {
		name     string
		input    string
		expected Template
		err      error
	}{
		{
			name:     "OK",
			input:    "csharp-template",
			expected: config.Templates[0],
		},
		{
			name:  "returns error when template not found",
			input: "not existing template",
			err:   ErrTemplateNotFound,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			config := newSampleAppConfig()
			actual, err := config.GetTemplate(testcase.input)

			if diff := cmp.Diff(actual, testcase.expected); diff != "" {
				t.Errorf("expected %v, but actual %v", testcase.expected, actual)
			}

			if !errors.Is(err, testcase.err) {
				t.Errorf("expected %v, but actual %v", testcase.err, err)
			}
		})
	}
}

func TestGetLanguage(t *testing.T) {
	origAppConfig := GlobalAppConfig
	origTaskConfig := GlobalTaskConfig

	defer func() {
		GlobalAppConfig = origAppConfig
		GlobalTaskConfig = origTaskConfig
	}()

	config := newSampleAppConfig()

	testcases := []struct {
		name       string
		appConfig  AppConfig
		taskConfig TaskConfig
		expected   Language
		err        error
	}{
		{
			name:       "OK",
			appConfig:  newSampleAppConfig(),
			taskConfig: newSampleTaskConfig(),
			expected:   config.Languages[0],
		},
		{
			name:       "returns error when language not found",
			appConfig:  newAppConfigWithInvalidLanguageName(),
			taskConfig: newSampleTaskConfig(),
			err:        ErrLanguageNotFound,
		},
		{
			name:       "returns error when template not found",
			appConfig:  newSampleAppConfig(),
			taskConfig: newTaskConfigWithInvalidTemplateName(),
			err:        ErrTemplateNotFound,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			GlobalAppConfig = testcase.appConfig
			GlobalTaskConfig = testcase.taskConfig

			actual, err := GetLanguage()

			if diff := cmp.Diff(actual, testcase.expected); diff != "" {
				t.Errorf("expected %v, but actual %v", testcase.expected, actual)
			}

			if !errors.Is(err, testcase.err) {
				t.Errorf("expected %v, but actual %v", testcase.err, err)
			}
		})
	}
}

func TestGetTemplate(t *testing.T) {
	origAppConfig := GlobalAppConfig
	origTaskConfig := GlobalTaskConfig

	defer func() {
		GlobalAppConfig = origAppConfig
		GlobalTaskConfig = origTaskConfig
	}()

	config := newSampleAppConfig()

	testcases := []struct {
		name       string
		appConfig  AppConfig
		taskConfig TaskConfig
		expected   Template
		err        error
	}{
		{
			name:       "OK",
			appConfig:  newSampleAppConfig(),
			taskConfig: newSampleTaskConfig(),
			expected:   config.Templates[0],
		},
		{
			name:       "returns error when template not found",
			appConfig:  newSampleAppConfig(),
			taskConfig: newTaskConfigWithInvalidTemplateName(),
			err:        ErrTemplateNotFound,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			GlobalAppConfig = testcase.appConfig
			GlobalTaskConfig = testcase.taskConfig

			actual, err := GetTemplate()

			if diff := cmp.Diff(actual, testcase.expected); diff != "" {
				t.Errorf("expected %v, but actual %v", testcase.expected, actual)
			}

			if !errors.Is(err, testcase.err) {
				t.Errorf("expected %v, but actual %v", testcase.err, err)
			}
		})
	}
}

func newSampleAppConfig() AppConfig {
	return AppConfig{
		Username: "foo",
		Languages: []Language{
			{
				Name:        "csharp",
				AtCoderName: "csharp(100.x)",
				Build:       "echo build command",
				Run:         "echo run command",
			},
		},
		Templates: []Template{
			{
				Name:              "csharp-template",
				Language:          "csharp",
				TemplateDirectory: "templates/csharp",
				SourceFile:        "Program.cs",
			},
		},
		configDir:       "sample",
		DefaultTemplate: "csharp",
	}
}

func newSampleTaskConfig() TaskConfig {
	return TaskConfig{
		ContestID: "fooContest",
		TaskID:    "barTask",
		Template:  "csharp-template",
	}
}

func newTaskConfigWithInvalidTemplateName() TaskConfig {
	config := newSampleTaskConfig()
	config.Template = "not-existing-template"

	return config
}

func newAppConfigWithInvalidLanguageName() AppConfig {
	config := newSampleAppConfig()
	config.Templates[0] = Template{
		Name:     "csharp-template",
		Language: "invalid-language",
	}

	return config
}
