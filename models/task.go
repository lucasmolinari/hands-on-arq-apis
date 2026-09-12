package models

type Task struct {
	ID          int    `json:"id" example:"1"`
	Title       string `json:"title" example:"Estudar Swaggo"`
	Description string `json:"description" example:"Preparar roteiro da apresentacao"`
	Done        bool   `json:"done" example:"false"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"task not found"`
}
