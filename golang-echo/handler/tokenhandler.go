package handler

import (
	"golang-echo/service"
)

type TokenHandler interface {
	GetToken() string
}

type tokenHandler struct {
	jwtService service.JWTService
}

func TokenHandle(jWtService service.JWTService) TokenHandler {
	return &tokenHandler{
		jwtService: jWtService,
	}
}

func (th *tokenHandler) GetToken() string {
	return th.jwtService.GenerateToken()
}
