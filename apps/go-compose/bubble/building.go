package bubble

import (
	"github.com/razmat145/learning/go-compose/builder"
)

type buildingView struct {
	state string
}

func NewBuildingView() buildingView {
	return buildingView{
		state: "starting",
	}
}

func (bv buildingView) View() string {
	s := "Building the previous choices... \n\n"

	switch bv.state {
	case "starting":
		s += "Starting to build compose... \n\n"

	case "building":
		s = "Building compose... \n\n"

	case "done":
		s = "Compose build complete! \n\n"

		s += "\nPress any key to exit.\n"
	}

	return s
}

func (bv *buildingView) Progress(msg builder.BuildProgress) {
	bv.state = msg.Progress
}
