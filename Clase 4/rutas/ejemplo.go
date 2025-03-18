package rutas

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Home maneja la ruta principal, renderizando la plantilla home.html
func Home(w http.ResponseWriter, r *http.Request) {
	// ParseFiles lee y parsea el archivo de template
	template, err := template.ParseFiles("template/ejemplo/home.html")
	if err != nil {
		panic(err)
	}
	// Execute renderiza el template con los datos proporcionados (nil en este caso)
	template.Execute(w, nil)
}

// Nosotros maneja la ruta /notosotros
func Nosotros(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/ejemplo/nosotros.html")
	if err != nil {
		panic(err)
	}
	template.Execute(w, nil)
}

// Parametros maneja rutas con parámetros en la URL
// Ejemplo: /parametros/123/mi-slug
func Parametros(w http.ResponseWriter, r *http.Request) {
	// mux.Vars obtiene los parámetros de la URL
	vars := mux.Vars(r)
	template, err := template.ParseFiles("template/ejemplo/parametros.html")
	if err != nil {
		panic(err)
	}
	// Pasamos los parámetros al template como un map
	template.Execute(w, vars)
}

// Query maneja rutas con parámetros y query strings
// Ejemplo: /query/123/mi-slug/hola?extra=valor
func Query(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Creamos un map para almacenar todos los datos
	data := make(map[string]string)
	
	// Copiamos los parámetros de la URL al map
	for key, value := range vars {
		data[key] = value
	}
	
	// Agregamos los query parameters al map
	data["extra"] = r.URL.Query().Get("extra")
	
	template, err := template.ParseFiles("template/ejemplo/query.html")
	if err != nil {
		panic(err)
	}
	template.Execute(w, data)
}

// Habilidad representa una habilidad con su nombre
type Habilidad struct {
	Nombre string
}

// Datos representa la estructura de datos principal
// que se usa en el template estructura.html
type Datos struct {
	Nombre      string      // Nombre de la persona
	Edad        int         // Edad de la persona
	Perfil      int         // Tipo de perfil
	Habilidades []Habilidad // Lista de habilidades
}

// Estructuras demuestra el uso de structs en templates
func Estructuras(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/ejemplo/estructura.html")
	if err != nil {
		panic(err)
	}

	// Creamos algunas habilidades de ejemplo
	habilidad1 := Habilidad{"Habilidad 1"}
	habilidad2 := Habilidad{"Habilidad 2"}
	habilidad3 := Habilidad{"Habilidad 3"}
	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}

	// Creamos una instancia de Datos con valores de ejemplo
	datos := Datos{
		Nombre:      "Juan",
		Edad:        22,
		Perfil:      1,
		Habilidades: habilidades,
	}

	// Renderizamos el template con nuestros datos
	template.Execute(w, datos)
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
