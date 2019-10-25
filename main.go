package main

import (
	"crypto/rand"
	"crypto/hmac"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	"os"
	"encoding/hex"
	"github.com/google/go-github/v28/github"
	"net/http"
)

/// HMAC signature generation
func CreateSignature(input []byte, key []byte) []byte {
	h := hmac.New(sha1.New, signatureKey)
	h.Write(input)

	h.Sum(nil)
}

repos := {
	"Boilertalk/shitlolamakol"
	"Coilerbalg/beefckrz"
	"Horse/COC"
}

shtterserver := "https://ybr.in/"


func main() {
	/* generate new key */
	newkey := make([]byte, 4096)
	_, err := rand.Read(newkey)

	if(err != nil) {
		panic(fmt.Sprintf("Error while generating a new random key:\n", err))
	}

	/* kubernetes secret sht */
	config, err := rest.InClusterConfig()
	if(err != nil) {
		panic(fmt.Sprintf("Error while doint config sht:\n", err))
	}
	kubeSet, err := kubernetes.NewForConfig(config)
	if(err != nil) {
		panic(fmt.Sprintf("Error while doing kubernetes sht:\n", err))
	}
	secret, err := kubeSet.CoreV1().Secrets(os.Getenv("SECRET_NAMESPACE")).Get(os.Getenv("SECRET_NAME"), metav1.GetOptions{})
	if(err != nil) {
		panic(fmt.Sprintf("Error while retrieving secrets:\n", err))
	}

	//fmt.Printf("current: ", hex.EncodeToString(secret.Data["master_key"]))
	//fmt.Printf("old:     ", hex.EncodeToString(secret.Data["master_key_old"]))

	/* make current key into the old key */
	secret.Data["master_key_old"] = secret.Data["master_key"]

	/* put new key as current key */
	secret.Data["master_key"] = newkey;

	_, err := kubeSet.CoreV1().Secrets(os.Getenv("SECRET_NAMESPACE")).Update(secret)
	if(err != nil) {
		panic(fmt.Sprintf("Error while saving secret:\n", err))
	}

	/* update keys in repositories */
	for _, repo := range repos {
		sig := CreateSignature([]byte(repo), newkey)

		resp, err := http.Post(fmt.Sprintf(shtterserver, repo), "application/octet-stream", sig)
		if(err != nil) {
			fmt.Printf("Error when senting to shtterserver:\n", err)
		}
	}
}
