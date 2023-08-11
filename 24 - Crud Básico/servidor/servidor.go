package servidor

import (
	"crud/banco"
	"encoding/json"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		http.Error(w, "Falha ao ler o corpo da requisição!", http.StatusBadRequest)
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		http.Error(w, "Erro ao converter o usuário para struct", http.StatusBadRequest)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados!", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if erro != nil {
		http.Error(w, "Erro ao criar o statement!", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		http.Error(w, "Erro ao executar o statement!", http.StatusInternalServerError)
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		http.Error(w, "Erro ao obter o id inserido!", http.StatusInternalServerError)
		return
	}

	response := struct {
		ID     int64  `json:"id"`
		Status string `json:"status"`
	}{
		ID:     idInserido,
		Status: "Usuário inserido com sucesso!",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Erro ao criar resposta JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
