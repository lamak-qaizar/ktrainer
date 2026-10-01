package ui

import "github.com/pterm/pterm"

type Spinner struct {
	spinner *pterm.SpinnerPrinter
}

func NewSpinner(str string) *Spinner {
	spinner, _ := pterm.DefaultSpinner.Start(str)
	return &Spinner{spinner: spinner}
}

func (s *Spinner) Success() {
	s.spinner.Success()
}

func (s *Spinner) Fail() {
	s.spinner.Fail()
}
