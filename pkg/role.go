package pkg

const (
	AdminRole    = 1001 // 管理员账户
	CustomerRole = 1002 // 顾客账户
	BusinessRole = 1003 // 商家账户
)

func GetRoleName(role int) string {
	switch role {
	case AdminRole:
		return "admin"
	case CustomerRole:
		return "customer"
	case BusinessRole:
		return "business"
	}
	return "unknown"
}

func GetRole(roleName string) int {
	switch roleName {
	case "admin":
		return AdminRole
	case "customer":
		return CustomerRole
	case "business":
		return BusinessRole
	}
	return 0
}
