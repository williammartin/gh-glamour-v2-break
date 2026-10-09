# gh-glamour-v2-break

Minimal `gh` extension reproducing a downstream breakage from
[cli/go-gh#281](https://github.com/cli/go-gh/pull/281) (glamour v1 → `charm.land/glamour/v2`).

The extension only calls `markdown.Render` / `markdown.WithTheme` / `markdown.WithWrap`
from go-gh. It never references a glamour type. It also uses `github.com/charmbracelet/lipgloss`
v1, which is common in extensions (directly, or via bubbletea / huh).

## The break

Upgrading go-gh to the #281 commit pulls in `github.com/charmbracelet/x/ansi` v0.11.x.
Minimal version selection then applies that version to lipgloss v1's dependency
`github.com/charmbracelet/x/cellbuf` v0.0.13, which no longer compiles:

```
github.com/charmbracelet/x/cellbuf/cell.go: not enough arguments in call to b.Italic
	have ()
	want (bool)
```

It also raises this module's `go` directive to `1.25.8`.

## Workaround

```sh
go get github.com/charmbracelet/x/cellbuf@v0.0.15
```
