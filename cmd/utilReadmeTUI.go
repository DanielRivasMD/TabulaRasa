////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// langItem wraps a string so it satisfies list.Item (FilterValue)
type langItem string

func (i langItem) FilterValue() string { return string(i) }

type step int

const (
	chooseLang step = iota
	enterDesc
	enterLicense
	confirm
	done
)

////////////////////////////////////////////////////////////////////////////////////////////////////

type model struct {
	step         step
	langList     list.Model
	descInput    textinput.Model
	licInput     textinput.Model
	selectedLang string
	description  string
	license      string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initialModel() model {
	// build list.Items from validLangs
	items := make([]list.Item, len(validLangs))
	for i, l := range validLangs {
		items[i] = langItem(l)
	}
	l := list.New(items, itemDelegate{}, 20, 5)
	l.Title = "Select README Language"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	// description input
	d := textinput.New()
	d.Placeholder = "Project description"
	d.Focus()

	// license input
	li := textinput.New()
	li.Placeholder = "License (e.g. MIT)"

	return model{
		step:      chooseLang,
		langList:  l,
		descInput: d,
		licInput:  li,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// itemDelegate renders our langItem
type itemDelegate struct{}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (itemDelegate) Height() int                               { return 1 }
func (itemDelegate) Spacing() int                              { return 0 }
func (itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	str := listItem.(langItem)
	cursor := " "
	if index == m.Index() {
		cursor = ">"
	}
	fmt.Fprintf(w, "%s %s\n", cursor, str)
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.step {
	case chooseLang:
		// allow arrow/enter handling in the list
		var cmd tea.Cmd
		m.langList, cmd = m.langList.Update(msg)
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			// capture selection
			sel := m.langList.SelectedItem().(langItem)
			m.selectedLang = string(sel)
			m.step = enterDesc
			m.descInput.Focus()
		}
		return m, cmd

	case enterDesc:
		// textinput handles typing
		var cmd tea.Cmd
		m.descInput, cmd = m.descInput.Update(msg)
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			m.description = m.descInput.Value()
			m.step = enterLicense
			m.licInput.Focus()
		}
		return m, cmd

	case enterLicense:
		var cmd tea.Cmd
		m.licInput, cmd = m.licInput.Update(msg)
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			m.license = m.licInput.Value()
			m.step = confirm
		}
		return m, cmd

	case confirm:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			m.step = done
			// fire off deploy in background
			go m.deploy()
		}
		return m, nil

	case done:
		// exit the TUI
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	switch m.step {
	case chooseLang:
		return "\n" + m.langList.View() + "\n(↑/↓: navigate • enter: select)\n"

	case enterDesc:
		return fmt.Sprintf(
			"\nEnter a short project description:\n\n  %s\n",
			m.descInput.View(),
		)

	case enterLicense:
		return fmt.Sprintf(
			"\nEnter license identifier:\n\n  %s\n",
			m.licInput.View(),
		)

	case confirm:
		return fmt.Sprintf(
			"\nConfirm settings:\n\n  language: %s\n  description: %s\n  license: %s\n\nPress ENTER to deploy\n",
			m.selectedLang, m.description, m.license,
		)

	case done:
		return "\nDeployment triggered! Exiting…\n"
	}

	return ""
}

// deploy() calls your existing Cobra logic using the collected values.
func (m model) deploy() {
	// 1) fallback to current dir
	repo, err := domovoi.CurrentDir()
	horus.CheckErr(err)

	// 2) find TabulaRasa home
	home, err := domovoi.FindHome(false)
	horus.CheckErr(err)

	// 3) build copy params
	srcDir := filepath.Join(home, readmeDir)
	destFile := filepath.Join(path, readme)

	params := newCopyParams(srcDir, destFile)
	params.Files = []string{
		overview,
		filepath.Join("02" + m.selectedLang + "_install.md"),
		usage,
		filepath.Join("04" + m.selectedLang + "_dev.md"),
		faq,
	}

	// 4) replacements + concat
	reps, repErr := buildReadmeReplacements(
		m.selectedLang,
		m.description,
		repo,
		user,
		author,
		m.license,
		path,
	)
	horus.CheckErr(repErr)

	params.Reps = reps
	horus.CheckErr(concatenateFiles(params, ""))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
