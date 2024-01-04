package exception

import "fmt"

type VideoAppError struct {
	httpStatus       int
	errorDescription string
}

func NewVideoAppError(httpStatus int, errDescription string) *VideoAppError {

	return &VideoAppError{
		httpStatus:       httpStatus,
		errorDescription: errDescription,
	}
}

func (videoAppError *VideoAppError) Error() string {

	return fmt.Sprintf("HttpStatus : [%d] msg : [%s]", videoAppError.httpStatus, videoAppError.errorDescription)
}

func (videoAppError *VideoAppError) GetHttpStatus() int {

	return videoAppError.httpStatus
}

func (videoAppError *VideoAppError) GetErrorDescription() string {

	return videoAppError.errorDescription
}
