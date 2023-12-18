package main

import (
	"Proyecto1_202111576/estructuras/ArbolAVL"
	"Proyecto1_202111576/estructuras/ColaPrioridad"
	"Proyecto1_202111576/estructuras/Listas"
	"Proyecto1_202111576/estructuras/MatrizDispersa"
	"fmt"
	"strconv"
)

var ListaDobleCircular *Listas.ListaDobleCircular = &Listas.ListaDobleCircular{Inicio: nil, Longitud: 0}
var ListaDoble *Listas.ListaDoble = &Listas.ListaDoble{Inicio: nil, Longitud: 0}
var colaPrioridad *ColaPrioridad.Cola = &ColaPrioridad.Cola{Primero: nil, Longitud: 0}
var matrizDispersa *MatrizDispersa.Matriz = &MatrizDispersa.Matriz{Raiz: &MatrizDispersa.NodoMatriz{PosX: -1, PosY: -1, Dato: &MatrizDispersa.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}
var arbolCursos *ArbolAVL.ArbolAVL = &ArbolAVL.ArbolAVL{Raiz: nil}

var loggeado_estudiante string = ""

func main() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. Inicio de Sesion")
		fmt.Println("2. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			MenuLogin()
		case 2:
			salir = true
		}
	}
}

func MenuLogin() {
	fmt.Print("\033[H\033[2J")
	usuario := ""
	password := ""
	fmt.Print("Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	if usuario == "ADMIN_202111576" && password == "admin" {
		fmt.Println("Administrador Inicio Sesion")
		MenuAdmin()
	} else if ListaDoble.Buscar(usuario, password) {
		loggeado_estudiante = usuario
		MenuEstudiantes()
	} else {
		fmt.Println("ERROR EN CREDENCIALES!!!!")
	}
}

func MenuAdmin() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. Carga de Estudiantes Tutores")
		fmt.Println("2. Carga de Estudiantes")
		fmt.Println("3. Cargar de Cursos") //falta
		fmt.Println("4. Control de Estudiantes")
		fmt.Println("5. Reportes") //Falta
		fmt.Println("6. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			CargaTutores()
		case 2:
			CargaEstudiantes()
		case 3:
			CargaCursos()
		case 4:
			ControlEstudiantes()
		case 5:
			MenuReportes()
		case 6:
			salir = true
		}

	}
}

func MenuReportes() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. Reporte de estudiantes.")
		fmt.Println("2. Reporte de tutores.")
		fmt.Println("3. Reporte de asignación.")
		fmt.Println("4. Reporte de cursos.")
		fmt.Println("5. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			ListaDoble.Reporte()
			fmt.Println("Reporte de estudiantes exitosamente graficado.")
		case 2:
			ListaDobleCircular.Reportev2()
			fmt.Println("Reporte de tutores exitosamente graficado.")
		case 3:
			matrizDispersa.Reporte("Asignaciones.jpg")
			fmt.Println("Reporte de asignacion exitosamente graficado.")
		case 4:
			arbolCursos.Graficar()
			fmt.Println("Reporte de cursos exitosamente graficado.")
		case 5:
			salir = true

		}
	}
}

func MenuEstudiantes() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("1. Ver Tutores Disponibles")
		fmt.Println("2. Asignarse Tutores")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Print("\033[H\033[2J")
			ListaDobleCircular.Mostrar()
		case 2:
			AsignarCurso()
		case 3:
			salir = true
		}
	}
}

func AsignarCurso() {
	opcion := ""
	salir := false
	for !salir {
		fmt.Println("Teclee el codigo del curso: ")
		fmt.Scanln(&opcion)
		//Iria el primer If del Arbol (pendiente)
		if arbolCursos.Busqueda(opcion) {
			if ListaDobleCircular.Buscar(opcion) {
				TutorBuscado := ListaDobleCircular.BuscarTutor(opcion)
				estudiante, err := strconv.Atoi(loggeado_estudiante)
				if err != nil {
					break
				}
				matrizDispersa.Insertar_Elemento(estudiante, TutorBuscado.Tutor.Carnet, opcion)
				fmt.Println("Se asigno Correctamente....")
				break
			} else {
				fmt.Println("No hay tutores para ese curso....")
				break
			}
		} else {
			fmt.Println("El curso no existe en el sistema")
			break
		}

	}
}

func CargaTutores() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	colaPrioridad.LeerCSV(ruta)
	fmt.Println("Se cargo a la Cola los tutores")

}

func CargaEstudiantes() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	ListaDoble.LeerCSV(ruta)
	fmt.Println("Se cargo los estudiantes")
}

func CargaCursos() {
	fmt.Print("\033[H\033[2J")
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	arbolCursos.LeerJson(ruta)
	fmt.Println("Se cargaron los cursos")
}

func ControlEstudiantes() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Print("\033[H\033[2J")
		colaPrioridad.PrimeroCola()
		fmt.Println("════════════════════")
		fmt.Println("1. Aceptar")
		fmt.Println("2. Rechazar")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			ListaDobleCircular.Agregar(colaPrioridad.Primero.Tutor.Carnet, colaPrioridad.Primero.Tutor.Nombre, colaPrioridad.Primero.Tutor.Curso, colaPrioridad.Primero.Tutor.Nota)
			colaPrioridad.Descolar()
		} else if opcion == 2 {
			colaPrioridad.Descolar()
		} else if opcion == 3 {
			salir = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}
