package ptr

// Of returns a pointer to the provided value.
// It is a generic replacement for previous helper functions.
func Of[T any](v T) *T { return &v }

// Int returns a pointer to the provided int value.
// Deprecated: use Of instead, e.g. Of(10).
func Int(v int) *int { return &v }

// String returns a pointer to the provided string value.
// Deprecated: use Of instead, e.g. Of("hello").
func String(v string) *string { return &v }

// Bool returns a pointer to the provided bool value.
// Deprecated: use Of instead, e.g. Of(true).
func Bool(v bool) *bool { return &v }


