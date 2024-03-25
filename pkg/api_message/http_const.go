package messages

import "net/http"

func HttpStatusPt(code int) string {
	switch code {
	case http.StatusOK:
		return "OK"
	case http.StatusCreated:
		return "Criado"
	case http.StatusAccepted:
		return "Aceito"
	case http.StatusNoContent:
		return "Sem conteúdo"
	case http.StatusMovedPermanently:
		return "Movido permanentemente"
	case http.StatusFound:
		return "Encontrado"
	case http.StatusSeeOther:
		return "Veja outro"
	case http.StatusNotModified:
		return "Não modificado"
	case http.StatusBadRequest:
		return "Requisição inválida"
	case http.StatusUnauthorized:
		return "Não autorizado"
	case http.StatusForbidden:
		return "Proibido"
	case http.StatusNotFound:
		return "Não encontrado"
	case http.StatusMethodNotAllowed:
		return "Método não permitido"
	case http.StatusInternalServerError:
		return "Erro interno do servidor"
	case http.StatusNotImplemented:
		return "Não implementado"
	case http.StatusServiceUnavailable:
		return "Serviço indisponível"
	default:
		return "Código desconhecido"
	}
}
