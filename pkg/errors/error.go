package errors

type ServiceError struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (e ServiceError) Error() string {
	return e.Msg
}

var InternalErrorMsg = "internal server error"

var (
	InsertDataBaseFailedError = &ServiceError{500, 50001, InternalErrorMsg}
	FetchDatabaseFailedError  = &ServiceError{500, 50002, InternalErrorMsg}
	EncodeHashFailedError     = &ServiceError{500, 50003, InternalErrorMsg}

	ValidateRequestFailedError = &ServiceError{400, 40001, "validate request failed"}
	GetDomainFailedError       = &ServiceError{400, 40002, "get domain failed"}
	UrlNotFoundError           = &ServiceError{404, 40401, "url not found"}
)
