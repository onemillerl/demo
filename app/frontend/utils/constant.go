package utils

const ServiceName = "frontend"

type (
	SessionUserIdKey string
	SessionTokenKey  string
)

const UserIdKey = SessionUserIdKey("user_id")

const TokenKey = SessionTokenKey("token")
