# Go CLI Todo App

Una aplicación de línea de comandos simple y eficiente para gestionar tareas,
construida en Go. Este proyecto fue desarrollado para demostrar las habilidades
básicas en el lenguaje, incluyendo el manejo de argumentos, la serialización de
datos a JSON y el testing unitario.

### Características

- **Agregar Tareas:** Añade nuevas tareas con un contenido específico.
- **Listar Tareas:** Muestra todas las tareas existentes con su estado de
  completado.
- **Completar Tareas:** Marca una tarea como completada usando su ID.
- **Eliminar Tareas:** Borra una tarea de la lista usando su ID.
- **Persistencia de Datos:** Las tareas se guardan automáticamente en un archivo
  `tasks.json`.

### Uso

Para ejecutar la aplicación, navega al directorio del proyecto en tu terminal.

#### 📝 Agregar una tarea

```bash
go run . -add -task "Mi primera tarea en Go"
```
