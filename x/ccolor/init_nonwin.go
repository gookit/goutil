//go:build !windows

package ccolor

func init() {
	CheckColorSupport()
}
