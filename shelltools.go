package grumble

import (
	"atomicgo.dev/keyboard"
	"github.com/byzk-project-deploy/promptui"
	"github.com/pterm/pterm"
)

type ShellTools struct {
	app              *App
	isIgnoreKeyboard bool
	exitChan         chan struct{}
}

func (s *ShellTools) emptyKeyboardFuncFilterInputRune(r rune) (rune, bool) {
	return 0, false
}

func (s *ShellTools) keyboardFuncFilterInputRune(r rune) (rune, bool) {
	_ = keyboard.SimulateKeyPress(r)
	return 0, false
}

func (s *ShellTools) beginKeyboardHandle(keyboardHandleFn func(rune) (rune, bool)) {
	s.app.rl.Terminal.EnterRawMode()
	go func() {
		rawInputRune := s.app.rl.Config.FuncFilterInputRune
		defer func() { s.app.rl.Config.FuncFilterInputRune = rawInputRune }()

		s.app.rl.Config.FuncFilterInputRune = keyboardHandleFn

		<-s.exitChan
		s.exitChan <- struct{}{}
	}()
}

func (s *ShellTools) exitKeyboardHandle() {
	s.exitChan <- struct{}{}
	<-s.exitChan
	s.app.rl.Terminal.ExitRawMode()
}

func (s *ShellTools) EnterIgnoreKeyboard() {
	if s.isIgnoreKeyboard {
		return
	}

	s.beginKeyboardHandle(s.emptyKeyboardFuncFilterInputRune)
}

func (s *ShellTools) ExitIgnoreKeyboard() {
	if !s.isIgnoreKeyboard {
		return
	}

	s.exitKeyboardHandle()
}

func (s *ShellTools) Prompt(prompt *promptui.Prompt) (string, error) {
	prompt.Readline = s.app.rl
	return prompt.Run()
}

func (s *ShellTools) Confirm(label string) (bool, error) {
	s.beginKeyboardHandle(s.keyboardFuncFilterInputRune)
	defer s.exitKeyboardHandle()

	return pterm.DefaultInteractiveConfirm.Show(label)
}
