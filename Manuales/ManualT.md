## Proyecto único, Estructura de Datos, Vacaciones 2do semestre
### Manual Técnico

Explicación del código de la aplicación realizada en consola para el proyecto único del curso Estructura de Datos.

1. La función principal se compone de un menú sencillo de inicio de sesión realizado con un switch.
![1](/Manuales/Imagenes/1.jpg)
![2](/Manuales/Imagenes/1.1.jpg)

2. Tras iniciar sesión, se mandará a dos diferentes funciones dependiendo del tipo de usuario que se cargó. Al principio solo podremos iniciar sesión como administrador. El menú administrador está codificado de manera sencilla de igual forma con un switch, pero con funciones escenciales para el correcto funcionamiento del programa.
![3](/Manuales/Imagenes/2.jpg)
3. Si nos vamos a la primera opción con la carga de tutores, esta nos lanzará a la función que carga los ya mencionados. Este los carga en una Cola de Prioridad, para posteriormente ser aceptados, lanzados a una Lista Doble Circular pues los carga desde un CSV, además de crear un archivo DOT para la posterior gráfica realizada por graphviz.
![4](/Manuales/Imagenes/3.jpg)
![5](/Manuales/Imagenes/3.1.jpg)
![6](/Manuales/Imagenes/3.2.jpg)
![7](/Manuales/Imagenes/3.3.jpg)
![8](/Manuales/Imagenes/3..4.jpg)
4. Si nos vamos por la segunda opción con la carga de carga de estudiantes, esta nos lanzará a la función de los ya mencionados. Este los cargará en una Lista Doble, pues los carga desde un archivo CSV y crea un archivo DOT para la posterior grafica que realiza Graphviz.
![9](/Manuales/Imagenes/4.jpg)
![10](/Manuales/Imagenes/4.1.jpg)
![11](/Manuales/Imagenes/4.2.jpg)
![12](/Manuales/Imagenes/4.3.jpg)
5. Si nos vamos a la tercera opción, con la carga de los cursos, esta nos lanzará a la opción de carga de estos. Este los cargará en una Matriz Dispersa, los cargará desde un archivo JSON y además, creará un archivo DOT para su posterior gráfica realizada por Graphviz.
![13](/Manuales/Imagenes/5.jpg)
![13](/Manuales/Imagenes/5.1.jpg)
![13](/Manuales/Imagenes/5.2.jpg)
![13](/Manuales/Imagenes/5.3.jpg)
![13](/Manuales/Imagenes/5.4.jpg)
![13](/Manuales/Imagenes/5.5.jpg)
6. Si nos vamos a la cuarta opción, con el control de estudiantes, esta nos permitirá aceptar o rechazar a los tutores dependiendo de su prioridad, si es baja o alta, tendremos la libertad de aceptarlos o rechazarlos. Los tutores se mandaran a una Cola de Prioridad, y si son aceptados, a la Lista Doble Circular.
![14](/Manuales/Imagenes/6.jpg)
![14](/Manuales/Imagenes/6.1.jpg)
![14](/Manuales/Imagenes/6.2.jpg)
![14](/Manuales/Imagenes/6.3.jpg)
7. La última opción, reportes, nos generará los reportes que previamente fueron almacenados en archivos DOT, se nos presentará un menú sencillo para la graficación en forma de reporte de cada una de las opciones anteriores que tuvimos.
![15](/Manuales/Imagenes/7.jpg)
8. Ahora, del lado de estudiantes, tendremos las opciones de Ver tutores, y Asignar cursos.
Los cursos que se le presenten a los estudiantes, están estrechamente ligados con los que se cargaron a partir del archivo JSON, y los tutores que se visualicen, son los que fueron aceptados durante el control de administrador.
![16](/Manuales/Imagenes/8.jpg)
![16](/Manuales/Imagenes/8.1.jpg)
9. Así termina el correcto funcionamiento del programa.