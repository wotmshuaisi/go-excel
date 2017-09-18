package internal

type Config struct {
	// Sheet name as string or sheet model as object or a slice of object.
	Sheet interface{}
	// Use the index row as title, every row before title-row will be ignore, default is 0.
	TitleRowIndex int
	// Skip n row after title, default is 0 (not skip).
	Skip int
	// Auto prefix to sheet name.
	Prefix string
	// Auto suffix to sheet name.
	Suffix string
}
