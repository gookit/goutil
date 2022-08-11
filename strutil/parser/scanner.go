package parser

const (
	TokInvalid = iota
	TokComments
	TokILComments
	TokMLComments
	TokValue
	TokMLValue
)

// TextScanner struct
type TextScanner struct {
}
