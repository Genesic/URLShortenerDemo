package errors

type ServiceError struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (e ServiceError) Error() string {
	return e.Msg
}

var DatabaseErrorMsg = "internal server error"

var (
	InsertDataBaseFailedError = &ServiceError{500, 50001, DatabaseErrorMsg}
	FetchDatabaseFailedError  = &ServiceError{500, 50002, DatabaseErrorMsg}
	UrlNotFoundError          = &ServiceError{404, 40401, "url not found"}
)
