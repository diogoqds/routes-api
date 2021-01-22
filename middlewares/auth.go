package middlewares

import (
	"fmt"
	"github.com/diogoqds/routes-challenge-api/repositories"
	"net/http"
	"strconv"

	"strings"

	"github.com/diogoqds/routes-challenge-api/infra"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{ "message": "Malformed Token" }`))
			return
		} else {
			body, err := infra.Jwt.Decoder.Decode(authHeader[1])
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{ "message": "Unauthorized"`))
				return
			}

			id, err := strconv.ParseInt(fmt.Sprintf("%.0f", body["id"]), 10, 64)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			_, err = repositories.AdminRepo.FinderById.FindById(id)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			next(w, r)
		}
	}
}
