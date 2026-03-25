package constant

type UserStatus string

const (
	UserStatusActive   UserStatus = "Active"
	UserStatusInActive UserStatus = "Inactive"
	UserStatusBlocked  UserStatus = "Blocked"
	UserStatusDeleted  UserStatus = "Deleted"
)

func (s UserStatus) IsValid() bool {
	switch s {
	case UserStatusActive, UserStatusInActive, UserStatusBlocked, UserStatusDeleted:
		return true
	}
	return false
}
