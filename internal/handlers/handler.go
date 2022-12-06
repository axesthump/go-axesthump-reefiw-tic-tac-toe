package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-axesthump-reefiw-tic-tac-toe/internal/generator"
	myMiddleware "go-axesthump-reefiw-tic-tac-toe/internal/middleware"
	"go-axesthump-reefiw-tic-tac-toe/internal/storage"
	"html/template"
	"log"
	"net/http"
)

var gameTemplate = template.Must(
	template.New("data").Parse(
		`
<h1>Game:</h1>
<h2>Player id: {{.}}</h2>
<p>
_|_|_
<br>
_|_|_
<br>
_|_|_
<br>
</p>
`,
	),
)

type AppHandler struct {
	// todo GameField
	Router  chi.Router
	storage storage.Storage
}

func NewAppHandler() *AppHandler {
	// todo init GameField
	handler := &AppHandler{
		storage: storage.NewInMemoryStorage(),
	}
	createRouter(handler)
	return handler
}

func createRouter(h *AppHandler) {
	r := chi.NewRouter()

	r.Use(myMiddleware.NewAuthService(generator.NewIDGenerator(0)).Auth)
	r.Use(middleware.SetHeader("Content-Type", "text/html"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", h.gamePage)
	r.Post("/", h.playerMove)
	h.Router = r
}

func (h *AppHandler) gamePage(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(myMiddleware.UserIDKey)

	gameTemplate.Execute(w, id)
}

func (h *AppHandler) playerMove(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(myMiddleware.UserIDKey).(uint32)
	_, err := h.storage.GetUser(int(id))
	if err != nil {
		log.Println(err.Error())
	}
	gameTemplate.Execute(w, id)
}
