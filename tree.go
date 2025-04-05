package ghost

import (
	"fmt"
	"strings"
)

func (c *Context) DumpTree() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	var sb strings.Builder
	for _, s := range c.rootSpans {
		dumpSpan(&sb, s, 0)
	}
	return sb.String()
}

func dumpSpan(sb *strings.Builder, s *Span, level int) {
	indent := strings.Repeat("  ", level)
	duration := s.End.Sub(s.Start)
	sb.WriteString(fmt.Sprintf("%s↳ %s (%s)\n", indent, s.Name, duration))
	if s.Panick != nil {
		sb.WriteString(fmt.Sprintf("%s  ⚠️ PANIC:\n%s\n", indent, s.Panick.Stack))
	}
	for _, child := range s.Children {
		dumpSpan(sb, child, level+1)
	}
}
