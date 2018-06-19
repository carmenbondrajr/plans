package test

import (
	"plans"
	"testing"
)

func TestUserMap(t *testing.T) {
	userMap := plans.UserMap{
		"12345": []plans.User{
			{UserId: "user1", Name: "Nick Foles", Address: "4133 Superbowl Blvd, Philadelphia, PA 12345-9999"},
		},
		"92929": []plans.User{
			{UserId: "user2", Name: "Johnny Cash", Address: "19 Road Street, Hometown, PA 92929"},
			{UserId: "user3", Name: "Claude Giroux", Address: "140 Apple Road, Hometown, PA 92929"},
		},
		"44444": []plans.User{
			{UserId: "user4", Name: "Alexander Hamilton", Address: "76 Reynolds Drive, Weehawken, NJ 44444"},
		},
	}

	inputOutput := map[string]string{
		"12345": "Residents:\n" +
			"user1 - Nick Foles - 4133 Superbowl Blvd, Philadelphia, PA 12345-9999",
		"92929": "Residents:\n" +
			"user2 - Johnny Cash - 19 Road Street, Hometown, PA 92929\n" +
			"user3 - Claude Giroux - 140 Apple Road, Hometown, PA 92929",
		"44444": "Residents:\n" +
			"user4 - Alexander Hamilton - 76 Reynolds Drive, Weehawken, NJ 44444",
		"90210": "No Residents.",
	}

	for in := range inputOutput {
		expected := inputOutput[in]
		actual := userMap.GetUsersFromZip(in)
		if actual != expected {
			t.Errorf("Formatting Plans from map failed for Zip '%s': "+
				"\nExpected \n%s but got \n%s ", in, expected, actual)
		}
	}

}
