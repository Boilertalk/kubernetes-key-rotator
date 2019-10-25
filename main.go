package main

import (
	"crypto/rand"
	"k8s.io/api/core/v1"
	"fmt"
	"encoding/hex"
)

type MessageGithub struct {
	Sha        string `json:"sha"`
	Repository string `json:"repository"`
	Ref        string `json:"ref"`
}

type Message struct {
	Github MessageGithub `json:"github"`
	Image  string        `json:"image"`
}

type ResponseMessage struct {
	Success bool   `json:"error"`
	Message string `json:"message"`
}

// GLOBAL VARIABLES
var hmacSecret string
var slackWebhookUrl string
var globalLogger *logger.Logger
var kubeSet *kubernetes.Clientset

/// HMAC signature generation
func CreateSignature(input []byte, key string) string {
	signatureKey := []byte(key)

	h := hmac.New(sha1.New, signatureKey)
	h.Write(input)

	return "sha1=" + hex.EncodeToString(h.Sum(nil))
}




func main() {
	/* generate new key */
	newkey := make([]byte, 4096);
	_, err := rand.Read(b);

	if(err == nil) {
		panic(fmt.Sprintf("Error while generating a new random key:\n", err))
	}

	/* kubernetes secret sht */
	kubeSet, err := kubernetes.NewForConfig(config)
	if(err != nil) {
		panic(fmt.Sprintf("Error while doing kubernetes sht:\n", err))
	}
	secret, err := kubeSet.CoreV1().Secrets(os.Getenv("SECRET_NAMESPACE")).Get(os.Getenv("SECRET_NAME"))
	if(err != nil) {
		panic(fmt.Sprintf("Error while retrieving secrets:\n", err))
	}

	fmt.Printf("current: ", hex.EncodeToString(secret.Data["master_key"]))
	fmt.Printf("old:     ", hex.EncodeToString(secret.Data["master_key_old"]));

	/* make current key into the old key */

	/* put new key as current key */

	/* update keys in repositories */
}
