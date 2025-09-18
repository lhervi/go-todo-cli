package main

import (
	"testing"
	"time"
)

// TestTaskCreation verifica que una tarea se crea correctamente.
func TestTaskCreation(t *testing.T) {
	content := "Aprender Go"
	task := Task{
		ID:        1,
		Content:   content,
		Completed: false,
		CreatedAt: time.Now(),
	}

	// Comprobamos que el contenido y el estado inicial sean correctos.
	if task.Content != content {
		t.Errorf("Content incorrecto, se esperaba %s, se obtuvo %s", content, task.Content)
	}

	if task.Completed {
		t.Error("La tarea no debería estar completada al ser creada")
	}

	if task.ID != 1 {
		t.Errorf("ID incorrecto, se esperaba 1, se obtuvo %d", task.ID)
	}

	// Verificamos que CreatedAt no sea un valor "cero" de tiempo, lo que indica que se inicializó correctamente.
	if task.CreatedAt.IsZero() {
		t.Error("CreatedAt no fue inicializado correctamente")
	}

}
