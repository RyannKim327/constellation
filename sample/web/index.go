package sample_web

import (
	constellation "constellation/src"
	"constellation/src/elements"
)

var Index = constellation.Endpoint{
	Path:       "/",
	Method:     constellation.GET,
	Permission: constellation.Unauthenticated,
	Handler: constellation.HandlerSetup{
		Web: handler,
	},
	Title: "Welcome to Constellation",
}

func handler(r *constellation.Request) string {
	features := []any{
		featureCard("01", "Declarative routes", "Describe an endpoint with an Endpoint value. Methods and paths are normalized before registration."),
		featureCard("02", "Composable elements", "Build markup with one helper per HTML tag. Text and attributes are escaped for you."),
		featureCard("03", "Standard library only", "No external dependencies. The stylesheet and documentation are embedded into the binary."),
	}

	routes := []any{
		routeRow("/docs", "Documentation index"),
		routeRow("/docs/getting-started", "Getting started"),
		routeRow("/docs/routing", "Routing guide"),
		routeRow("/docs/elements", "Elements guide"),
		routeRow("/docs/styling", "Styling guide"),
	}

	return elements.Div(constellation.WebComponent{
		Class: "page-shell mx-auto w-[calc(100%_-_2rem)] max-w-[1120px] pb-20 pt-8 md:w-[calc(100%_-_3rem)]",
		Children: []any{
			elements.Header(constellation.WebComponent{
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
							pageLink("/docs", "Documentation", "text-slate-400 hover:text-slate-50"),
							pageLink("/docs/getting-started", "Getting started", "text-slate-400 hover:text-slate-50"),
							pageLink("/docs/elements", "Elements", "text-slate-400 hover:text-slate-50"),
						},
					}),
				},
			}),
			elements.Main(constellation.WebComponent{
				Children: []any{
					elements.Header(constellation.WebComponent{
						Class: "hero flex flex-col items-start pb-16 pt-20",
						Children: []any{
							elements.Span(constellation.WebComponent{
								Class: "eyebrow block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400",
								Text:  "Built with Go and Tailwind CSS",
							}),
							elements.H1(constellation.WebComponent{
								Class: "hero-title my-5 mb-6 text-[clamp(2.75rem,7vw,5.5rem)] font-bold leading-[0.98] tracking-[-0.045em]",
								Text:  "Small framework. Clear conventions.",
							}),
							elements.P(constellation.WebComponent{
								Class: "hero-lead mb-8 max-w-2xl text-[clamp(1rem,2vw,1.2rem)] text-slate-400",
								Text:  "Constellation is a small web framework for Go. Declare an endpoint, return HTML from a handler, and the framework wraps it in a page.",
							}),
							elements.Div(constellation.WebComponent{
								Class: "hero-actions flex flex-wrap gap-3",
								Children: []any{
									pageLink("/docs", "Read the docs", "button button-primary inline-flex items-center justify-center rounded-lg border border-transparent bg-orange-500 px-[1.1rem] py-[0.7rem] text-sm font-semibold text-gray-900 transition hover:-translate-y-px hover:bg-orange-400"),
									pageLink("/docs/elements", "Browse elements", "button button-secondary inline-flex items-center justify-center rounded-lg border border-white/10 bg-transparent px-[1.1rem] py-[0.7rem] text-sm font-semibold text-slate-50 transition hover:-translate-y-px hover:border-orange-400"),
								},
							}),
						},
					}),
					elements.Section(constellation.WebComponent{
						Class:    "feature-grid mb-16 grid grid-cols-[repeat(auto-fit,minmax(15rem,1fr))] gap-4",
						Children: features,
					}),
					elements.Section(constellation.WebComponent{
						Class: "split-section mb-12 grid grid-cols-1 items-center gap-8 border-y border-white/10 py-14 md:grid-cols-2",
						Children: []any{
							elements.Div(constellation.WebComponent{
								Class: "prose-block",
								Children: []any{
									elements.Span(constellation.WebComponent{
										Class: "section-label block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400",
										Text:  "Elements",
									}),
									elements.H2(constellation.WebComponent{
										Class: "mt-[0.65rem] mb-4 text-[clamp(1.5rem,3vw,2.1rem)] tracking-[-0.02em]",
										Text:  "Compose pages in Go",
									}),
									elements.P(constellation.WebComponent{
										Class: "mb-4 text-slate-400",
										Text:  "Every HTML tag has a helper. Nest components with Children and let the renderer escape text and attributes.",
									}),
									pageLink("/docs/elements", "Read the elements guide →", "text-link font-semibold text-orange-400 hover:underline"),
								},
							}),
							elements.Pre(constellation.WebComponent{
								Class: "code-block overflow-x-auto rounded-[0.9rem] border border-white/10 bg-[#07070a] p-[1.4rem] text-[0.82rem] leading-[1.65] text-slate-200",
								Children: []any{
									elements.Code(constellation.WebComponent{
										Text: `elements.Div(constellation.WebComponent{
    Class: "page-shell",
    Children: []any{
        elements.H1(constellation.WebComponent{Text: "Hello"}),
        elements.P(constellation.WebComponent{Text: "Built with Go."}),
    },
})`,
									}),
								},
							}),
						},
					}),
					elements.Section(constellation.WebComponent{
						Class: "route-list flex flex-col",
						Children: []any{
							elements.Span(constellation.WebComponent{
								Class: "section-label mb-3 block text-[0.72rem] font-bold uppercase tracking-[0.18em] text-orange-400",
								Text:  "Available pages",
							}),
							routes[0],
							routes[1],
							routes[2],
							routes[3],
							routes[4],
						},
					}),
				},
			}),
			elements.Footer(constellation.WebComponent{
				Class: "page-footer mt-16 flex flex-col items-start justify-between gap-4 border-t border-white/10 pt-6 text-[0.8rem] text-slate-400 md:flex-row md:items-center",
				Children: []any{
					elements.Span(constellation.WebComponent{Text: "Constellation"}),
					elements.Span(constellation.WebComponent{Text: "Go · Tailwind CSS · Markdown"}),
				},
			}),
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

func featureCard(index, title, body string) string {
	return elements.Article(constellation.WebComponent{
		Class: "feature-card rounded-[0.9rem] border border-white/10 bg-white/[0.04] p-6",
		Children: []any{
			elements.Span(constellation.WebComponent{
				Class: "feature-index text-xs font-bold tracking-[0.12em] text-orange-400",
				Text:  index,
			}),
			elements.H2(constellation.WebComponent{
				Class: "mt-5 mb-2 text-[1.1rem]",
				Text:  title,
			}),
			elements.P(constellation.WebComponent{
				Class: "m-0 text-[0.9rem] text-slate-400",
				Text:  body,
			}),
		},
	})
}

func routeRow(href, title string) string {
	return elements.A(constellation.WebComponent{
		Class:      "flex items-center justify-between border-b border-white/10 py-4 text-[0.95rem] text-slate-50 hover:text-orange-400",
		Attributes: map[string]string{"href": href},
		Children: []any{
			elements.Span(constellation.WebComponent{Text: title}),
			elements.Span(constellation.WebComponent{Text: "→"}),
		},
	})
}
