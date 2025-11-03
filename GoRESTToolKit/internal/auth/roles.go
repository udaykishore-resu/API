package auth

// Define roles for RBAC
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
	RoleGuest = "guest"
)

func HasAnyRole(userRoles []string, allowedRoles []string) bool {
	roleSet := make(map[string]bool)
	for _, r := range userRoles {
		roleSet[r] = true
	}
	for _, allowed := range allowedRoles {
		if roleSet[allowed] {
			return true
		}
	}
	return false
}
