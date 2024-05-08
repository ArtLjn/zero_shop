package pkg

const (
	AdminRole    = 1001 // 管理员账户
	CustomerRole = 1002 // 顾客账户
	SellerRole   = 1003 // 商家账户
)

func GetRoleName(role int) string {
	switch role {
	case AdminRole:
		return "admin"
	case CustomerRole:
		return "customer"
	case SellerRole:
		return "seller"
	}
	return "unknown"
}

func GetRole(roleName string) int {
	switch roleName {
	case "admin":
		return AdminRole
	case "customer":
		return CustomerRole
	case "seller":
		return SellerRole
	}
	return 0
}
