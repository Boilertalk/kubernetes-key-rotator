package main

import (
	"crypto/rand"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	"os"
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


func main() {
	/* generate new key */
	newkey := make([]byte, 4096);
	_, err := rand.Read(newkey);

	if(err == nil) {
		panic(fmt.Sprintf("Error while generating a new random key:\n", err))
	}

	/* kubernetes secret sht */
	config, err := rest.InClusterConfig()
	if(err != nil) {
		panic(fmt.Sprintf("Error while doint config sht:\n", err));
	}
	kubeSet, err := kubernetes.NewForConfig(config)
	if(err != nil) {
		panic(fmt.Sprintf("Error while doing kubernetes sht:\n", err))
	}
	secret, err := kubeSet.CoreV1().Secrets(os.Getenv("SECRET_NAMESPACE")).Get(os.Getenv("SECRET_NAME"), metav1.GetOptions{})
	if(err != nil) {
		panic(fmt.Sprintf("Error while retrieving secrets:\n", err))
	}

	fmt.Printf("current: ", hex.EncodeToString(secret.Data["master_key"]))
	fmt.Printf("old:     ", hex.EncodeToString(secret.Data["master_key_old"]));

	/* make current key into the old key */

	/* put new key as current key */

	/* update keys in repositories */
}
