package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ev-the-dev/bootmaker/models"
)

var version string = "123"

/*
*
* TODO:
  - Auto select all options on the WizardAnswers
  - For each selected, create DTO and Adapter
*
*/

type WizardState struct {
	cursor int

	moduleName                  string
	moduleNameTextInput         textinput.Model
	moduleNameTextInputRendered bool

	selected map[int]bool
	styles   *Styles

	height int
	width  int
}

func Execute() {
	wS := &WizardState{
		moduleName: "",
		selected: map[int]bool{
			models.CONTROLLER:     true,
			models.QUEUE_CONSUMER: true,
			models.QUEUE_PRODUCER: true,
			models.REPOSITORY:     true,
			models.SERVICE:        true,
		},
		styles: DefaultStyles(),
	}

	p := tea.NewProgram(wS)
	m, err := p.Run()
	if err != nil {
		log.Fatalf("Oopsie doopsie! Can't start TUI: %v", err)
	}

	updatedWS := m.(WizardState)
	wA := &models.WizardAnswers{
		Controller:    updatedWS.selected[models.CONTROLLER],
		ModuleName:    updatedWS.moduleName,
		QueueConsumer: updatedWS.selected[models.QUEUE_CONSUMER],
		QueueProducer: updatedWS.selected[models.QUEUE_PRODUCER],
		Repository:    updatedWS.selected[models.REPOSITORY],
		Service:       updatedWS.selected[models.SERVICE],
	}

	// pass wA to write stream
	fmt.Printf("wizard answers::: %+v: ", wA)
	generateFiles(wA)
}

func (w WizardState) Init() tea.Cmd {
	return nil
}

func (w WizardState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.height = msg.Height
		w.width = msg.Width
	}

	if w.moduleName == "" {
		if w.moduleNameTextInputRendered == false {
			w.moduleNameTextInput = textinput.New()
			w.moduleNameTextInput.Placeholder = "Type your Module Name"
			w.moduleNameTextInput.Focus()
			w.moduleNameTextInputRendered = true
		}
		return w.updateTextInput(msg)
	} else {
		return w.updateSelector(msg)
	}
}

func (w WizardState) View() string {
	if w.width == 0 {
		return "Loading..."
	}

	if w.moduleName == "" {
		return lipgloss.JoinVertical(
			lipgloss.Center,
			w.moduleNameTextInput.Value(),
			w.styles.InputField.Render(w.moduleNameTextInput.View()),
		)
	} else {
		s := "\nSelect each component you would like to include in your module:\n\n"
		for i := 0; i < len(w.selected); i++ {
			cursor := " "
			if w.cursor == i {
				cursor = ">"
			}

			selected := " "
			if v, ok := w.selected[i]; ok && v {
				selected = "X"
			}

			keyName := mapSelectedKeyToName(i)
			s += renderRow(cursor, selected, keyName)
		}

		return s
	}
}

func mapSelectedKeyToName(key int) string {
	switch key {
	case models.CONTROLLER:
		return "Controller"
	case models.QUEUE_CONSUMER:
		return "Queue Consumer"
	case models.QUEUE_PRODUCER:
		return "Queue Producer"
	case models.REPOSITORY:
		return "Repository"
	case models.SERVICE:
		return "Service"
	}
	return ""
}

func renderRow(cursor string, checked string, choice string) string {
	return fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
}

func (w WizardState) updateSelector(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return w, tea.Quit
		case "up", "k":
			if w.cursor > 0 {
				w.cursor--
			}
		case "down", "j":
			if int(w.cursor) < len(w.selected)-1 {
				w.cursor++
			}
		case "enter", " ":
			b, ok := w.selected[w.cursor]
			if ok {
				w.selected[w.cursor] = !b
			} else {
				log.Fatalf("Couldn't access selected map. Accessed with key: (%d)", w.cursor)
			}
		}
	}

	return w, nil
}

func (w WizardState) updateTextInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	w.moduleNameTextInput, cmd = w.moduleNameTextInput.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		case "enter":
			w.moduleName = w.moduleNameTextInput.Value()
			w.moduleNameTextInput.Blur()
		}
	}
	return w, cmd
}

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderTopForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)

	return s
}
