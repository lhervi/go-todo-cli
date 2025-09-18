package main

import (
	"os"
	"testing"
	"time"
)

func TestSaveAndLoadTasks(t *testing.T) {
	testFile := "test_tasks.json"

	defer os.Remove(testFile)

	tasksToSave := []Task{
		{ID: 1, Content: "Comprar pan", Completed: false, CreatedAt: time.Now()},
		{ID: 2, Content: "Aprender Go", Completed: true, CreatedAt: time.Now()},
	}

	if err := SaveTasks(testFile, tasksToSave); err != nil {
		t.Fatalf("Error al guardar las tareas: %v", err)
	}

	tasksLoaded, err := LoadTasks(testFile)
	if err != nil {
		t.Fatalf("Error al cargar las tareas: %v", err)
	}

	if len(tasksToSave) != len(tasksLoaded) {
		t.Fatalf("El número de tareas no coincide. Se esperaba %d, se obtuvo %d", len(tasksToSave), len(tasksLoaded))
	}

	for i := range tasksToSave {
		// Verificamos que los campos se hayan guardado y cargado correctamente.
		if tasksToSave[i].ID != tasksLoaded[i].ID {
			t.Errorf("ID de tarea incorrecto en el índice %d", i)
		}
		if tasksToSave[i].Content != tasksLoaded[i].Content {
			t.Errorf("Contenido de tarea incorrecto en el índice %d", i)
		}
		if tasksToSave[i].Completed != tasksLoaded[i].Completed {
			t.Errorf("Estado de tarea incorrecto en el índice %d", i)
		}

		// Aserción para el campo CreatedAt
		// Usamos un delta de tiempo para permitir pequeñas diferencias de precisión
		// entre la fecha guardada y la cargada.
		timeDelta := time.Since(tasksLoaded[i].CreatedAt)
		if timeDelta > time.Second || timeDelta < -time.Second {
			t.Errorf("La fecha de creación de la tarea en el índice %d es incorrecta", i)
		}
	}
}
