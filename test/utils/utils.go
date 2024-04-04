/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo/v2" //nolint:golint,revive
)

// Run executes the provided command within this context
func Run(cmd *exec.Cmd) ([]byte, error) {
	dir, err := GetProjectDir()
	if err != nil {
		return nil, err
	}
	cmd.Dir = dir

	if err := os.Chdir(cmd.Dir); err != nil {
		fmt.Fprintf(GinkgoWriter, "chdir dir: %s\n", err)
	}

	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	command := strings.Join(cmd.Args, " ")
	fmt.Fprintf(GinkgoWriter, "running: %s\n", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("%s failed with error: (%v) %s", command, err, string(output))
	}

	return output, nil
}

// MinikubeSetProfile sets the minikube profile to the given name
func MinikubeSetProfile(name string) error {
	opts := []string{"profile", name}
	cmd := exec.Command("minikube", opts...)
	_, err := Run(cmd)
	return err
}

// LoadImageToMinikubeProfileWithName loads the host image to the nodes of the
// given minikube profile.
func LoadImageToMinikubeProfileWithName(image string) error {
	// We cannot use `minikube image load ...` on recent Docker versions.
	// See [1] for more details. Until [1] is resolved we will use this
	// workaround instead.
	//
	// [1]: https://github.com/kubernetes/minikube/issues/18021
	archiveName := "image.tar"
	dockerOpts := []string{"image", "save", "-o", archiveName, image}
	dockerCmd := exec.Command("docker", dockerOpts...)
	if _, err := Run(dockerCmd); err != nil {
		return err
	}
	defer os.Remove(archiveName)

	minikubeOpts := []string{"image", "load", "--overwrite", "true", archiveName}
	minikubeCmd := exec.Command("minikube", minikubeOpts...)
	if _, err := Run(minikubeCmd); err != nil {
		return err
	}

	return nil
}

// GetNonEmptyLines converts given command output string into individual objects
// according to line breakers, and ignores the empty elements in it.
func GetNonEmptyLines(output string) []string {
	var res []string
	elements := strings.Split(output, "\n")
	for _, element := range elements {
		if element != "" {
			res = append(res, element)
		}
	}

	return res
}

// GetProjectDir will return the directory where the project is
func GetProjectDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return wd, err
	}
	wd = strings.Replace(wd, "/test/e2e", "", -1)
	return wd, nil
}

// Make invokes `make` with the given args
func Make(args ...string) error {
	cmd := exec.Command("make", args...)
	_, err := Run(cmd)

	return err
}

// Kubectl invokes `kubectl` with the given args
func Kubectl(args ...string) error {
	cmd := exec.Command("kubectl", args...)
	_, err := Run(cmd)

	return err
}
