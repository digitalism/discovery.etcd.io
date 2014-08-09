package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"net/http"
	"os"
	"path"
)

func generateCluster() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}

func setupToken() (string, error) {
	token := generateCluster()
	if token == "" {
		return "", errors.New("Couldn't generate a token")
	}

	machines := make([]string, 0)

	origin_addr := os.Getenv("DISCOVERY_ORIGIN_ADDR")

	if origin_addr != "" {
		machines = append(machines, origin_addr)
	}
	client := etcd.NewClient(machines)

	key := path.Join("_etcd", "registry", token)
	resp, err := client.CreateDir(key, 0)

	if err != nil || resp.Node == nil || resp.Node.Key != "/"+key || resp.Node.Dir != true {
		return "", errors.New(fmt.Sprintf("Couldn't setup state %v %v", resp, err))
	}

	return token, nil
}

func deleteToken(token string) error {
	client := etcd.NewClient(nil)

	if token == "" {
		return errors.New("No token given")
	}

	_, err := client.Delete(path.Join("_etcd", "registry", token), true)

	return err
}

func NewTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := setupToken()

	if err != nil {
		log.Printf("setupToken returned: %v", err)
		http.Error(w, "Unable to generate token", 400)
		return
	}

	log.Println("New cluster created", token)

	var domain = os.Getenv("DISCOVERY_ROOT_URL")
	if domain == "" {
		domain = "https://discovery.etcd.io"
	}

	fmt.Fprintf(w, path.Join(domain, token))
}
