////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	enterDesc step = iota
	chooseLicense
	confirm
	done
)

var (
	licenses = []string{"MIT", "GPLv3"}
)

type step int
type licenseItem string

////////////////////////////////////////////////////////////////////////////////////////////////////

func (i licenseItem) FilterValue() string { return string(i) }

////////////////////////////////////////////////////////////////////////////////////////////////////

type model struct {
	step        step
	descInput   textinput.Model
	licList     list.Model
	description string
	license     string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initialModel() model {
	// description input
	d := textinput.New()
	d.Placeholder = "Project description"
	d.Focus()

	// license list
	licenses := licenses
	items := make([]list.Item, len(licenses))
	for i, l := range licenses {
		items[i] = licenseItem(l)
	}
	l := list.New(items, itemDelegate{}, 20, 5)
	l.Title = "Select License"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return model{
		step:      enterDesc,
		descInput: d,
		licList:   l,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// itemDelegate renders our list items
type itemDelegate struct{}

func (itemDelegate) Height() int                               { return 1 }
func (itemDelegate) Spacing() int                              { return 0 }
func (itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	str := listItem.(licenseItem)
	cursor := " "
	if index == m.Index() {
		cursor = ">"
	}
	fmt.Fprintf(w, "%s %s\n", cursor, str)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.step {
	case enterDesc:
		var cmd tea.Cmd
		m.descInput, cmd = m.descInput.Update(msg)
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			m.description = m.descInput.Value()
			m.step = chooseLicense
		}
		return m, cmd

	case chooseLicense:
		var cmd tea.Cmd
		m.licList, cmd = m.licList.Update(msg)
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			sel := m.licList.SelectedItem().(licenseItem)
			m.license = string(sel)
			m.step = confirm
		}
		return m, cmd

	case confirm:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "enter" {
			m.step = done
		}
		return m, nil

	case done:
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	switch m.step {
	case enterDesc:
		return fmt.Sprintf(
			"\nEnter a short project description:\n\n  %s\n",
			m.descInput.View(),
		)

	case chooseLicense:
		return "\n" + m.licList.View() + "\n(↑/↓: navigate • enter: select)\n"

	case confirm:
		return fmt.Sprintf(
			"\nConfirm settings:\n\n  description: %s\n  license: %s\n\nPress ENTER to finish\n",
			m.description, m.license,
		)

	case done:
		return "\nCaptured values! Exiting…\n"
	}

	return ""
}

////////////////////////////////////////////////////////////////////////////////////////////////////
