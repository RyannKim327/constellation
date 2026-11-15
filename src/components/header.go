package components

import (
	"strings"

	"github.com/RyannKim327/constellation/config/utils"
)

/*
 * INFO: This is an example of reusable component
 * This may also use as "var" or "func"
 * Use "var" if you don't pass any parameters
 * Use func if you pass parameters
 */

type navs struct {
	Name string
	Url  string
}

func Header(currentPage string) string {
	_endpoints := []navs{
		{
			Url:  "about",
			Name: "About",
		},
		{
			Url:  "docs",
			Name: "Documentation",
		},
		{
			Url:  "samples",
			Name: "Samples",
		},
		{
			Url:  "api",
			Name: "API",
		},
	}

	if strings.HasPrefix(currentPage, "/") {
		currentPage = currentPage[1:]
	}
	navigation := ""

	for _, endpoint := range _endpoints {
		// INFO: This will remove the active tab
		if endpoint.Url == currentPage {
			navigation += utils.HTParser(`
				<a class="border-b-2 border-b-white border-b-solid px-2" href="/{url}">{name}</a>
				`, utils.PARAMS{
				"url":  endpoint.Url,
				"name": endpoint.Name,
			})
		} else {
			navigation += utils.HTParser(`
				<a class="px-2" href="/{url}">{name}</a>
				`, utils.PARAMS{
				"url":  endpoint.Url,
				"name": endpoint.Name,
			})
		}
	}

	return utils.HTParser(`
		<nav class="flex flex-col items-start justify-between gap-4 pt-4 pb-8 text-sm text-slate-400 md:flex-row md:items-center [&_a:hover]:text-slate-50">
			<a href="/">← Back to main</a>
			<div class="flex flex-wrap gap-5">
				{navigation}
			</div>
		</nav>
	`, utils.PARAMS{
		"navigation": navigation,
	})
}
