package bubble

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/razmat145/learning/go-compose/builder"
)

type model struct {
	inChan   chan builder.BuildProgress
	outChan  chan builder.RunEnvironment
	location string

	runSelectionView runSelectionView
	appSelectionView appSelectionView
	buildingView     buildingView
}

type selection struct {
	label string
	value string
}

func InitialiseModel(in chan builder.BuildProgress, out chan builder.RunEnvironment) model {
	return model{
		inChan:           in,
		outChan:          out,
		location:         "run-selection",
		runSelectionView: NewRunSelectionView(),
		appSelectionView: NewAppSelectionView(),
		buildingView:     NewBuildingView(),
	}
}

func waitForActivity(inChan chan builder.BuildProgress) tea.Cmd {
	return func() tea.Msg {
		return builder.BuildProgress(<-inChan)
	}
}

func (m model) Init() tea.Cmd {
	return waitForActivity(m.inChan)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case builder.BuildProgress:
		m.buildingView.Progress(msg)

		return m, waitForActivity(m.inChan)

	case tea.KeyMsg:
		if (m.location == "building") && (m.buildingView.state == "done") {
			return m, tea.Quit
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "backspace", "p":
			switch m.location {
			case "app-selection":
				m.location = "run-selection"
			}

		case "b":
			switch m.location {
			case "app-selection":
				m.location = "building"

				m.outChan <- builder.RunEnvironment{
					BenchmarkTarget: m.runSelectionView.selected.value,
					AppsToRun:       extractFromMap(m.appSelectionView.selected),
				}
			}

		case "enter", " ":
			switch m.location {
			case "run-selection":
				m.runSelectionView.Select()
				m.location = "app-selection"

			case "app-selection":
				m.appSelectionView.Select()
			}

		default:
			switch m.location {
			case "run-selection":
				m.runSelectionView.Update(msg.String())

			case "app-selection":
				m.appSelectionView.Update(msg.String())
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string

	switch m.location {
	case "run-selection":
		s = m.runSelectionView.View()

	case "app-selection":
		s = m.appSelectionView.View(m)

	case "building":
		s = m.buildingView.View()
	}

	if m.location != "building" {
		s += "\nPress q or ctrl+c to quit.\n"
	}
	return s
}

func extractFromMap(m map[int]string) []string {
	var res []string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
