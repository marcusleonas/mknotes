package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/marcusleonas/mknotes/internal"
	"github.com/spf13/cobra"
)

var (
	Name     string
	Template string
	Open     bool
)

func init() {
	rootCmd.AddCommand(newCommand)
	newCommand.Flags().StringVarP(&Name, "name", "n", "", "name")
	newCommand.Flags().StringVarP(&Template, "template", "t", "default.md", "template")
	newCommand.Flags().BoolVarP(&Open, "open", "o", false, "open")
	newCommand.MarkFlagRequired("name")
}

var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Create a new note using a template",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("config.toml"); err != nil {
			fmt.Println("error: no config.toml found in this directory")
			return
		}

		err := os.MkdirAll(filepath.Dir(Name), 0750)
		if err != nil {
			fmt.Println("error creating directories", err)
			return
		}

		var nameWithExtension string
		if strings.Contains(Name, ".md") {
			nameWithExtension = Name
		} else {
			nameWithExtension = Name + ".md"
		}

		f, err := os.Create(nameWithExtension)
		if err != nil {
			fmt.Println("error creating file", err)
			return
		}

		conf, err := internal.GetConfig()
		if err != nil {
			fmt.Println("error reading config")
			return
		}

		if _, err = os.Stat(path.Join(conf.TemplateDir, Template)); err != nil {
			fmt.Println("error: template does not exist!")
			return
		}

		tmpl, err := template.ParseFiles(path.Join(conf.TemplateDir, Template))
		if err != nil {
			fmt.Println("error parsing template:", err)
			return
		}

		data := struct {
			Name      string
			Timestamp string
		}{
			Name:      strings.Trim(filepath.Base(Name), ".md"),
			Timestamp: time.Now().Format(time.RFC3339),
		}

		err = tmpl.Execute(f, data)
		if err != nil {
			fmt.Println("error executing template:", err)
			return
		}

		defer f.Close()

		if Open {
			editor := os.Getenv("EDITOR")

			if editor == "" {
				if runtime.GOOS == "windows" {
					editor = "nodepad"
				} else {
					editor = "nano"
				}
			}

			cmd := exec.Command(editor, nameWithExtension)

			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("error opening in editor:", err)
				return
			}

		}

		fmt.Printf("successfully created note %s\n", filepath.Base(Name))
	},
}
