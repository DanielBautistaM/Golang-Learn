package rutas

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	texto := "si we"
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
		"texto": texto,
	}
	template, err := template.ParseFiles("template/ejemplo/parametros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, data)
	}
}
func Query(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"id":    r.URL.Query().Get("id"),
		"slug":  r.URL.Query().Get("slug"),
		"texto": r.URL.Query().Get("texto"),
	}
	template, err := template.ParseFiles("template/ejemplo/query.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, data)
	}
}

type Habilidad struct {
	Nombre string
}
type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

func Estructuras(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/ejemplo/estructura.html")
	if err != nil {
		panic(err)
	}
    
    habilidad1 := Habilidad{"Habilidad 1"}
    habilidad2 := Habilidad{"Habilidad 2"}
    habilidad3 := Habilidad{"Habilidad 3"}
    habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}

    datos := Datos{
        Nombre:      "Juan",
        Edad:        22,
        Perfil:      1,
        Habilidades: habilidades,
    }

    err = template.Execute(w, datos)
    if err != nil {
        panic(err)
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
