package docs

import (
	"html"
	"strconv"
	"strings"
	"unicode/utf8"
)

func markdownToHTML(source string) string {
	lines := strings.Split(strings.ReplaceAll(source, "\r\n", "\n"), "\n")

	var out strings.Builder
	var paragraph []string
	var quote []string
	var listItems []string
	var listTag string
	var code strings.Builder
	inCode := false

	flushParagraph := func() {
		if len(paragraph) == 0 {
			return
		}
		out.WriteString("<p>")
		out.WriteString(inlineMarkdown(strings.Join(paragraph, " ")))
		out.WriteString("</p>")
		paragraph = nil
	}

	flushQuote := func() {
		if len(quote) == 0 {
			return
		}
		out.WriteString("<blockquote><p>")
		out.WriteString(inlineMarkdown(strings.Join(quote, " ")))
		out.WriteString("</p></blockquote>")
		quote = nil
	}

	flushList := func() {
		if len(listItems) == 0 {
			return
		}
		out.WriteByte('<')
		out.WriteString(listTag)
		out.WriteByte('>')
		for _, item := range listItems {
			out.WriteString("<li>")
			out.WriteString(item)
			out.WriteString("</li>")
		}
		out.WriteString("</")
		out.WriteString(listTag)
		out.WriteByte('>')
		listItems = nil
		listTag = ""
	}

	flushCode := func() {
		out.WriteString("<pre><code>")
		out.WriteString(html.EscapeString(code.String()))
		out.WriteString("</code></pre>")
		code.Reset()
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		if inCode {
			if strings.HasPrefix(trimmed, "```") {
				flushCode()
				inCode = false
				continue
			}
			code.WriteString(line)
			code.WriteByte('\n')
			continue
		}

		if strings.HasPrefix(trimmed, "```") {
			flushParagraph()
			flushQuote()
			flushList()
			inCode = true
			continue
		}

		if trimmed == "" {
			flushParagraph()
			flushQuote()
			flushList()
			continue
		}

		if strings.Contains(trimmed, "|") && i+1 < len(lines) && tableSeparator(lines[i+1]) {
			flushParagraph()
			flushQuote()
			flushList()

			header := tableCells(trimmed)
			i++
			var rows [][]string
			for i+1 < len(lines) {
				candidate := strings.TrimSpace(lines[i+1])
				if candidate == "" || !strings.Contains(candidate, "|") {
					break
				}
				rows = append(rows, tableCells(candidate))
				i++
			}

			out.WriteString("<table><thead><tr>")
			for _, cell := range header {
				out.WriteString("<th>")
				out.WriteString(inlineMarkdown(cell))
				out.WriteString("</th>")
			}
			out.WriteString("</tr></thead><tbody>")
			for _, row := range rows {
				out.WriteString("<tr>")
				for col := range header {
					cell := ""
					if col < len(row) {
						cell = row[col]
					}
					out.WriteString("<td>")
					out.WriteString(inlineMarkdown(cell))
					out.WriteString("</td>")
				}
				out.WriteString("</tr>")
			}
			out.WriteString("</tbody></table>")
			continue
		}

		if level := headingLevel(trimmed); level > 0 {
			flushParagraph()
			flushQuote()
			flushList()
			out.WriteString("<h")
			out.WriteByte(byte('0' + level))
			out.WriteByte('>')
			out.WriteString(inlineMarkdown(strings.TrimSpace(trimmed[level:])))
			out.WriteString("</h")
			out.WriteByte(byte('0' + level))
			out.WriteByte('>')
			continue
		}

		if trimmed == "---" || trimmed == "***" {
			flushParagraph()
			flushQuote()
			flushList()
			out.WriteString("<hr>")
			continue
		}

		if strings.HasPrefix(trimmed, ">") {
			flushParagraph()
			flushList()
			quote = append(quote, strings.TrimSpace(strings.TrimPrefix(trimmed, ">")))
			continue
		}

		flushQuote()
		if tag, text, ok := listItem(trimmed); ok {
			if listTag != "" && listTag != tag {
				flushList()
			}
			listTag = tag
			listItems = append(listItems, inlineMarkdown(text))
			continue
		}

		flushList()
		paragraph = append(paragraph, trimmed)
	}

	flushParagraph()
	flushQuote()
	flushList()
	if inCode {
		flushCode()
	}

	return out.String()
}

