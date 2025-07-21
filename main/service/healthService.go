package service

import (
	"fmt"
	"github.com/Gircen/go-library-api/main/api/rest/response"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/healthCheck" {

		if req.Method == http.MethodGet {
			_, err := fmt.Fprint(w, response.OKNoContent())
			if err != nil {
				return
			}
		}
	}
}
