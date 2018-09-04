package SimpleResourceApi

import (
	"fmt"
	"log"
	"net/http"

	"SimpleResourceApi/config"

	"github.com/Vilsol/GoLib"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"SimpleResourceApi/database"
	"SimpleResourceApi/schema"
	"SimpleResourceApi/nodes"
)

func Serve() {
	config.InitializeConfig()
	schema.InitializeSchemas()
	database.InitializeDatabase()

	router := mux.NewRouter()
	router.NotFoundHandler = GoLib.LoggerHandler(GoLib.NotFoundHandler())

	v1 := GoLib.RouteHandler(router, "/v1")
	nodes.RegisterResourceRoutes(v1)

	CORSHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
	)

	var finalRouter http.Handler = router
	finalRouter = GoLib.LoggerHandler(finalRouter)
	finalRouter = handlers.CompressHandler(finalRouter)
	finalRouter = handlers.ProxyHeaders(finalRouter)
	finalRouter = CORSHandler(finalRouter)

	fmt.Printf("Listening on port %d\n", config.Get().Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Port), finalRouter))
}
