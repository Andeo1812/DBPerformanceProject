package pkg

import (
	"net/http"

	"github.com/pkg/errors"
)

var (
	// Common delivery
	ErrBadBodyRequest                      = errors.New("bad body request")
	ErrJSONUnexpectedEnd                   = errors.New("unexpected end of JSON input")
	ErrContentTypeUndefined                = errors.New("content-type undefined")
	ErrUnsupportedMediaType                = errors.New("unsupported media type")
	ErrEmptyBody                           = errors.New("empty body")
	ErrConvertQueryType                    = errors.New("bad input query")
	ErrQueryRequiredEmpty                  = errors.New("miss query params")
	ErrBadRequestParams                    = errors.New("bad query params")
	ErrBadRequestParamsEmptyRequiredFields = errors.New("bad params, empty required field")
	ErrUpdateWebSocketProtocol             = errors.New("err update ws protocol")
	ErrGetEasyJSON                         = errors.New("err get easyjson")

	// Common repository
	ErrNotFoundInDB             = errors.New("not found")
	ErrWorkDatabase             = errors.New("error sql")
	ErrGetParamsConvert         = errors.New("err get sql params")
	ErrUnsupportedSortParameter = errors.New("unsupported sort parameter")

	// Collection service
	ErrNotFindSuchTarget     = errors.New("not found such target")
	ErrCollectionIsNotPublic = errors.New("this collection is not public")

	// Auth delivery
	ErrNoCookie = errors.New("no such cookie")

	// Auth repository
	ErrUserExist     = errors.New("such user exists")
	ErrUserNotExist  = errors.New("no such user")
	ErrCreateSession = errors.New("can't create new session")

	// Auth service
	ErrInvalidNickname   = errors.New("invalid nickname")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrIncorrectPassword = errors.New("incorrect password")

	// Image delivery
	ErrBigImage     = errors.New("big image")
	ErrBadImageType = errors.New("bad image type")

	// Image repository
	ErrImage = errors.New("service picture not work")

	// User delivery
	ErrGetUserRequest     = errors.New("fatal getting user")
	ErrWrongValidPassword = errors.New("bad pass")

	// User service
	ErrBadUserCollectionID      = errors.New("this collection doesn't belong to current user")
	ErrFilmExistInCollection    = errors.New("such film exist in collection")
	ErrFilmNotExistInCollection = errors.New("such film not found in collection")
	ErrFilmRatingNotExist       = errors.New("film rating not exist")

	// Middleware
	ErrBigRequest    = errors.New("big request")
	ErrConvertLength = errors.New("getting content-length failed")

	// Security
	ErrCsrfTokenCreate        = errors.New("csrf token create error")
	ErrCsrfTokenCheck         = errors.New("csrf token check error")
	ErrCsrfTokenCheckInternal = errors.New("csrf token check internal error")
	ErrCsrfTokenExpired       = errors.New("csrf token expired")
	ErrCsrfTokenInvalid       = errors.New("invalid csrf token")

	// Not Found
	ErrGenreNotFound       = errors.New("genre not found")
	ErrTagNotFound         = errors.New("tag not found")
	ErrFilmNotFound        = errors.New("film not found")
	ErrPersonNotFound      = errors.New("person not found")
	ErrImageNotFound       = errors.New("image not found")
	ErrSessionNotFound     = errors.New("session not found")
	ErrUserNotFound        = errors.New("user not found")
	ErrFilmsNotFound       = errors.New("films not found")
	ErrCollectionsNotFound = errors.New("collections not found")
	ErrCollectionNotFound  = errors.New("such collection doesn't exist")
)

type ErrHTTPClassifier struct {
	table map[string]int
}

