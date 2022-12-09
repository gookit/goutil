package termctrl

// Giant list of key constants.
//
// Everything above KeyUnknown matches an actual ASCII key value.
// After that, we have various pseudo-keys in order to
// represent complex byte sequences that correspond to keys like Page up, Right
// arrow, etc.
const (
	KeyCharA = 1 + iota
	KeyCharB
	KeyCharC
	KeyCharD
	KeyCharE
	KeyCharF
	KeyCharG
	KeyCharH
	KeyCharI
	KeyCharJ
	KeyCharK
	KeyCharL
	KeyCharM
	KeyCharN
	KeyCharO
	KeyCharP
	KeyCharQ
	KeyCharR
	KeyCharS
	KeyCharT
	KeyCharU
	KeyCharV
	KeyCharW
	KeyCharX
	KeyCharY
	KeyCharZ
	KeyEscape

	KeyLeftBracket  = '['
	KeyRightBracket = ']'
	KeyEnter        = '\r'
	KeyBackspace    = 127

	KeyUnknown = 0xd800 /* UTF-16 surrogate area */ + iota
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyHome
	KeyEnd
	KeyPasteStart
	KeyPasteEnd
	KeyInsert
	KeyDelete
	KeyPgUp
	KeyPgDn
	KeyPause
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
)

// TODO ...
const (
	KeyCtrl = ' '
	KeyAlt  = ' '
)

// special chars on paste start or end
var (
	PasteStart = []byte{KeyEscape, '[', '2', '0', '0', '~'}
	PasteEnd   = []byte{KeyEscape, '[', '2', '0', '1', '~'}
)
