package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"main.go/conectar"
	"main.go/modelos"
)

func Listar() {
	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes ORDER BY id DESC;"
	datos, err := conectar.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer conectar.CerrarConexion()

	fmt.Println("--------------DATOS--------------")
	fmt.Println("----------------------------")
	for datos.Next() {
		var data modelos.Cliente
		err := datos.Scan(&data.Id, &data.Nombre, &data.Correo, &data.Telefono)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Id: %v | Nombre %v | Email: %v | Telefono %v \n", data.Id, data.Nombre, data.Correo, data.Telefono)
		fmt.Println("----------------------------")
	}
}

func ListarPorIDchan(id int) {
	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes WHERE id = ?;"
	datos, err := conectar.Db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer conectar.CerrarConexion()

	fmt.Println("--------------DATO POR ID--------------")
	fmt.Println("----------------------------")
	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Id: %v | Nombre %v | Email: %v | Telefono %v \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
		fmt.Println("----------------------------")
	}
}

func Insertar(cliente modelos.Cliente) {
	conectar.Conectar()
	sql := "insert into clientes values(null,?,?,?);"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("se creo el registro")
}

func Editar(cliente modelos.Cliente, id int) {
	conectar.Conectar()
	sql := "Update clientes set nombre=?, correo=?, telefono=? where id = ?;"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se edito good")
}

func Eliminar(id int) {
	conectar.Conectar()
	sql := "DELETE FROM clientes WHERE id=?;"
	_, err := conectar.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("se elimino bien we")
}

var ID int
var nombre, correo, telefono string

func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Seleccione una opcion \n")
	fmt.Println("1.Listar cliente \n")
	fmt.Println("2.Listar X ID\n")
	fmt.Println("3.Crear cliente \n")
	fmt.Println("4.Editar Cliente \n")
	fmt.Println("5.Eliminar Cliente \n ")
	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				Listar()
				return
			}
			if scanner.Text() == "2" {
				fmt.Println("Ingrese el Id del cliente: \n")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				ListarPorIDchan(ID)
				return
			}
			if scanner.Text() == "3" {
				fmt.Println("Ingrese el nombre: ")
				if scanner.Scan() {
					nombre = scanner.Text() // ✅ Se asigna correctamente el nombre
				}

				fmt.Println("Ingrese el Correo: ")
				if scanner.Scan() {
					correo = scanner.Text() // ✅ Se asigna correctamente el correo
				}

				fmt.Println("Ingrese el telefono: ")
				if scanner.Scan() {
					telefono = scanner.Text() // ✅ Se asigna correctamente el teléfono
				}

				cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Insertar(cliente)
				Ejecutar()
				continue
			}

			if scanner.Text() == "4" {
				fmt.Println("Ingrese ID cliente: ")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}

				fmt.Println("Ingrese el nombre: ")
				var nombre, correo, telefono string
				if scanner.Scan() {
					nombre = scanner.Text()
				}

				fmt.Println("Ingrese el Correo: ")
				if scanner.Scan() {
					correo = scanner.Text()
				}

				fmt.Println("Ingrese el telefono: ")
				if scanner.Scan() {
					telefono = scanner.Text()
				}

				cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Editar(cliente, ID) // 🔹 Aquí se usa Editar() en vez de Insertar()
				Ejecutar()
				continue
			}
			if scanner.Text() == "5" {
				fmt.Println("Ingresa el ID del cliente: \n")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
					Eliminar(ID)
					return
				}
			}

		}
	}

}
