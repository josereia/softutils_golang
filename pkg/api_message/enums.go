package messages

type Language int

const (
	EN Language = iota
	PT
)

func (lang Language) String() string {
	return []string{"EN", "PT"}[lang]
}

type Level int

const (
	PANIC Level = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

func (lvl Level) String() string {
	return []string{"PANIC", "ERROR", "INFO", "DEBUG"}[lvl]
}
