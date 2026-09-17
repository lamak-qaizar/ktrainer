package main

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed prompts/*.tmpl
var promptFiles embed.FS

func generatePromptFromTemplate(instructions string) string {
	tmpl, _ := template.ParseFS(promptFiles, "prompts/test.tmpl")
	var prompt bytes.Buffer
	tmpl.Execute(&prompt, struct{ Instructions string }{Instructions: instructions})
	return prompt.String()
}
