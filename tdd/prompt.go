package tdd

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed prompts/*.tmpl prompts/*.txt
var promptFiles embed.FS

type PromptData struct {
	Instructions string
	Intro        string
	CommonRules  string
}

func PromptFrom(templatePath string, instructions string) string {
	intro, _ := promptFiles.ReadFile("prompts/intro.txt")
	commonRules, _ := promptFiles.ReadFile("prompts/common_rules.txt")
	tmpl, _ := template.ParseFS(promptFiles, "prompts/"+templatePath)
	var prompt bytes.Buffer
	tmpl.Execute(&prompt, PromptData{
		Instructions: instructions,
		Intro:        string(intro),
		CommonRules:  string(commonRules)})
	return prompt.String()
}
