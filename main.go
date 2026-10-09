// Command gh-glamour-v2-break is a minimal gh extension demonstrating that
// upgrading go-gh to glamour v2 (cli/go-gh#281) breaks extensions that never
// reference glamour types directly, purely through shared transitive deps.
package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh/v2/pkg/markdown"
)

const doc = `# Hello from gh-glamour-v2-break

This extension renders markdown with **go-gh** and styles a banner with
**lipgloss v1**. Neither call references a glamour type.

See https://github.com/cli/go-gh/pull/281
`

func main() {
	banner := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Render("gh glamour-v2-break")
	fmt.Println(banner)

	out, err := markdown.Render(doc, markdown.WithTheme("dark"), markdown.WithWrap(80))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(out)
}
