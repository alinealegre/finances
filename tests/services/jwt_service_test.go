package controllers

import (
	"finances/services"
	"testing"
)

func TestJWTService_GenerateToken(t *testing.T) { // Inicialização do serviço JWTfunc TestJWTService_GenerateToken(t *testing.T) {
	jwtService := services.NewJWTService()

	// Gerando um token
	cpf := "12345678900"
	userType := true
	userEmail := "admin3@finances.com"
	token, err := jwtService.GenerateToken(cpf, userType, userEmail)

	// Verificando se não ocorreu erro durante a geração do token
	if err != nil {
		t.Errorf("Erro ao gerar token: %v", err)
		return // Termina o teste se houver erro na geração do token
	}

	// Verificando se o token não está vazio
	if token == "" {
		t.Error("Token vazio")
		return // Termina o teste se o token estiver vazio
	}
}
