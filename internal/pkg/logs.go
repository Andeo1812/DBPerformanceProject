package pkg

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CreateLog(ctx context.Context, errFull error) {
	errCause := errors.Cause(errFull)

	logger, ok := ctx.Value(LoggerKey).(*logrus.Entry)
	if !ok {
		logrus.Infof("CreateLog: errFull convert context -> *logrus.Logger on errFull [%s]", errCause)
		return
	}

	logLevel, err := GetLogLevelErr(errCause)
	if err != nil {
		logger.Error(errors.Wrap(errFull, "Undefined error"))
		return
	}

	switch logLevel {
	case errLogLevel:
		logger.Error(errFull)
	case debugLogLevel:
		logger.Debug(errFull)
	default:
		logger.Info(errCause)
	}
}

type ErrLogClassifier struct {
	table map[string]string
}

const (
	infoLogLevel  = "info"
	errLogLevel   = "error"
	debugLogLevel = "api"
)

func NewErrLogClassifier() ErrLogClassifier {
	res := make(map[string]string)

	// Common delivery
	res[ErrBadBodyRequest.Error()] = errLogLevel
	res[ErrJSONUnexpectedEnd.Error()] = errLogLevel
	res[ErrContentTypeUndefined.Error()] = errLogLevel
	res[ErrUnsupportedMediaType.Error()] = errLogLevel
	res[ErrEmptyBody.Error()] = errLogLevel
	res[ErrConvertQueryType.Error()] = errLogLevel
	res[ErrQueryRequiredEmpty.Error()] = errLogLevel
	res[ErrBadRequestParams.Error()] = errLogLevel
	res[ErrBadRequestParamsEmptyRequiredFields.Error()] = errLogLevel
	res[ErrBadRequestParams.Error()] = errLogLevel
	res[ErrGetEasyJSON.Error()] = errLogLevel

	// Common repository
	res[ErrNotFoundInDB.Error()] = errLogLevel
	res[ErrWorkDatabase.Error()] = errLogLevel
	res[ErrGetParamsConvert.Error()] = errLogLevel

	// Middleware
	res[ErrBigRequest.Error()] = errLogLevel
	res[ErrConvertLength.Error()] = errLogLevel

	// User
	res[ErrSuchUserExist.Error()] = errLogLevel
	res[ErrSuchUserNotFound.Error()] = errLogLevel
	res[ErrUpdateUserDataConflict.Error()] = errLogLevel

	// Thread
	res[ErrSuchThreadNotFound.Error()] = errLogLevel

	// Thread
	res[ErrNoSuchRuleSortPosts.Error()] = errLogLevel

	return ErrLogClassifier{
		table: res,
	}
}

func GetLogLevelErr(err error) (string, error) {
	level, exist := errLogCsf.table[err.Error()]
	if !exist {
		return "", errors.New("error not found")
	}

	return level, nil
}

var errLogCsf = NewErrLogClassifier()
