package prompts

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed *.tmpl
var promptFiles embed.FS

func PromptFrom(templatePath string, instructions string) string {
	tmpl, _ := template.ParseFS(promptFiles, templatePath)
	var prompt bytes.Buffer
	tmpl.Execute(&prompt, struct{ Instructions string }{Instructions: instructions})
	return prompt.String()
}
