package numutil

import "github.com/gookit/goutil/mathutil"

//goland:noinspection GoUnusedGlobalVariable
var (
	// TryToString try convert intX/floatX value to string
	TryToString   = mathutil.TryToString
	StringOrPanic = mathutil.StringOrPanic
	MustString    = mathutil.MustString
	ToString      = mathutil.ToString
	// ToFloat convert value to float64
	ToFloat = mathutil.ToFloat
	RandInt = mathutil.RandInt

	RandIntWithSeed = mathutil.RandIntWithSeed
)
