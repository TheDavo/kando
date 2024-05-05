package ui

import (
	"fmt"
	"kando/kando"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type addTaskModel struct {
	focusIndex        int
	focusStatus       int
	inputs            []textinput.Model
	totalInputs       int
	cursorMode        cursor.Mode
	project           string
	statuses          kando.Statuses
	selStatus         kando.Status
	selStatusDetailed kando.StatusDetailed
	k                 *kando.Kando
	dbg               string
}

func AddTaskInitialModel(proj string) addTaskModel {
	m := addTaskModel{
		inputs: make([]textinput.Model, 1),
		k:      kando.Open(),
	}

	m.statuses = m.k.Projects[proj].Statuses

	m.totalInputs += len(m.inputs)

	for _, v := range m.statuses {
		m.totalInputs += len(v)
	}

	var t textinput.Model

	t = textinput.New()
	t.Cursor.Style = cursorStyle
	t.CharLimit = 50

	t.Placeholder = "description"
	t.Focus()
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle

	m.inputs[0] = t
	m.project = proj

	return m
}

func (m addTaskModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m addTaskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":

			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}

			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == m.totalInputs {

				if m.inputs[0].Value() == "" {
					m.dbg = "Please enter a description, task entry canceled!"
					return m, tea.Quit
				}
				newTask := kando.Task{
					Description:    m.inputs[0].Value(),
					Status:         m.selStatus,
					StatusDetailed: m.selStatusDetailed,
				}

				m.k.Projects[m.project].AddTask(newTask)

				err := m.k.Save()

				if err != nil {
					panic(err)
				}
				return m, tea.Quit
			}

			inStatusRange := m.focusIndex >= len(m.inputs) &&
				m.focusIndex < m.totalInputs
			statusOffset := len(m.inputs)
			newIdx := m.focusIndex - statusOffset

			lenTodo := len(m.statuses["todo"])
			lenInProg := len(m.statuses["in-progress"])

			inTodoRange := newIdx < lenTodo
			inInProgRange := lenTodo <= newIdx &&
				newIdx < (lenInProg+lenTodo)

			if s == "enter" && inStatusRange {
				if inTodoRange {
					m.selStatus = kando.Todo
					m.selStatusDetailed = m.statuses["todo"][newIdx]
				} else if inInProgRange {
					m.selStatus = kando.InProgress
					m.selStatusDetailed =
						m.statuses["in-progress"][newIdx-lenTodo]
				} else {
					m.selStatus = kando.Done
					m.selStatusDetailed =
						m.statuses["done"][newIdx-lenTodo-lenInProg]
				}
				m.focusStatus = m.focusIndex
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > m.totalInputs {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = m.totalInputs
			}

			cmds := make([]tea.Cmd, m.totalInputs)
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}

				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle

			}

			return m, tea.Batch(cmds...)
		}

	}
	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *addTaskModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m addTaskModel) View() string {
	var b strings.Builder
	checked := " "

	for i := range m.inputs {
		// Update the textinput Views
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	b.WriteRune('\n')
	b.WriteRune('\n')

	// Calc breakpoints for focused Style rendering
	todoStart := len(m.inputs)
	inProgStart := todoStart + len(m.statuses["todo"])
	doneStart := inProgStart + len(m.statuses["in-progress"])

	b.WriteString(todoStatusStyle.Render("Todo"))
	b.WriteRune('\n')
	for i, v := range m.statuses["todo"] {
		if v == m.selStatusDetailed {
			checked = "x"
		} else {
			checked = " "
		}
		renderStr := fmt.Sprintf("[%s] %s", checked, v)
		if m.focusIndex == todoStart+i {
			b.WriteString(focusedStyle.Render(renderStr))
		} else {
			b.WriteString(renderStr)
		}
	}
	b.WriteRune('\n')
	b.WriteRune('\n')
	b.WriteString(inProgStatuStyle.Render("In Progress"))
	b.WriteRune('\n')
	for i, v := range m.statuses["in-progress"] {
		if v == m.selStatusDetailed {
			checked = "x"
		} else {
			checked = " "
		}
		renderStr := fmt.Sprintf("[%s] %s", checked, v)
		if m.focusIndex == inProgStart+i {
			b.WriteString(focusedStyle.Render(renderStr))
		} else {
			b.WriteString(renderStr)
		}
	}
	b.WriteRune('\n')
	b.WriteRune('\n')
	b.WriteString(doneStyle.Render("Done"))
	b.WriteRune('\n')
	for i, v := range m.statuses["done"] {
		if v == m.selStatusDetailed {
			checked = "x"
		} else {
			checked = " "
		}
		renderStr := fmt.Sprintf("[%s] %s", checked, v)
		if m.focusIndex == doneStart+i {
			b.WriteString(focusedStyle.Render(renderStr))
		} else {
			b.WriteString(renderStr)
		}
	}
	b.WriteRune('\n')

	button := &blurredButton
	if m.focusIndex == m.totalInputs {
		button = &focusedButton
	}

	b.WriteString(m.dbg)
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}
