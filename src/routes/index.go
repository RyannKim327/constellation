package routes

import (
	"encoding/json"
	"net/http"

	"constellation/config/middleware"
	"constellation/config/utils"

	"github.com/gin-gonic/gin"
)

var Index = middleware.Route{
	Method:     utils.METHOD_GET,
	Path:       "/",
	Handler:    layout,
	Permission: utils.PERMISSION_ALL,
	Title:      "Hello Constellation",
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func layout(ctx *gin.Context) (string, utils.PARAMS) {
	name := ctx.DefaultQuery("name", "Test")

	// TODO: This is how to fetch data
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var user []User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	// TODO: This is a sample of mapping
	data := ""

	for _, d := range user {
		data += "<li>" +
			d.Name +
			"</li>"
	}

	return (utils.HTParser(`
		<div class="flex flex-col py-[5rem] px-[10rem] w-full h-full">
			<span class="self-start text-slate-500">Back to main</span>
			<div class="bg-slate-950 w-full p-2 rounded">
				<h1>Hello {name}</h1>
				<div>
					<ul>
						{data}
					</ul>
				</div>
			</div>
		</div>
	`,
			utils.PARAMS{
				"data": data,
			})), utils.PARAMS{
			"name": name,
		}
}
