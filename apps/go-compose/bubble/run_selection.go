package bubble

import "fmt"

type runSelectionView struct {
	choices []selection

	cursor int

	selected selection
}

func NewRunSelectionView() runSelectionView {
	return runSelectionView{
		choices: []selection{
			{label: "Hello World", value: "hello-world"},
			{label: "Factorial", value: "factorial"},
			{label: "Make Garbage", value: "make-garbage"},
		},
		cursor: 0,
	}
}

func (rsv *runSelectionView) Update(msg string) {
	switch msg {
	case "up", "k":
		if rsv.cursor > 0 {
			rsv.cursor--
		}
	case "down", "j":
		if rsv.cursor < len(rsv.choices)-1 {
			rsv.cursor++
		}
	}
}

func (srv runSelectionView) View() string {
	s := "Hello, choose the which of the following setups to run: \n\n"

	for i, choice := range srv.choices {
		cursor := "  "

		if srv.cursor == i {
			cursor = "->"
		}

		s += fmt.Sprintf("%s Run %s\n", cursor, choice.label)
	}

	return s
}

func (rsv *runSelectionView) Select() {
	rsv.selected = rsv.choices[rsv.cursor]
}
