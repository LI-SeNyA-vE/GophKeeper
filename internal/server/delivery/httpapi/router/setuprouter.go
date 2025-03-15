package router

import (
	"github.com/go-chi/chi/v5"
)

func (rout *Router) SetupRouter() {
	rout.Mux = chi.NewRouter()

	// Группа без проверки JWT
	rout.Mux.Group(func(public chi.Router) {
		public.Post("/registration", rout.handler.PostRegistrationUser)   // Регистрация нового пользователя
		public.Post("/authorization", rout.handler.PostAuthorizationUser) // Авторизация пользователя
		public.Post("/logout", rout.handler.PostLogoutUser)               // Выход пользователя (удаляет refresh токен)
	})

	// все остальные запросы
	rout.Mux.Group(func(protected chi.Router) {
		protected.Use(rout.middleware.JwtCheck)
		protected.Post("/addLoginAndPassword", rout.handler.PostAddLoginAndPassword)
	})
}
