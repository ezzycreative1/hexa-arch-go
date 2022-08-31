package main

import (
	"flag"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	handlerTodo "github.com/ezzycreative1/hexa-arch/app/handlers/todo"
	"github.com/ezzycreative1/hexa-arch/helpers"
	"github.com/ezzycreative1/hexa-arch/helpers/logging"
	"github.com/ezzycreative1/hexa-arch/internal/core/ports"
	usecases "github.com/ezzycreative1/hexa-arch/internal/core/usecases"
	repoTodo "github.com/ezzycreative1/hexa-arch/internal/repositories/todo"
	"go.uber.org/zap"
)

var (
	repo    string
	binding string

	log *zap.SugaredLogger = logging.NewLogger()
)

func init() {
	flag.StringVar(&repo, "repo", "mysql", "Mongo or MySql")
	flag.StringVar(&binding, "httpbind", ":8080", "address/port to bind listen socket")

	flag.Parse()
}

func main() {
	var todoRepo ports.TodoRepository
	if repo == "mysql" {
		todoRepo = startMysqlRepo()
	} else {
		todoRepo = startMongoRepo()
	}

	todoUseCase := usecases.NewToDoUseCase(todoRepo)

	ws := new(restful.WebService)
	ws = ws.Path("/api")
	handlerTodo.NewTodoHandler(todoUseCase, ws)
	restful.Add(ws)

	log.Info("Listening...")

	log.Panic(http.ListenAndServe(binding, nil))
}

func startMongoRepo() ports.TodoRepository {
	return repoTodo.NewTodoMongoRepo(helpers.StartMongoDb())
}

func startMysqlRepo() ports.TodoRepository {
	return repoTodo.NewTodoMysqlRepo(helpers.StartMysqlDb())
}
