//go:build windows
// +build windows

package sysutil

// ChangeUserByName change work user by new username.
func ChangeUserByName(newUname string) (err error) {
	return ChangeUserUIDGid(0, 0)
}

// ChangeUserUidGid change work user by new username uid,gid
//
// Deprecated: use ChangeUserUIDGid instead
func ChangeUserUidGid(newUid int, newGid int) (err error) {
	return ChangeUserUIDGid(newUid, newGid)
}

// ChangeUserUIDGid change work user by new username uid,gid
func ChangeUserUIDGid(newUid int, newGid int) (err error) {
	return nil
}
