package bubble

import "fmt"

type appSelectionView struct {
	choices []string

	cursor int

	selected map[int]string
}

func NewAppSelectionView() appSelectionView {
	choices := []string{
		"node-express",
		"go-echo",
		"node-std",
		"go-std",
	}

	return appSelectionView{
		choices:  choices,
		cursor:   0,
		selected: make(map[int]string, len(choices)),
	}
}

func (asv *appSelectionView) Update(msg string) {
	switch msg {
	case "up", "k":
		if asv.cursor > 0 {
			asv.cursor--
		}
	case "down", "j":
		if asv.cursor < len(asv.choices)-1 {
			asv.cursor++
		}
	}
}

func (asv appSelectionView) View(m model) string {
	s := fmt.Sprintf("Choose the which of the following apps you want to include in the '%s' run: \n\n", m.runSelectionView.selected.label)

	for i, choice := range asv.choices {
		cursor := "  "

		if asv.cursor == i {
			cursor = "->"
		}

		checked := " "
		if _, ok := asv.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress b to start building the test to run."
	s += "\nPress p to re-select the test to run."
	return s
}

func (asv appSelectionView) Select() {
	_, ok := asv.selected[asv.cursor]
	if ok {
		delete(asv.selected, asv.cursor)
	} else {
		asv.selected[asv.cursor] = asv.choices[asv.cursor]
	}
}
