package servidor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ...

func TestCriarUsuario(t *testing.T) {
	body := `{"nome": "Malteus", "email": "malteus@example.com"}`
	req, err := http.NewRequest("POST", "/usuarios", bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CriarUsuario(rr, r)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Código de status incorreto: esperado %d, obtido %d", http.StatusCreated, rr.Code)
	}

	responseBody := rr.Body.String()
	fmt.Println("Corpo da resposta:", responseBody)

	var response struct {
		ID int64 `json:"id"`
	}
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.ID == 0 {
		t.Error("ID inserido inválido")
	}

	assert.Contains(t, responseBody, "Usuário inserido com sucesso!", "Resposta incorreta")
}
