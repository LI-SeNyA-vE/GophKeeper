package middleware

import (
	"GophKeeper/pkg/jwttoken"
	"net/http"
	"strings"
)

func (m *Middleware) JwtCheck(h http.Handler) http.Handler {
	jwtCheck := func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if !strings.HasPrefix(bearerToken, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		// Вытаскиваем JWT токен
		token := strings.TrimPrefix(bearerToken, "Bearer ")
		// Проверяем его валидность
		_, err := jwttoken.ValidateToken(token, "m.secret") // TODO исправить получение секретного ключа
		if err != nil {
			//TODO поправить правильны вывод ошибки
			return
		}

		// Передаём управление дальше
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(jwtCheck)
}
