package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"swaggo-demo/models"

	"github.com/gin-gonic/gin"
)

var (
	tasks = map[int]models.Task{
		1: {ID: 1, Title: "Estudar Swaggo", Description: "Ler a documentacao oficial", Done: false},
		2: {ID: 2, Title: "Montar projeto exemplo", Description: "Criar API com anotacoes", Done: true},
	}
	nextID = 3
	mu     sync.Mutex
)

// ListTasks godoc
// @Summary      Lista todas as tarefas
// @Description  Retorna todas as tarefas cadastradas
// @Tags         tasks
// @Produce      json
// @Success      200  {array}   models.Task
// @Router       /tasks [get]
func ListTasks(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	result := make([]models.Task, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, t)
	}
	c.JSON(http.StatusOK, result)
}

// GetTask godoc
// @Summary      Busca uma tarefa pelo ID
// @Description  Retorna uma tarefa especifica
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID da tarefa"
// @Success      200  {object}  models.Task
// @Failure      404  {object}  models.ErrorResponse
// @Router       /tasks/{id} [get]
func GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	task, ok := tasks[id]
	if !ok {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask godoc
// @Summary      Cria uma nova tarefa
// @Description  Adiciona uma tarefa a lista
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      models.Task  true  "Dados da tarefa"
// @Success      201   {object}  models.Task
// @Failure      400   {object}  models.ErrorResponse
// @Router       /tasks [post]
func CreateTask(c *gin.Context) {
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	input.ID = nextID
	nextID++
	tasks[input.ID] = input

	c.JSON(http.StatusCreated, input)
}

// UpdateTask godoc
// @Summary      Atualiza uma tarefa existente
// @Description  Atualiza os campos de uma tarefa pelo ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "ID da tarefa"
// @Param        task  body      models.Task  true  "Dados atualizados"
// @Success      200   {object}  models.Task
// @Failure      404   {object}  models.ErrorResponse
// @Router       /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "invalid id"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := tasks[id]; !ok {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "task not found"})
		return
	}

	input.ID = id
	tasks[id] = input
	c.JSON(http.StatusOK, input)
}

// DeleteTask godoc
// @Summary      Remove uma tarefa
// @Description  Remove uma tarefa pelo ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "ID da tarefa"
// @Success      204
// @Failure      404  {object}  models.ErrorResponse
// @Router       /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := tasks[id]; !ok {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "task not found"})
		return
	}
	delete(tasks, id)
	c.Status(http.StatusNoContent)
}
