package errors

import "net/http"

// http status code
const (
	_                            err = iota
	HttpStatusContinue           err = http.StatusContinue           // RFC 9110, 15.2.1
	HttpStatusSwitchingProtocols err = http.StatusSwitchingProtocols // RFC 9110, 15.2.2
	HttpStatusProcessing         err = http.StatusProcessing         // RFC 2518, 10.1
	HttpStatusEarlyHints         err = http.StatusEarlyHints         // RFC 8297
	HttpStatusOK                 err = http.StatusOK                 // RFC 9110, 15.3.1

	HttpStatusCreated              err = http.StatusCreated              // RFC 9110, 15.3.2
	HttpStatusAccepted             err = http.StatusAccepted             // RFC 9110, 15.3.3
	HttpStatusNonAuthoritativeInfo err = http.StatusNonAuthoritativeInfo // RFC 9110, 15.3.4
	HttpStatusNoContent            err = http.StatusNoContent            // RFC 9110, 15.3.5
	HttpStatusResetContent         err = http.StatusResetContent         // RFC 9110, 15.3.6
	HttpStatusPartialContent       err = http.StatusPartialContent       // RFC 9110, 15.3.7
	HttpStatusMultiStatus          err = http.StatusMultiStatus          // RFC 4918, 11.1
	HttpStatusAlreadyReported      err = http.StatusAlreadyReported      // RFC 5842, 7.1
	HttpStatusIMUsed               err = http.StatusIMUsed               // RFC 3229, 10.4.1

	HttpStatusMultipleChoices   err = http.StatusMultipleChoices   // RFC 9110, 15.4.1
	HttpStatusMovedPermanently  err = http.StatusMovedPermanently  // RFC 9110, 15.4.2
	HttpStatusFound             err = http.StatusFound             // RFC 9110, 15.4.3
	HttpStatusSeeOther          err = http.StatusSeeOther          // RFC 9110, 15.4.4
	HttpStatusNotModified       err = http.StatusNotModified       // RFC 9110, 15.4.5
	HttpStatusUseProxy          err = http.StatusUseProxy          // RFC 9110, 15.4.6
	_                                                              // RFC 9110, 15.4.7 (Unused)
	HttpStatusTemporaryRedirect err = http.StatusTemporaryRedirect // RFC 9110, 15.4.8
	HttpStatusPermanentRedirect err = http.StatusPermanentRedirect // RFC 9110, 15.4.9

	HttpStatusBadRequest                   err = http.StatusBadRequest                   // RFC 9110, 15.5.1
	HttpStatusUnauthorized                 err = http.StatusUnauthorized                 // RFC 9110, 15.5.2
	HttpStatusPaymentRequired              err = http.StatusPaymentRequired              // RFC 9110, 15.5.3
	HttpStatusForbidden                    err = http.StatusForbidden                    // RFC 9110, 15.5.4
	HttpStatusNotFound                     err = http.StatusNotFound                     // RFC 9110, 15.5.5
	HttpStatusMethodNotAllowed             err = http.StatusMethodNotAllowed             // RFC 9110, 15.5.6
	HttpStatusNotAcceptable                err = http.StatusNotAcceptable                // RFC 9110, 15.5.7
	HttpStatusProxyAuthRequired            err = http.StatusProxyAuthRequired            // RFC 9110, 15.5.8
	HttpStatusRequestTimeout               err = http.StatusRequestTimeout               // RFC 9110, 15.5.9
	HttpStatusConflict                     err = http.StatusConflict                     // RFC 9110, 15.5.10
	HttpStatusGone                         err = http.StatusGone                         // RFC 9110, 15.5.11
	HttpStatusLengthRequired               err = http.StatusLengthRequired               // RFC 9110, 15.5.12
	HttpStatusPreconditionFailed           err = http.StatusPreconditionFailed           // RFC 9110, 15.5.13
	HttpStatusRequestEntityTooLarge        err = http.StatusRequestEntityTooLarge        // RFC 9110, 15.5.14
	HttpStatusRequestURITooLong            err = http.StatusRequestURITooLong            // RFC 9110, 15.5.15
	HttpStatusUnsupportedMediaType         err = http.StatusUnsupportedMediaType         // RFC 9110, 15.5.16
	HttpStatusRequestedRangeNotSatisfiable err = http.StatusRequestedRangeNotSatisfiable // RFC 9110, 15.5.17
	HttpStatusExpectationFailed            err = http.StatusExpectationFailed            // RFC 9110, 15.5.18
	HttpStatusTeapot                       err = http.StatusTeapot                       // RFC 9110, 15.5.19 (Unused)
	HttpStatusMisdirectedRequest           err = http.StatusMisdirectedRequest           // RFC 9110, 15.5.20
	HttpStatusUnprocessableEntity          err = http.StatusUnprocessableEntity          // RFC 9110, 15.5.21
	HttpStatusLocked                       err = http.StatusLocked                       // RFC 4918, 11.3
	HttpStatusFailedDependency             err = http.StatusFailedDependency             // RFC 4918, 11.4
	HttpStatusTooEarly                     err = http.StatusTooEarly                     // RFC 8470, 5.2.
	HttpStatusUpgradeRequired              err = http.StatusUpgradeRequired              // RFC 9110, 15.5.22
	HttpStatusPreconditionRequired         err = http.StatusPreconditionRequired         // RFC 6585, 3
	HttpStatusTooManyRequests              err = http.StatusTooManyRequests              // RFC 6585, 4
	HttpStatusRequestHeaderFieldsTooLarge  err = http.StatusRequestHeaderFieldsTooLarge  // RFC 6585, 5
	HttpStatusUnavailableForLegalReasons   err = http.StatusUnavailableForLegalReasons   // RFC 7725, 3

	HttpStatusInternalServerError           err = http.StatusInternalServerError           // RFC 9110, 15.6.1
	HttpStatusNotImplemented                err = http.StatusNotImplemented                // RFC 9110, 15.6.2
	HttpStatusBadGateway                    err = http.StatusBadGateway                    // RFC 9110, 15.6.3
	HttpStatusServiceUnavailable            err = http.StatusServiceUnavailable            // RFC 9110, 15.6.4
	HttpStatusGatewayTimeout                err = http.StatusGatewayTimeout                // RFC 9110, 15.6.5
	HttpStatusHTTPVersionNotSupported       err = http.StatusHTTPVersionNotSupported       // RFC 9110, 15.6.6
	HttpStatusVariantAlsoNegotiates         err = http.StatusVariantAlsoNegotiates         // RFC 2295, 8.1
	HttpStatusInsufficientStorage           err = http.StatusInsufficientStorage           // RFC 4918, 11.5
	HttpStatusLoopDetected                  err = http.StatusLoopDetected                  // RFC 5842, 7.2
	HttpStatusNotExtended                   err = http.StatusNotExtended                   // RFC 2774, 7
	HttpStatusNetworkAuthenticationRequired err = http.StatusNetworkAuthenticationRequired // RFC 6585, 6
)

