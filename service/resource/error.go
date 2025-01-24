package resource

import "errors"

var (
	ErrTenantIDNotAllow          = errors.New("not allow action in your tenant")
	ErrMutationOtherUserNotAllow = errors.New("not allow mutation other user")
)
