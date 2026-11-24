package status

import (
	constellation "constellation/lib"
	"constellation/lib/elements"
)

func Status404(r *constellation.Request) string {
	return page(
		"404",
		"Page not found",
		"This page is not part of the constellation. Check the address or start again from the home page.",
		requestedPath(r),
	)
}

func Status403(r *constellation.Request) string {
	return page(
		"403",
		"Access denied",
		"You do not have permission to view this page. Sign in with an authorized account or return home.",
		requestedPath(r),
	)
}

func page(code, title, message, path string) string {
	return elements.Div(constellation.WebComponent{
		Class: "page-shell mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pb-20 pt-8 md:w-[calc(100%_-_3rem)]",
		Children: []any{
			navigation(),
			elements.Section(constellation.WebComponent{
				Class: "status-page",
				Children: []any{
					elements.Span(constellation.WebComponent{Class: "status-code", Text: code}),
					elements.Span(constellation.WebComponent{Class: "eyebrow mb-3 block", Text: "Error " + code}),
					elements.H1(constellation.WebComponent{Class: "page-title mt-0 mb-4", Text: title}),
					elements.P(constellation.WebComponent{Class: "page-lead mx-auto", Text: message}),
					elements.Div(constellation.WebComponent{
						Class: "status-actions mt-8",
						Children: []any{
							link("/", "Back home", "button button-primary"),
							link("/docs", "Read the docs", "button button-secondary"),
						},
					}),
					pathRow(path),
				},
			}),
			footer(),
		},
	})
}

func navigation() string {
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
					link("/", "Home", "text-slate-400 hover:text-slate-50"),
					link("/docs", "Documentation", "text-slate-400 hover:text-slate-50"),
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
			elements.Span(constellation.WebComponent{Text: "Go · Tailwind CSS · Markdown"}),
		},
	})
}

func link(href, text, class string) string {
	return elements.A(constellation.WebComponent{
		Class:      class,
		Attributes: map[string]string{"href": href},
		Text:       text,
	})
}

func pathRow(path string) string {
	if path == "" {
		return ""
	}

	return elements.P(constellation.WebComponent{
		Class: "status-path mt-10",
		Children: []any{
			elements.Span(constellation.WebComponent{Text: "Requested path: "}),
			elements.Code(constellation.WebComponent{Class: "inline-code", Text: path}),
		},
	})
}

func requestedPath(r *constellation.Request) string {
	if r == nil || r.URL == nil {
		return ""
	}
	return r.URL.Path
}
