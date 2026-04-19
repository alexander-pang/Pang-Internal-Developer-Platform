package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/go-chi/chi/v5"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()
	r.Get("/api/namespaces", func(w http.ResponseWriter, r *http.Request) {
		namespaces, err := client.CoreV1().Namespaces().List(r.Context(), metav1.ListOptions{})
		// handle err, then write namespaces as JSON
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(namespaces)

	})

	r.Get("/api/namespaces/{name}/pods", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		pods, err := client.CoreV1().Pods(name).List(r.Context(), metav1.ListOptions{})
		// handle err, then write namespaces as JSON
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pods)
	})

	log.Printf("Successfully created Kubernetes client")

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "ok"}`)
	})
	// create the server explicitly
	srv := &http.Server{Addr: ":9090", Handler: r}

	// run it in a goroutine so it doesn't block
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// wait for a signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
}
