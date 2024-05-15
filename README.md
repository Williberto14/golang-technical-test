# Golang Technical Test

Este proyecto es una aplicación de prueba técnica construida con Go. Utiliza una arquitectura limpia, que separa las responsabilidades en diferentes capas para facilitar el mantenimiento y la escalabilidad del código.

## Arquitectura

La arquitectura del proyecto se divide en varias capas:

* [`cmd`]: Contiene el punto de entrada principal de la aplicación (`main.go`).
* [`config`]: Contiene la configuración de la aplicación (`config.go` y [`config.yml`]).
* [`database`]: Contiene la lógica para la conexión a la base de datos (`database.go`).
* [`internal`]: Contiene la lógica principal de la aplicación, dividida en varias subcapas:
  * `delivery`: Define cómo se entregan los datos al usuario. En este caso, se utiliza HTTP.
  * `domain`: Define las estructuras de datos principales utilizadas en la aplicación.
  * `repository`: Contiene la lógica para interactuar con la base de datos.
  * `usecase`: Contiene la lógica de negocio de la aplicación.
* [`middlewares`]: Contiene los middlewares utilizados en la aplicación, como la autenticación JWT.
* [`utils`]: Contiene funciones de utilidad que se utilizan en toda la aplicación.

## Tecnologías

* **Go** : El lenguaje de programación principal utilizado para construir la aplicación.
* **Gin** : Un marco web HTTP de alto rendimiento que proporciona una forma sencilla de crear aplicaciones web en Go.
* **JWT** : Se utiliza para la autenticación y la generación de tokens.
* **Viper** : Se utiliza para la gestión de la configuración.

## Cómo ejecutar el proyecto

1. Asegúrate de tener Go instalado en tu máquina.
2. Clona el repositorio en tu máquina local.
3. Navega hasta el directorio del proyecto.
4. Ejecuta `go run cmd/main.go` para iniciar la aplicación.

Por favor, consulta el archivo [`config.yml`] para configurar la aplicación según tus necesidades.

## Base de datos

```bash
docker run --name golang-technical-test -p 127.0.0.1:3306:3306 -e MYSQL_ROOT_PASSWORD=qwerty -e MYSQL_DATABASE=golang_technical_test -d mariadb:latest
```
