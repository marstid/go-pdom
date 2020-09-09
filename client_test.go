package gopingdom

import (
	"fmt"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	fmt.Println("Test: Client Connect")

	c, err := NewRestClient(os.Getenv("PINGDOM_TOKEN"), true, 3)
	if err != nil {
		t.Error(err)
	}

	checks, err := c.UptimeGetChecks()
	if err != nil {
		t.Error(err)
	}
	fmt.Println()
	//sort.Sort(checks)
	for _, check := range checks {
		if check.Status != "paused" {
			fmt.Printf("ID: %d, Name: %s, Host: %s, Status: %s\n", check.ID, check.Name, check.Hostname, check.Status)
		}

	}
	check2, err := c.UptimeGetCheckDetails(checks[0].ID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println()
	fmt.Printf("ID: %d, Name: %s, Host: %s, Status: %s\n", check2.ID, check2.Name, check2.Hostname, check2.Status)

}
