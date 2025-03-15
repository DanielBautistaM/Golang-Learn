package rutas

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/ejemplo/home.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)
	}
}

func Nosotros(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/ejemplo/nosotros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)
	}
}
func Parametros(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/ejemplo/parametros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)
	}
}
func Query(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/ejemplo/query.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)
	}
}

// func Home(response http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintln(response, "hola zzz")
// }

// func Nosotros(response http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintln(response, "sobre ssss zzz")
// }

// func Parametros(response http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)
// 	fmt.Fprintln(response, "ID = "+vars["id"]+" | SLUG = "+vars["slug"]+"")
// }

// func Query(response http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintln(response, request.URL)
// 	fmt.Fprintln(response, request.URL.RawQuery)
// 	fmt.Fprintln(response, request.URL.Query())
// 	id:=request.URL.Query().Get("id")
// 	slug:=request.URL.Query().Get("slug")
// 	fmt.Fprintln(response, id)
// 	fmt.Fprintln(response, slug)

// }