func headingLevel(line string) int {
	level := 0
	for level < len(line) && line[level] == '#' {
		level++
	}
	if level == 0 || level > 6 || level >= len(line) || line[level] != ' ' {
		return 0
	}
	return level
}

func listItem(line string) (string, string, bool) {
	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
		return "ul", strings.TrimSpace(line[2:]), true
	}

	dot := strings.Index(line, ". ")
	if dot > 0 {
		if _, err := strconv.Atoi(line[:dot]); err == nil {
			return "ol", strings.TrimSpace(line[dot+2:]), true
		}
	}

	return "", "", false
}

func tableCells(line string) []string {
	cells := strings.Split(strings.TrimSpace(line), "|")
	if len(cells) > 0 && strings.TrimSpace(cells[0]) == "" {
		cells = cells[1:]
	}
	if len(cells) > 0 && strings.TrimSpace(cells[len(cells)-1]) == "" {
		cells = cells[:len(cells)-1]
	}
	for i := range cells {
		cells[i] = strings.TrimSpace(cells[i])
	}
	return cells
}

func tableSeparator(line string) bool {
	cells := tableCells(line)
	if len(cells) == 0 {
		return false
	}
	for _, cell := range cells {
		cell = strings.Trim(cell, ":")
		if cell == "" || strings.Trim(cell, "-") != "" {
			return false
		}
	}
	return true
}

func inlineMarkdown(value string) string {
	return inline(value, 0)
}

func inline(value string, depth int) string {
	var out strings.Builder

	for i := 0; i < len(value); {
		switch {
		case value[i] == '`':
			end := strings.IndexByte(value[i+1:], '`')
			if end < 0 {
				writeByte(&out, value[i])
				i++
				continue
			}
			out.WriteString("<code>")
			out.WriteString(html.EscapeString(value[i+1 : i+1+end]))
			out.WriteString("</code>")
			i += end + 2

		case strings.HasPrefix(value[i:], "**"):
			end := strings.Index(value[i+2:], "**")
			if end < 0 {
				out.WriteString("**")
				i += 2
				continue
			}
			out.WriteString("<strong>")
			out.WriteString(inline(value[i+2:i+2+end], depth))
			out.WriteString("</strong>")
			i += end + 4

		case value[i] == '[' && depth == 0:
			labelEnd := strings.Index(value[i:], "](")
			if labelEnd < 0 {
				writeByte(&out, value[i])
				i++
				continue
			}
			hrefStart := i + labelEnd + 2
			end := strings.IndexByte(value[hrefStart:], ')')
			if end < 0 {
				writeByte(&out, value[i])
				i++
				continue
			}
			href := docHref(value[hrefStart : hrefStart+end])
			out.WriteString(`<a href="`)
			out.WriteString(html.EscapeString(href))
			out.WriteString(`">`)
			out.WriteString(inline(value[i+1:i+labelEnd], depth+1))
			out.WriteString("</a>")
			i = hrefStart + end + 1

		case value[i] == '*':
			end := strings.IndexByte(value[i+1:], '*')
			if end < 0 {
				writeByte(&out, value[i])
				i++
				continue
			}
			out.WriteString("<em>")
			out.WriteString(inline(value[i+1:i+1+end], depth))
			out.WriteString("</em>")
			i += end + 2

		default:
			size := writeRune(&out, value[i:])
			i += size
		}
	}

	return out.String()
}

func writeByte(out *strings.Builder, value byte) {
	out.WriteString(html.EscapeString(string(rune(value))))
}

func writeRune(out *strings.Builder, value string) int {
	r, size := utf8.DecodeRuneInString(value)
	out.WriteString(html.EscapeString(string(r)))
	return size
}

func docHref(href string) string {
	if strings.Contains(href, "://") || !strings.HasSuffix(href, ".md") {
		return href
	}
	name := href[strings.LastIndex(href, "/")+1:]
	return "/docs/" + strings.TrimSuffix(name, ".md")
}
