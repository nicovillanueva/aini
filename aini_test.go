package aini

import (
	"strings"
	"testing"
)

func input1() string {
	return `
myhost1
[dbs]
dbhost1
dbhost2

[apps]
my-app-server1
my-app-server2

`
}

func createHosts(input string) Hosts {
	testInput := strings.NewReader(input)
	v, _ := NewParser(testInput)
	v.Parse()
	return *v
}

func TestGroupExists(t *testing.T) {
	v := createHosts(input1())
	matched := false
	if _, ok := v.Groups["dbs"]; ok {
		matched = true
	}
	if !matched {
		t.Error("Expected to find the group \"dbs\"")
	}
}

func TestHostExistsInGroups(t *testing.T) {
	v := createHosts(input1())
	exportedHosts := map[string][]Host{
		"dbs": []Host{Host{Name: "dbhost1"},
			Host{Name: "dbhost2"}},
		"ungrouped": []Host{Host{Name: "myhost1"}},
	}

	for group, ehosts := range exportedHosts {
		for _, ehost := range ehosts {
			if hosts, ok := v.Groups[group]; ok {
				matched := false
				for _, host := range hosts {
					if host.Name == ehost.Name {
						matched = true
					}
				}
				if !matched {
					t.Error("Server ", ehost.Name, " was not found in ", group)
				}
			} else {
				t.Error(group, " group doesn't exist")
			}
		}

	}
}