func NewErrHTTPClassifier() ErrHTTPClassifier {
	res := make(map[string]int)

	// Common delivery
	res[ErrBadBodyRequest.Error()] = http.StatusBadRequest
	res[ErrJSONUnexpectedEnd.Error()] = http.StatusBadRequest
	res[ErrContentTypeUndefined.Error()] = http.StatusBadRequest
	res[ErrUnsupportedMediaType.Error()] = http.StatusUnsupportedMediaType
	res[ErrEmptyBody.Error()] = http.StatusBadRequest
	res[ErrConvertQueryType.Error()] = http.StatusBadRequest
	res[ErrQueryRequiredEmpty.Error()] = http.StatusBadRequest
	res[ErrBadRequestParams.Error()] = http.StatusBadRequest
	res[ErrBadRequestParamsEmptyRequiredFields.Error()] = http.StatusBadRequest
	res[ErrBadRequestParams.Error()] = http.StatusBadRequest
	res[ErrUpdateWebSocketProtocol.Error()] = http.StatusBadRequest
	res[ErrGetEasyJSON.Error()] = http.StatusInternalServerError

	// Common repository
	res[ErrNotFoundInDB.Error()] = http.StatusNotFound
	res[ErrWorkDatabase.Error()] = http.StatusInternalServerError
	res[ErrGetParamsConvert.Error()] = http.StatusInternalServerError
	res[ErrUnsupportedSortParameter.Error()] = http.StatusBadRequest

	// Collection service
	res[ErrNotFindSuchTarget.Error()] = http.StatusNotFound
	res[ErrCollectionIsNotPublic.Error()] = http.StatusForbidden

	// Auth delivery
	res[ErrNoCookie.Error()] = http.StatusNotFound

	// Auth repository
	res[ErrUserExist.Error()] = http.StatusBadRequest
	res[ErrUserNotExist.Error()] = http.StatusNotFound
	res[ErrCreateSession.Error()] = http.StatusInternalServerError

	// Auth service
	res[ErrInvalidNickname.Error()] = http.StatusBadRequest
	res[ErrInvalidEmail.Error()] = http.StatusBadRequest
	res[ErrInvalidPassword.Error()] = http.StatusBadRequest
	res[ErrIncorrectPassword.Error()] = http.StatusForbidden

	// Image delivery
	res[ErrBigImage.Error()] = http.StatusBadRequest
	res[ErrBadImageType.Error()] = http.StatusBadRequest

	// Image repository
	res[ErrImage.Error()] = http.StatusInternalServerError

	// User delivery
	res[ErrGetUserRequest.Error()] = http.StatusInternalServerError
	res[ErrWrongValidPassword.Error()] = http.StatusForbidden

	// User service
	res[ErrBadUserCollectionID.Error()] = http.StatusForbidden
	res[ErrFilmExistInCollection.Error()] = http.StatusBadRequest
	res[ErrFilmNotExistInCollection.Error()] = http.StatusNotFound
	res[ErrFilmRatingNotExist.Error()] = http.StatusNotFound

	// Middleware
	res[ErrBigRequest.Error()] = http.StatusBadRequest
	res[ErrConvertLength.Error()] = http.StatusBadRequest

	// Security
	res[ErrCsrfTokenCreate.Error()] = http.StatusInternalServerError
	res[ErrCsrfTokenCheck.Error()] = http.StatusForbidden
	res[ErrCsrfTokenCheckInternal.Error()] = http.StatusInternalServerError
	res[ErrCsrfTokenExpired.Error()] = http.StatusForbidden
	res[ErrCsrfTokenInvalid.Error()] = http.StatusForbidden

	// Not found
	res[ErrGenreNotFound.Error()] = http.StatusNotFound
	res[ErrTagNotFound.Error()] = http.StatusNotFound
	res[ErrPersonNotFound.Error()] = http.StatusNotFound
	res[ErrFilmNotFound.Error()] = http.StatusNotFound
	res[ErrSessionNotFound.Error()] = http.StatusNotFound
	res[ErrUserNotFound.Error()] = http.StatusNotFound
	res[ErrImageNotFound.Error()] = http.StatusNotFound
	res[ErrFilmsNotFound.Error()] = http.StatusNotFound
	res[ErrCollectionsNotFound.Error()] = http.StatusNotFound
	res[ErrCollectionNotFound.Error()] = http.StatusNotFound

	return ErrHTTPClassifier{
		table: res,
	}
}

var errHTTPCsf = NewErrHTTPClassifier()

func GetErrorCodeHTTP(err error) (int, bool) {
	code, exist := errHTTPCsf.table[err.Error()]
	if !exist {
		return http.StatusInternalServerError, exist
	}

	return code, exist
}
