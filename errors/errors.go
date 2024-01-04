package errors

import "github.com/rotisserie/eris"

var (
	ErrLock        = eris.New("failed to lock event")
	ErrUnlock      = eris.New("failed to unlock event")
	ErrEventLocked = eris.New("failed to lock, event still running")
	ErrClose       = eris.New("failed to close connection")
)

func Is(err1, err2 error) bool { return eris.Is(err1, err2) }

func IsAny(err error, errAny ...error) bool {
	for i := range errAny {
		if Is(err, errAny[i]) {
			return true
		}
	}

	return false
}
