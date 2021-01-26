package cliutil

import "strings"

// LineParser struct
// parse input command line to []string, such as cli os.Args
type LineParser struct {
	// Line the full input command line text
	// eg `kite top sub -a "the a message" --foo val1 --bar "val 2"`
	Line string
	// the exploded nodes by space.
	nodes []string
	// the parsed args
	args []string
}

// StringToOSArgs parse input command line text to os.Args
func StringToOSArgs(line string) []string {
	p := &LineParser{
		Line: line,
	}

	return p.Parse()
}

// ParseLine input command line text. alias of the StringToOSArgs()
func ParseLine(line string) []string {
	p := &LineParser{
		Line: line,
	}

	return p.Parse()
}

// Parse input command line text to os.Args
func (p *LineParser) Parse() []string {
	p.Line = strings.TrimSpace(p.Line)
	if p.Line == "" {
		return p.args
	}

	p.nodes = strings.Split(p.Line, " ")
	if len(p.nodes) == 1 {
		p.args = p.nodes
		return p.args
	}

	// temp value
	var quoteChar, fullNode string
	for _, node := range p.nodes {
		if node == "" {
			continue
		}

		nodeLen := len(node)
		start, end := node[:1], node[nodeLen-1:]

		var clearTemp bool
		if start == "'" || start == `"` {
			noStart := node[1:]
			if quoteChar == "" { // start
				// only one words. eg: `-m "msg"`
				if end == start {
					p.args = append(p.args, node[1:nodeLen-1])
					continue
				}

				fullNode += noStart
				quoteChar = start
			} else if quoteChar == start { // invalid. eg: `-m "this is "message` `-m "this is "message"`
				p.appendWithPrefix(strings.Trim(node, quoteChar), fullNode)
				clearTemp = true // clear temp value
			} else if quoteChar == end { // eg: `"has inner 'quote'"`
				p.appendWithPrefix(node[:nodeLen-1], fullNode)
				clearTemp = true // clear temp value
			} else { // goon. eg: `-m "the 'some' message"`
				fullNode += " " + node
			}
		} else if end == "'" || end == `"` {
			noEnd := node[:nodeLen-1]
			if quoteChar == "" { // end
				p.appendWithPrefix(noEnd, fullNode)
				clearTemp = true // clear temp value
			} else if quoteChar == end { // end
				p.appendWithPrefix(noEnd, fullNode)
				clearTemp = true // clear temp value
			} else { // goon. eg: `-m "the 'some' message"`
				fullNode += " " + node
			}
		} else {
			if quoteChar != "" {
				fullNode += " " + node
			} else {
				p.args = append(p.args, node)
			}
		}

		// clear temp value
		if clearTemp {
			quoteChar, fullNode = "", ""
		}
	}

	if fullNode != "" {
		p.args = append(p.args, fullNode)
	}

	return p.args
}

func (p *LineParser) appendWithPrefix(node, prefix string) {
	if prefix != "" {
		p.args = append(p.args, prefix+" "+node)
	} else {
		p.args = append(p.args, node)
	}
}
