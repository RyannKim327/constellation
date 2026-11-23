package docs

import (
	"embed"
	"sort"
	"strings"

	constellation "constellation/src"
	"constellation/src/elements"
)

type document struct {
	Slug     string
	Title    string
	Summary  string
	Markdown string
}

var documents []document

func Register(fsys embed.FS) {
	documents = loadDocuments(fsys)

	constellation.AddEndpoint(constellation.Endpoint{
		Path:       "/docs",
		Method:     constellation.GET,
		Permission: constellation.Unauthenticated,
		Title:      "Documentation · Constellation",
		Handler: constellation.HandlerSetup{
			Web: indexPage,
		},
	})

	for _, doc := range documents {
		doc := doc
		constellation.AddEndpoint(constellation.Endpoint{
			Path:       "/docs/" + doc.Slug,
			Method:     constellation.GET,
			Permission: constellation.Unauthenticated,
			Title:      doc.Title + " · Constellation",
			Handler: constellation.HandlerSetup{
				Web: func(r *constellation.Request) string {
					return documentPage(doc)
				},
			},
		})
	}
}

func loadDocuments(fsys embed.FS) []document {
	entries, err := fsys.ReadDir("docs")
	if err != nil {
		panic(err)
	}

	var docs []document
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		source, err := fsys.ReadFile("docs/" + entry.Name())
		if err != nil {
			panic(err)
		}

		content := string(source)
		title := docTitle(content)
		docs = append(docs, document{
			Slug:     strings.TrimSuffix(entry.Name(), ".md"),
			Title:    title,
			Summary:  docSummary(content),
			Markdown: stripTitle(content, title),
		})
	}

	sort.Slice(docs, func(i, j int) bool {
		return docs[i].Slug < docs[j].Slug
	})

	return docs
}

func docTitle(source string) string {
	for _, line := range strings.Split(source, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}
	return "Documentation"
}

func stripTitle(source, title string) string {
	lines := strings.SplitN(source, "\n", 2)
	if len(lines) == 2 && strings.TrimSpace(lines[0]) == "# "+title {
		return lines[1]
	}
	return source
}

func docSummary(source string) string {
	lines := strings.Split(source, "\n")
	var parts []string

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i == 0 && strings.HasPrefix(line, "#") {
			continue
		}
		if line == "" {
			if len(parts) > 0 {
				break
			}
			continue
		}
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "|") || strings.HasPrefix(line, "```") {
			if len(parts) > 0 {
				break
			}
			continue
		}
		parts = append(parts, line)
	}

	summary := strings.NewReplacer("`", "", "**", "", "*", "", "_", "").Replace(strings.Join(parts, " "))
	if len(summary) > 160 {
		summary = summary[:157]
		if space := strings.LastIndex(summary, " "); space > 0 {
			summary = summary[:space]
		}
		summary = strings.TrimSpace(summary) + "..."
	}
	return summary
}

func indexPage(r *constellation.Request) string {
	cards := make([]any, 0, len(documents))
	for _, doc := range documents {
		cards = append(cards, elements.A(constellation.WebComponent{
			Class:      "docs-card flex flex-col gap-2 rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-[1.4rem] transition duration-150 hover:-translate-y-0.5 hover:border-orange-400/50",
			Attributes: map[string]string{"href": "/docs/" + doc.Slug},
			Children: []any{
				elements.Span(constellation.WebComponent{Class: "docs-card-title text-[1.05rem] font-bold", Text: doc.Title}),
				elements.Span(constellation.WebComponent{Class: "docs-card-summary text-[0.88rem] text-slate-400", Text: doc.Summary}),
				elements.Span(constellation.WebComponent{Class: "docs-card-link mt-auto pt-2 text-[0.8rem] font-semibold text-orange-400", Text: "Read document →"}),
			},
		}))
	}

	return elements.Div(constellation.WebComponent{
		Class: "page-shell mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pb-20 pt-8 md:w-[calc(100%_-_3rem)]",
		Children: []any{
			navigation("/docs"),
			elements.Header(constellation.WebComponent{
				Class: "page-header pt-12 pb-10",
				Children: []any{
					elements.Span(constellation.WebComponent{Class: "eyebrow block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400", Text: "Documentation"}),
					elements.H1(constellation.WebComponent{Class: "page-title mt-3 mb-4 text-[clamp(2rem,5vw,3.5rem)] leading-[1.08] tracking-[-0.035em]", Text: "Constellation handbook"}),
					elements.P(constellation.WebComponent{
						Class: "page-lead m-0 max-w-[48rem] text-[1.02rem] text-slate-400",
						Text:  "Short guides for building with Constellation. Every page below is written in Markdown under docs/ and embedded into the server at build time.",
					}),
				},
			}),
			elements.Section(constellation.WebComponent{Class: "docs-grid mb-10 grid grid-cols-[repeat(auto-fit,minmax(17rem,1fr))] gap-4", Children: cards}),
			elements.Div(constellation.WebComponent{
				Class: "callout rounded-[0.9rem] border border-orange-400/30 bg-orange-400/[0.07] px-[1.4rem] py-[1.2rem]",
				Children: []any{
					elements.P(constellation.WebComponent{Class: "callout-title m-0 mb-[0.35rem] text-[0.8rem] font-bold uppercase tracking-[0.08em] text-orange-400", Text: "Markdown is the source"}),
					elements.P(constellation.WebComponent{
						Class: "callout-body m-0 text-[0.9rem] text-slate-400",
						Text:  "Edit a file in docs/ and rebuild to update the site. No documentation is read from disk at request time.",
					}),
				},
			}),
			footer(),
		},
	})
}

