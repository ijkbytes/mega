package api

const (
	CodeSuccess       = 200
	CodeInvalidParams = 400
	CodeError         = 500

	CodeNotExistTag     = 10001
	CodeNotExistArticle = 10002

	CodeUserError      = 20001
	CodeUserExist      = 20002
	CodeUserLoginError = 20003
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
	case CodeUserError:
		return "user error"
	case CodeUserExist:
		return "user has been exist"
	case CodeUserLoginError:
		return "password or user name error"
	default:
		return "unknown error"
	}
}
