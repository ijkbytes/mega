package api

const (
	CodeSuccess       = 200
	CodeInvalidParams = 400
	CodeError         = 500

	CodeNotExistTag     = 10001
	CodeNotExistArticle = 10002
)

func getErrMsg(code int) string {
	switch code {
	case CodeSuccess:
		return "success"
	case CodeInvalidParams:
		return "invalid params"
	case CodeError:
		return "internal error"
	case CodeNotExistTag:
		return "cannot find tag"
	case CodeNotExistArticle:
		return "cannot find article"
	default:
		return "unknown error"
	}
}
