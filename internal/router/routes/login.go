package routes

import (
	"api/internal/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	AuthRequired: false,
}
