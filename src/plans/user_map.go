package plans

import (
	"fmt"
	"strings"
)

type UserMap map[string][]User

func (r UserMap) GetUsersFromZip(zip string) string {
	userSlice := r[zip]

	users := make([]string, 0)
	for _, u := range userSlice {
		users = append(users, fmt.Sprintf("%s - %s - %s", u.UserId, u.Name, u.Address))
	}

	if len(users) > 0 {
		return fmt.Sprintf("Residents:\n%s", strings.Join(users, "\n"))
	}

	return "No Residents."
}
