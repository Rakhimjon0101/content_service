package errs

var (
	ErrEntityRelationNotLoaded = New("relation is not loaded")
	ErrWrongInput              = New("wrong input")
	ErrUnauthorized            = New("unauthorized")
	ErrNotFound                = New("entity not found")
	ErrConnection              = New("connection error")
	ErrInternal                = New("internal error")
)
