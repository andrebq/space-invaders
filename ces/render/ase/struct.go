package ase

type (
	// Animation handles the ASE animation
	Animation struct {
		Frames []Frame `json:"frames"`
	}

	// Frame controls a single animation frame
	Frame struct {
		Frame    Rect `json:"frame"`
		Duration float64
	}

	// Rect holds a rectangle
	Rect struct {
		X int32 `json:"x"`
		Y int32 `json:"y"`
		W int32 `json:"w"`
		H int32 `json:"h"`
	}
)
