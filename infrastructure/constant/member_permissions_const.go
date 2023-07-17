package constant

const (
	OWNER  = "owner"
	URGE   = "urge"
	MEMBER = "member"
)

const (
	Read   = "read"
	Create = "create"
	Update = "update"
	Delete = "delete"
	Invite = "invite"
)

func Owner() []string {
	return []string{Read, Create, Update, Delete, Invite}
}

func Urge() []string {
	return []string{Read, Create, Delete, Invite}
}

func Member() []string {
	return []string{Read, Invite}
}
