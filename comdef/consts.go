package comdef

// constants for compare operation
const (
	OpEq  = "="
	OpNeq = "!="
	OpLt  = "<"
	OpLte = "<="
	OpGt  = ">"
	OpGte = ">="
)

// constants quote chars
const (
	SingleQuote = '\''
	DoubleQuote = '"'
	SlashQuote  = '\\'

	SingleQuoteStr = "'"
	DoubleQuoteStr = `"`
	SlashQuoteStr  = "\\"
)

// NoIdx invalid index or length
const NoIdx = -1

// const VarPathReg = `(\w[\w-]*(?:\.[\w-]+)*)`

// Align define align, position: L, C, R, Auto
type Align uint8
type Position = Align // Position alias of Align

// constants for align, position: L, C, R, Auto
const (
	Left Align = iota
	Center
	Right
	PosAuto
)