var httpErr = []string{}

func init() {
	httpErr[HttpStatusContinue] = ""
	httpErr[HttpStatusSwitchingProtocols] = ""
	httpErr[HttpStatusProcessing] = ""
	httpErr[HttpStatusEarlyHints] = ""
	httpErr[HttpStatusOK] = ""
	httpErr[HttpStatusCreated] = ""
	httpErr[HttpStatusAccepted] = ""
	httpErr[HttpStatusNonAuthoritativeInfo] = ""
	httpErr[HttpStatusNoContent] = ""
	httpErr[HttpStatusResetContent] = ""
	httpErr[HttpStatusPartialContent] = ""
	httpErr[HttpStatusMultiStatus] = ""
	httpErr[HttpStatusAlreadyReported] = ""
	httpErr[HttpStatusIMUsed] = ""
	httpErr[HttpStatusMultipleChoices] = ""
	httpErr[HttpStatusMovedPermanently] = ""
	httpErr[HttpStatusFound] = ""
	httpErr[HttpStatusSeeOther] = ""
	httpErr[HttpStatusNotModified] = ""
	httpErr[HttpStatusUseProxy] = ""
	httpErr[HttpStatusTemporaryRedirect] = ""
	httpErr[HttpStatusPermanentRedirect] = ""
	httpErr[HttpStatusBadRequest] = ""
	httpErr[HttpStatusUnauthorized] = ""
	httpErr[HttpStatusPaymentRequired] = ""
	httpErr[HttpStatusForbidden] = ""
	httpErr[HttpStatusNotFound] = ""
	httpErr[HttpStatusMethodNotAllowed] = ""
	httpErr[HttpStatusNotAcceptable] = ""
	httpErr[HttpStatusProxyAuthRequired] = ""
	httpErr[HttpStatusRequestTimeout] = ""
	httpErr[HttpStatusConflict] = ""
	httpErr[HttpStatusGone] = ""
	httpErr[HttpStatusLengthRequired] = ""
	httpErr[HttpStatusPreconditionFailed] = ""
	httpErr[HttpStatusRequestEntityTooLarge] = ""
	httpErr[HttpStatusRequestURITooLong] = ""
	httpErr[HttpStatusUnsupportedMediaType] = ""
	httpErr[HttpStatusRequestedRangeNotSatisfiable] = ""
	httpErr[HttpStatusExpectationFailed] = ""
	httpErr[HttpStatusTeapot] = ""
	httpErr[HttpStatusMisdirectedRequest] = ""
	httpErr[HttpStatusUnprocessableEntity] = ""
	httpErr[HttpStatusLocked] = ""
	httpErr[HttpStatusFailedDependency] = ""
	httpErr[HttpStatusTooEarly] = ""
	httpErr[HttpStatusUpgradeRequired] = ""
	httpErr[HttpStatusPreconditionRequired] = ""
	httpErr[HttpStatusTooManyRequests] = ""
	httpErr[HttpStatusRequestHeaderFieldsTooLarge] = ""
	httpErr[HttpStatusUnavailableForLegalReasons] = ""
	httpErr[HttpStatusInternalServerError] = ""
	httpErr[HttpStatusNotImplemented] = ""
	httpErr[HttpStatusBadGateway] = ""
	httpErr[HttpStatusServiceUnavailable] = ""
	httpErr[HttpStatusGatewayTimeout] = ""
	httpErr[HttpStatusHTTPVersionNotSupported] = ""
	httpErr[HttpStatusVariantAlsoNegotiates] = ""
	httpErr[HttpStatusInsufficientStorage] = ""
	httpErr[HttpStatusLoopDetected] = ""
	httpErr[HttpStatusNotExtended] = ""
	httpErr[HttpStatusNetworkAuthenticationRequired] = ""
}
