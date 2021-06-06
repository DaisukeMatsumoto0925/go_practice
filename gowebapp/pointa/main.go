package main

import "gorm.io/gorm"

type DB interface {
	Fetch(id int) (err error)
}

type App struct {
	DB DB
}

// Models DB models
type Models struct {
	DB *gorm.DB
}

// NewModels is constructor
func NewModels() (*Models, error) {
	db, err := SQLConnect()
	if err != nil {
		return &Models{}, err
	}
	return &Models{DB: db}, nil
}

func NewApp(models *Models) *App {
	return &App{
		DB: models,
	}
}
