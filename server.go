package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	client "docker.io/go-docker"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

var DockerCli *client.Client

func init() {
	DockerCli, _ = client.NewEnvClient()
}


func setAPIVersion() {
	cmd := exec.Command("docker", "version", "--format", "{{.Server.APIVersion}}")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	apiVersion := strings.TrimSpace(string(cmdOutput.Bytes()))
	fmt.Println(apiVersion)
	os.Setenv("DOCKER_API_VERSION", apiVersion)
}

func listFilesInDirectory(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func main() {
	setAPIVersion()
	mux := StartUp()

	handler := cors.Default().Handler(mux)
	n := negroni.Classic()
	n.UseHandler(handler)
	n.Run(":3000")

	//log.Fatal(http.ListenAndServe(":3000", handler))

	//listFilesInDirectory("/home/joerivrij/Projects")
}
