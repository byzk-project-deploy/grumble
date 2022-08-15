package grumble

import (
	"atomicgo.dev/keyboard"
	"fmt"
	"github.com/byzk-project-deploy/promptui"
	"github.com/pterm/pterm"
)

type ShellTools struct {
	app *App
}

func (s *ShellTools) keyboardFuncFilterInputRune(r rune) (rune, bool) {
	fmt.Println("处罚Key...")
	_ = keyboard.SimulateKeyPress(r)
	return 0, false
}

func (s *ShellTools) keyboardHandle(exit <-chan struct{}) {
	rawInputRune := s.app.rl.Config.FuncFilterInputRune
	defer func() { s.app.rl.Config.FuncFilterInputRune = rawInputRune }()

	s.app.rl.Config.FuncFilterInputRune = s.keyboardFuncFilterInputRune

	<-exit
}

func (s *ShellTools) Prompt(prompt *promptui.Prompt) (string, error) {
	prompt.Readline = s.app.rl
	return prompt.Run()
}

func (s *ShellTools) Confirm(label string) (bool, error) {
	s.app.rl.Terminal.EnterRawMode()
	defer s.app.rl.Terminal.ExitRawMode()

	exitChain := make(chan struct{}, 1)
	defer close(exitChain)

	go s.keyboardHandle(exitChain)

	return pterm.DefaultInteractiveConfirm.Show(label)
}
