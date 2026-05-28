package main

import (
	"encoding/json"
	"sample/src/api/middleware"
	"fmt"
	"os"
)

const permFile = "permissions.json"

func main() {
	existing := loadPermissions()

	for _, p := range middleware.AllPermissions {
		if !contains(existing, string(p)) {
			fmt.Printf("Adding permission: %s\n", p)
			existing = append(existing, string(p))
		}
	}

	savePermissions(existing)
	fmt.Println("Permissions synced!")
}

func loadPermissions() []string {
	file, err := os.ReadFile(permFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}
		}
		panic(err)
	}
	var perms []string
	if err := json.Unmarshal(file, &perms); err != nil {
		panic(err)
	}
	return perms
}

func savePermissions(perms []string) {
	data, _ := json.MarshalIndent(perms, "", "  ")
	os.WriteFile(permFile, data, 0644)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
