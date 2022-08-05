package grumble

import "github.com/byzk-project-deploy/promptui"

type ShellTools struct {
	app *App
}

func (s *ShellTools) Prompt(prompt promptui.Prompt) (string, error) {
	prompt.Readline = s.app.rl
	return prompt.Run()
}
