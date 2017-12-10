//go:generate stringer -type=Pill

package pill

// Pill ...
type Pill int

// ...
const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
