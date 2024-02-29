package goutil

type RuneWriter interface {
	WriteRune(rune) (int, error)
}

type PrintLike interface {
	Print(...any) error
	Printf(string, ...any) error
	Println(...any) error
}

type Discard struct {}

func(Discard) Write(p []byte) (int, error) {
	return len(p), nil
}
