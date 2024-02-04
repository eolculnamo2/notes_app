package main

import (
	"context"
	"desktop_notes/core"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ListFiles() []string {
	dirFiles, err := fileops.ListSavedFiles()

	if err != nil {
		return make([]string, 0)
	}

	files := make([]string, len(dirFiles))
	for _, file := range dirFiles {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		files = append(files, name)
	}
	return files
}

func (a *App) SaveFile(name string, contents string, password string) string {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	err := fileops.SaveFile(name, contents, password)
	if err != nil {
		return "Failed to save file"
	}
	return "File Saved!"
}

func (a *App) ReadSaved(name string, password string) string {
	fileContents := fileops.ReadFile(name, password)
	return fileContents
}
