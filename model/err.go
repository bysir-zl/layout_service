package model

var (
	ErrDefault   = NewErrCode(400, "")
	ErrNotAuth   = NewErrCode(401, "")
	ErrForbidden = NewErrCode(403, "")
	ErrNotFind   = NewErrCode(404, "")
	ErrBadParams = NewErrCode(407, "")
)

type ErrCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e ErrCode) Append(ext string) ErrCode {
	if e.Msg != "" {
		ext = " " + ext
	}
	e.Msg += ext
	return e
}

func (e ErrCode) Coder() int {
	return e.Code
}

func (e ErrCode) Error() string {
	return e.Msg
}

func (e ErrCode) Is(err error) bool {
	switch x := err.(type) {
	case ErrCode:
		return x.Code == e.Code
	}
	return false
}

func NewErrCode(code int, msg string) ErrCode {
	return ErrCode{code, msg}
}
