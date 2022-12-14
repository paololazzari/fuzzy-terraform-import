package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/paololazzari/fuzzy-terraform-import/internal/builder"
	"github.com/paololazzari/fuzzy-terraform-import/internal/util"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

// get all terraform resources available for import
func getTerraformResourcesToImport(working_dir string) *tfconfig.Module {
	module, _ := tfconfig.LoadModule(working_dir)
	_ = (*module).ManagedResources
	var cmd strings.Builder
	cmd.WriteString("terraform state list")
	stdout, _, err := shellout(cmd.String(), true)
	if err != nil {
		fmt.Printf("No terraform state found\n")
		os.Exit(1)
	}
	stateResources := strings.Split(stdout, "\n")
	for _, resource := range stateResources {
		if _, ok := (*module).ManagedResources[resource]; ok {
			delete((*module).ManagedResources, resource)
		}
	}
	return module
}

// get the current working directory
func getDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	working_dir := filepath.Dir(ex)
	return working_dir
}

// build the fuzzyfinder menu
func fuzzyMenu(r []any) string {
	idx, err := fuzzyfinder.Find(
		r,
		func(i int) string {
			// resource must have either Name or Id as a property
			id := reflect.ValueOf(r[i]).FieldByName("Id")
			if !id.IsValid() {
				name := reflect.ValueOf(r[i]).FieldByName("Name")
				return name.String()
			}
			return id.String()
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			f := ""
			v := []interface{}{}
			f, v = util.FormatFuzzyInput(r[i])
			return fmt.Sprintf(f, v...)

		}))
	if err != nil {
		log.Fatal(err)
	}
	id := reflect.ValueOf(r[idx]).FieldByName("Id")
	if !id.IsValid() {
		name := reflect.ValueOf(r[idx]).FieldByName("Name")
		return name.String()
	}
	return id.String()
}

// execute the given command in either bash or powershell depending on the detected os
func shellout(command string, silent bool) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := &exec.Cmd{}
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-command", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if silent != true {
		fmt.Println(cmd.Stdout)
		fmt.Println(cmd.Stderr)
	}
	return stdout.String(), stderr.String(), err
}

func main() {

	workingDir := getDir()
	module := getTerraformResourcesToImport(workingDir)
	managedResources := (*module).ManagedResources

	if len(managedResources) == 0 {
		fmt.Printf("No resources available for import were found\n")
		os.Exit(1)
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	reader := bufio.NewReader(os.Stdin)
	for _, resource := range managedResources {
		fmt.Printf("Import %s.%s ? [y/n]:\n", resource.Type, resource.Name)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "y" {
				maps, exists := builder.GetObjects(resource.Type, sess)
				if exists == false {
					fmt.Printf("Resource %s is not supported for import\n", resource.Type)
					break
				}
				if len(maps) == 0 {
					fmt.Printf("%s: %s\n", "Did not find any resources of type", resource.Type)
					break
				}
				structs := util.BuildSliceOfStructs(maps)
				var cmd strings.Builder
				tfImportCommand := "terraform import " + resource.Type + "." + resource.Name + " "
				cmd.WriteString(tfImportCommand)
				cmd.WriteString(fuzzyMenu(structs))
				fmt.Printf("Executing: %s\n", cmd.String())
				shellout(cmd.String(), false)
				break
			} else if input == "n" {
				break
			} else {
				fmt.Printf("%s\n", "Invalid selection. Please select [y/n]")
			}
		}
	}
}