func documentPage(doc document) string {
	sidebar := []any{
		elements.P(constellation.WebComponent{Class: "docs-sidebar-title m-0 mb-3 text-[0.7rem] font-bold uppercase tracking-[0.16em] text-slate-400", Text: "Documents"}),
	}

	links := make([]any, 0, len(documents))
	for _, item := range documents {
		class := "docs-sidebar-link block rounded px-[0.6rem] py-[0.45rem] text-[0.85rem] text-slate-400 hover:text-slate-50"
		if item.Slug == doc.Slug {
			class += " docs-sidebar-link-active bg-orange-400/[0.12] font-semibold text-orange-400"
		}
		links = append(links, elements.Li(constellation.WebComponent{
			Children: []any{pageLink("/docs/"+item.Slug, item.Title, class)},
		}))
	}
	sidebar = append(sidebar, elements.Ul(constellation.WebComponent{
		Class:    "docs-sidebar-list m-0 flex list-none flex-col gap-0.5 p-0",
		Children: links,
	}))

	return elements.Div(constellation.WebComponent{
		Class: "page-shell mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pb-20 pt-8 md:w-[calc(100%_-_3rem)]",
		Children: []any{
			navigation("/docs"),
			elements.Div(constellation.WebComponent{
				Class: "docs-layout grid grid-cols-1 items-start gap-12 lg:grid-cols-[minmax(13rem,16rem)_minmax(0,1fr)]",
				Children: []any{
					elements.Aside(constellation.WebComponent{Class: "docs-sidebar rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-[1.2rem] lg:sticky lg:top-8", Children: sidebar}),
					elements.Article(constellation.WebComponent{
						Class: "docs-article",
						Children: []any{
							elements.Header(constellation.WebComponent{
								Class: "docs-article-header mb-8 border-b border-white/10 pb-7",
								Children: []any{
									elements.Span(constellation.WebComponent{Class: "eyebrow block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400", Text: "Documentation"}),
									elements.H1(constellation.WebComponent{Class: "mt-[0.65rem] mb-3 text-[clamp(1.9rem,4vw,2.8rem)] tracking-[-0.03em]", Text: doc.Title}),
									elements.P(constellation.WebComponent{Class: "page-lead m-0 max-w-[48rem] text-[1.02rem] text-slate-400", Text: doc.Summary}),
								},
							}),
							elements.Div(constellation.WebComponent{
								Class:    "markdown",
								Children: []any{markdownToHTML(doc.Markdown)},
							}),
						},
					}),
				},
			}),
			footer(),
		},
	})
}

func navigation(active string) string {
	class := "text-slate-400 hover:text-slate-50"
	if active == "/docs" {
		class = "text-link font-semibold text-orange-400 hover:text-orange-300"
	}

	return elements.Nav(constellation.WebComponent{
		Class: "page-nav flex flex-col items-start justify-between gap-4 pb-8 pt-4 text-sm md:flex-row md:items-center",
		Children: []any{
			elements.A(constellation.WebComponent{
				Class:      "brand flex items-center gap-2 text-base font-bold uppercase tracking-[0.08em] text-slate-50",
				Attributes: map[string]string{"href": "/"},
				Children: []any{
					elements.Logo("h-7 w-7 shrink-0"),
					elements.Span(constellation.WebComponent{Text: "Constellation"}),
				},
			}),
			elements.Div(constellation.WebComponent{
				Class: "page-nav-links flex flex-wrap gap-5",
				Children: []any{
					pageLink("/", "Home", "text-slate-400 hover:text-slate-50"),
					pageLink("/docs", "Documentation", class),
					pageLink("/docs/elements", "Elements", "text-slate-400 hover:text-slate-50"),
				},
			}),
		},
	})
}

func footer() string {
	return elements.Footer(constellation.WebComponent{
		Class: "page-footer mt-16 flex flex-col items-start justify-between gap-4 border-t border-white/10 pt-6 text-[0.8rem] text-slate-400 md:flex-row md:items-center",
		Children: []any{
			elements.Span(constellation.WebComponent{Text: "Constellation"}),
			elements.Span(constellation.WebComponent{Text: "Go · Markdown · Standard library"}),
		},
	})
}

func pageLink(href, text, class string) string {
	return elements.A(constellation.WebComponent{
		Class:      class,
		Attributes: map[string]string{"href": href},
		Text:       text,
	})
}
