package it

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
	"time"

	commands "github.com/adamdecaf/vault-backend-migrator/cmd"
	"github.com/adamdecaf/vault-backend-migrator/vault"
)

var (
	vaultVersion = "0.9.1"
)

func hasDocker() bool {
	err := exec.Command("docker", "version").Run()
	return err == nil
}

// Quick sanity check
func TestMigrator__integration(t *testing.T) {
	if !hasDocker() {
		t.Skip("docker isn't installed / running")
	}

	// Start vault container
	cmd := exec.Command("docker", "run", "-d", "-p", "8200:8200", "-t", fmt.Sprintf("vault:%s", vaultVersion))
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	// Grab root token
	r := regexp.MustCompile(`Root Token: ([a-f0-9]{8}\-[a-f0-9]{4}\-[a-f0-9]{4}\-[a-f0-9]{4}\-[a-f0-9]{12})`)
	id := strings.TrimSpace(stdout.String())
	defer func() {
		err = exec.Command("docker", "kill", id).Run()
		if err != nil {
			t.Fatal(err)
		}
	}()
	var token string
	for {
		out, err := exec.Command("docker", "logs", id).CombinedOutput()
		if err != nil {
			t.Fatal(err)
			break
		}
		loc := r.FindIndex(out)
		if len(loc) > 0 {
			s := string(out[loc[0]:loc[1]])
			token = strings.TrimPrefix(s, "Root Token: ")
			break
		}
		time.Sleep(1 * time.Second)
	}

	if token == "" {
		t.Fatal("empty token")
	}

	// Write a couple values into secret/, export, delete and import
	data := []struct {
		path  string
		key   string
		value interface{}
	}{
		{"secret/foo", "foo", "YmFyCg=="},          // bar
		{"secret/bar/baz", "username", "YWRhbQo="}, // adam
		{"secret/baz", "integer", 100},
	}
	os.Setenv("VAULT_ADDR", "http://127.0.0.1:8200")
	os.Setenv("VAULT_TOKEN", token)
	client, err := vault.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	// write values
	for i := range data {
		m := make(map[string]interface{})

		switch d := data[i].value.(type) {
		case int:
			m[data[i].key] = d
		case string:
			m[data[i].key] = d
		default:
			t.Fatal("Error: unsupported data type")
		}
		client.Write(data[i].path, m, "1")

		kv := client.Read(data[i].path)
		if kv[data[i].key] != data[i].value {
			t.Fatalf("path=%q, kv[%s]=%q, value=%q, err=%v", data[i].path, data[i].key, kv[data[i].key], data[i].value, err)
		}
	}

	// export
	tmp, err := ioutil.TempFile("", "vault-backend-migrator")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	err = commands.Export("secret/", tmp.Name(), "", "1")
	if err != nil {
		t.Fatal(err)
	}

	// delete
	for i := range data {
		client.Client().Logical().Delete(data[i].path)
		// read to verify it's gone
		kv, err := client.Client().Logical().Read(data[i].path)
		if err == nil && kv != nil {
			t.Fatalf("path=%q, kv=%v", data[i].path, kv)
		}
	}

	// import
	err = commands.Import("secret/", tmp.Name(), "1")
	if err != nil {
		t.Fatal(err)
	}
	for i := range data {
		kv := client.Read(data[i].path)
		if kv[data[i].key] != data[i].value {
			t.Fatalf("path=%q, kv[%s]=%q, value=%q, err=%v", data[i].path, data[i].key, kv[data[i].key], data[i].value, err)
		}
	}
}
