package capp

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// defaultNameWidth default max width for command name, on render help.
const defaultNameWidth = 12

// cmdSet manages a group of commands: names, commands map, aliases and name width.
//
// It is used by App for the top commands, and by Cmd for the sub commands.
type cmdSet struct {
	names []string
	cmds  map[string]*Cmd
	cmdAs maputil.Aliases
	// nameWidth max width for command name
	nameWidth int
}

// newCmdSet create a new cmdSet instance
func newCmdSet() *cmdSet {
	return &cmdSet{
		cmds:      make(map[string]*Cmd),
		cmdAs:     make(maputil.Aliases),
		nameWidth: defaultNameWidth,
	}
}

// add a command to the set.
//
// NOTE:
//   - will call c.initCmd() and fire the c.OnAdd hook.
//   - c.Name should not be empty, and should not be repeated.
func (s *cmdSet) add(c *Cmd) error {
	if len(c.Name) == 0 {
		return errors.New("command name cannot be empty")
	}

	if _, ok := s.cmds[c.Name]; ok {
		return fmt.Errorf("command name %s has been exists", c.Name)
	}

	s.names = append(s.names, c.Name)
	s.cmds[c.Name] = c
	s.nameWidth = mathutil.MaxInt(s.nameWidth, s.nameLen(c))

	// add aliases
	if len(c.Aliases) > 0 {
		s.cmdAs.AddAliases(c.Name, c.Aliases)
	}

	// attach handle func
	c.initCmd()

	if c.OnAdd != nil {
		c.OnAdd(c)
	}
	return nil
}

// nameLen the render width for a command: command name and aliases.
func (s *cmdSet) nameLen(c *Cmd) int {
	ln := len(c.Name)
	if len(c.Aliases) > 0 {
		ln += len(strings.Join(c.Aliases, ", ")) + 4
	}
	return ln
}

// find a command by name. will resolve the alias name.
func (s *cmdSet) find(name string) (*Cmd, bool) {
	if name == "" {
		return nil, false
	}

	// resolve alias
	name = s.cmdAs.ResolveAlias(name)
	c, ok := s.cmds[name]
	return c, ok
}

// hasCommands check the set has commands.
func (s *cmdSet) hasCommands() bool {
	return s != nil && len(s.names) > 0
}

// render the "Commands:" help list to buffer.
//
// - width: max width for render the command name.
func (s *cmdSet) render(buf *strutil.Buffer, width int) {
	sort.Strings(s.names)

	gaMap := s.cmdAs.GroupAliases()
	for _, name := range s.names {
		c := s.cmds[name]
		if len(gaMap[name]) > 0 {
			name = name + ", " + strings.Join(gaMap[name], ", ")
		}

		name = strutil.PadRight(name, " ", width)
		buf.Printf("  <green>%s</>    %s\n", name, strutil.UpperFirst(c.getDesc()))
	}
}

// propagateApp set the app to all commands in the set.
func (s *cmdSet) propagateApp(a *App) {
	for _, name := range s.names {
		s.cmds[name].setApp(a)
	}
}
