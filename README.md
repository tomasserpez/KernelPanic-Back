# SpaceTraders

En este repositorio se encuentra construido el backend de nuestro cliente http para poder acceder a la API de SpaceTraders.

## Requisitos

- Go v1.23.2^
- SQLite3
- Git

## Instalacion

1. Clonar el repositorio

Usando git con SSH
```bash
git clone git@github.com:tomasserpez/KernelPanic-Back.git
cd KernelPanic-Back
```

Usando git con HTTP
```bash
git clone https://github.com/tomasserpez/KernelPanic-Back.git
cd KernelPanic-Back
```
2. Ejecutamos la aplicación:
```bash
go run main.go
```

Esto iniciará la aplicación usando el puerto 8080

## Consideraciones para Windows

Si el sistema operativo es Windows se debe tener instalado y en el PATH el compilador gcc, y se debe configurar la variable de entorno de Go `CGO_ENABLED` con el valor `1`. Para esto, debemos correr el siguiente comando en una terminal:

```bash
go env -w CGO_ENABLED=1
```
Luego verificamos las variables de entorno con el comando:

```bash
go env
```

## Uso

La aplicación está configurada para usar el puerto 8080, podemos acceder a los endpoints publicados a travez de `http://localhost:8080`

## Endpoints

Ejemplo:

```bash
curl --location 'localhost:8080/agents/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "LAUTARO",
    "faction": "COSMIC"
}'
```

Este endpoint es para registrar un agente. Hay varios endpoints, entre ellos para obtener agentes, obtener información de un agente, listar contratos y aceptar contrato.

## Contribuciones

Si deseas contribuir, por favor realize un fork de este repositorio y envía un Pull Request.
Cabe aclarar que este proyecto es universitario.

---

Gracias por utilizar nuestro software.
