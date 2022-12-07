package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-axesthump-reefiw-tic-tac-toe/internal/game"
	"go-axesthump-reefiw-tic-tac-toe/internal/generator"
	myMiddleware "go-axesthump-reefiw-tic-tac-toe/internal/middleware"
	"go-axesthump-reefiw-tic-tac-toe/internal/storage"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type GameModel struct {
	Map        []string
	WinMessage string
}

var gameTemplate = template.Must(
	template.New("data").Parse(
		`
<h1>Game:</h1>
<h2>{{ .WinMessage}}</h2>
{{range .Map}}
     {{.}}
	<br>
{{end}}
</p>
<form action="/move" method="get" class="form-example">
    <div class="form-example">
        <label for="x">Enter x: </label>
        <input type="text" name="x" id="x" required>
    </div>
    <div class="form-example">
        <label for="y">Enter y: </label>
        <input type="text" name="y" id="y" required>
    </div>
    <div class="form-example">
        <input type="submit" value="Move!">
    </div>
</form>
`,
	),
)

type AppHandler struct {
	gameField  *game.Map
	Router     chi.Router
	storage    storage.Storage
	WinMessage string
}

func NewAppHandler() (*AppHandler, error) {
	gameMap, err := game.InitMap(3, 3)
	if err != nil {
		return nil, err
	}
	handler := &AppHandler{
		gameField: gameMap,
		storage:   storage.NewInMemoryStorage(),
	}
	createRouter(handler)
	return handler, nil
}

func createRouter(h *AppHandler) {
	r := chi.NewRouter()

	r.Use(myMiddleware.NewAuthService(generator.NewIDGenerator(0)).Auth)
	r.Use(middleware.SetHeader("Content-Type", "text/html"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", h.gamePage)
	r.Get("/move", h.playerMove)
	h.Router = r
}

func (h *AppHandler) gamePage(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(myMiddleware.UserIDKey).(uint32)
	user, err := h.storage.GetUser(int(id))
	if err != nil {
		return // todo
	}
	if user.ID == -1 {
		h.storage.AddUser(int(id))
	}
	gameModel := GameModel{
		Map: h.gameField.GetMapForResponse(),
	}
	gameTemplate.Execute(w, gameModel)
}

func (h *AppHandler) playerMove(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(myMiddleware.UserIDKey).(uint32)
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")
	user, err := h.storage.GetUser(int(id))
	if err != nil {
		log.Println(err.Error())
		return
	}
	xInt, err := strconv.Atoi(x)
	if err != nil {
		return //todo
	}
	yInt, err := strconv.Atoi(y)
	if err != nil {
		return //todo
	}
	if user.PlayerType == storage.OPlayer {
		if h.gameField.LastPlayerSymbol != 'o' {
			win, _ := h.gameField.Move(yInt, xInt, 'o') //todo
			if win {
				h.WinMessage = "Player o win!"
			}
		}
	} else {
		if h.gameField.LastPlayerSymbol != 'x' {
			win, _ := h.gameField.Move(yInt, xInt, 'x') //todo
			if win {
				h.WinMessage = "Player x win!"
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
