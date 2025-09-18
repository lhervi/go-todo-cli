package main

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestSaveAndLoadTasks(t *testing.T) {

	testFile := "test_tasks.json"

	// Creamos un archivo de prueba temporal para las tareas.
	testFile := "test_task.json"
	tasksFile = testFile

	// Aseguramos que el archivo de prueba se elimine al finalizar la prueba.
	defer os.Remove(testFile)

	tasksToSave := []Task{
		{ID: 1, Content: "Comprar pan", Completed: false, CreatedAt: time.Now()},
		{ID: 2, Content: "Aprender Go", Completed: true, CreatedAt: time.Now()},
	}

	// Cargamos las tareas desde el archivo.
	tasksLoaded, err := LoadTasks()
	if err != nil {
		t.Fatalf("Error al cargar las tareas: %v", err)
	}

	if !reflect.DeepEqual(tasksToSave, tasksLoaded) {
		t.Errorf("Las tareas guardas y cargadas no son iguales")
	}
}
