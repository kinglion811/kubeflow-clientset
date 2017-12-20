/*
Copyright 2016 The Kubernetes Authors.

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

package main

import (
	"path/filepath"

	"github.com/caicloud/kubeflow-clientset/cmd/lister-gen/generators"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"k8s.io/gengo/args"
)

func main() {
	arguments := &args.GeneratorArgs{
		// Note: I DO NOT know why it DOES NOT work via cli.
		InputDirs: []string{
			"github.com/caicloud/kubeflow-clientset/apis/kubeflow/v1alpha1",
		},
		OutputBase:        args.DefaultSourceTree(),
		OutputPackagePath: "github.com/caicloud/kubeflow-clientset/listers",
		GoHeaderFilePath:  filepath.Join(args.DefaultSourceTree(), "github.com/caicloud/kubeflow-clientset/hack/custom-boilerplate.go.txt"),
	}
	arguments.AddFlags(pflag.CommandLine)
	// Run it.
	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.V(2).Info("Completed successfully.")
}
