package vocab

type Object interface {
	GetNote() ObjectNote
	SetNote(note ObjectNote)
}
