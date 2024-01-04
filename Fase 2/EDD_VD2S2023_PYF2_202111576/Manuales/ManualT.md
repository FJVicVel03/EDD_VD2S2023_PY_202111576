## Proyecto único, Estructura de Datos, Vacaciones 2do semestre
### Manual Técnico

Explicación del código de la aplicación realizada en web con React + Vite para el proyecto único del curso Estructura de Datos.

1. El programa toma su inicio desde el archivo main, donde se mandan a llamar todas las funciones que realizan las diferentes acciones en la página web, que estas funciones a su vez se conectarán por Vite + React para darles su respectiva funcionalidad.

![1](/Manuales/imagenes/1.PNG)
2. Se empieza por las estructuras, estan son varias y diversas para cumplir con los requisitos del proyecto. Empezando por el ArbolB. Funcionando con el siguiente código:
![2](/Manuales/imagenes/2.PNG)
![2](/Manuales/imagenes/2.1.PNG)
![2](/Manuales/imagenes/2.2.PNG)
![2](/Manuales/imagenes/2.3.PNG)
Estas funcionan a traves de una lista enlazada, por nodos y una rama. 
3. Se realiza una estructura de arbol merkle, esta tiene como fin mostrar los libros autorizados por el administrador al tutor, y así mostrarlos en un grafo de manera ordenada.
![3](/Manuales/imagenes/3.PNG)
![3](/Manuales/imagenes/3.1.PNG)
Este arbol funciona con sus propias funciones, como tambien del uso de un nodo como es costumbre. En su función principal se realiza lo más importante, como por ejemplo, la inserción o graficar lo que se obtiene.

4. Se continúa con los grafos, los cuales tienen la función de graficar los cursos que tiene el alumno, o mejor dicho, que tienen los alumnos. Así se mostrarán de manera ordenada.
![4](/Manuales/imagenes/4.PNG)
![4](/Manuales/imagenes/4.1.PNG)
Estos grafos funcionan a partir de inserción de datos en columnas y filas, además de apoyarse usando un nodo como es costumbre.

5. Le siguen las peticiones, que como dice, realizan esto, las peticiones del usuario manejando el programa desde la página, desde almacenar un alumno o tutor, además de los cursos. Todo eso pasa a través de las peticiones con el siguiente código.
![5](/Manuales/imagenes/5.PNG)
![5](/Manuales/imagenes/5.1.PNG)

6. Finalizando el backend, es con la tabla Hash, esta almacena los alumnos registrados en un arreglo, el cual posteriormente se mostrará en una tabla dentro de la página web, en la vista de administrador.

![6](/Manuales/imagenes/6.PNG)
![6](/Manuales/imagenes/6.1.PNG)

### Comenzando con el frontend

1. El frontend se realizó con Vite + React usando Javascript, esto instalando sus respectivas dependencias y servicios, también se utilizó Bootstrap. Para la incialización de la App se usó de esta manera:
![10](/Manuales/imagenes/10.PNG)
![10](/Manuales/imagenes/10.1.PNG)

2. Ahora para trabajar la parte de administrador, estudiante y tutor se separó en diferentes carpetas para mantener el orden, Por su puesto, se requiere de un login funcional, el cual se realizó de la siguiente manera:

![11](/Manuales/imagenes/11.PNG)
![11](/Manuales/imagenes/11.1.PNG)
![11](/Manuales/imagenes/11.2.PNG)

3. Para la parte administrativa se utilizaron 4 diferentes "paneles", la carga de archivos, control de libros, la pagina principal, los reportes y la tabla de alumnos visibles y cargados. 

a. Carga de archivos
![12](/Manuales/imagenes/12.PNG)
b. Control de Libros
![12](/Manuales/imagenes/12.1.PNG)
c. Administrador (pagina principal)
![12](/Manuales/imagenes/12.2.PNG)
d. Reportes
![12](/Manuales/imagenes/12.3.PNG)
e. Tabla de alumnos
![12](/Manuales/imagenes/12.4.PNG)

4. Pasando a la parte de tutores, se utilizaron 2 paneles para su funcionamiento, Carga de publicaciones y carga de libros. Estos se hicieron de esta manera:

a. Carga de publicaciones
![13](/Manuales/imagenes/13.PNG)
b. Carga de libros
![13](/Manuales/imagenes/13.1.PNG)

5. Ahora finalizamos con el apartado de Estudiante, se utilizaron 3 paneles para su funcionamiento, Cursos de los estudiantes, Pagina principal del estudiante y publicaciones para estudiantes. Se realizaron de la siguiente forma:

a. Cursos asignados de estudiante:
![14](/Manuales/imagenes/14.PNG)
b. Pagina principal de estudiante: 
![14](/Manuales/imagenes/14.1.PNG)
c. Publicación de estudiante:
![14](/Manuales/imagenes/14.2.PNG)