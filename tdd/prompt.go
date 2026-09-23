package tdd

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed prompts/*.tmpl
var promptFiles embed.FS

func PromptFrom(templatePath string, instructions string) string {
	tmpl, _ := template.ParseFS(promptFiles, "prompts/"+templatePath)
	var prompt bytes.Buffer
	tmpl.Execute(&prompt, struct{ Instructions string }{Instructions: instructions})
	return prompt.String()
}
