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

}
