package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Directory string
	Git       bool
)

const defaultTemplatesDir = ".templates"

var configTemplate = fmt.Sprintf(`template-dir="%s"`, defaultTemplatesDir)

var defaultTemplate = `# {{ .Name }}

Created at {{ .Timestamp }}`

func init() {
	rootCmd.AddCommand(initCommand)
	initCommand.Flags().StringVarP(&Directory, "directory", "d", "", "directory")
	initCommand.Flags().BoolVarP(&Git, "git", "g", false, "git")

	// Make sure that the directory flag is provided
	initCommand.MarkFlagRequired("directory")
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "initialise an empty vault",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initialising vault...")

		dirTrimmed := strings.TrimSpace(Directory) // Remove white space

		if dirTrimmed == "" {
			fmt.Println("error: please provide a directory name")
			return
		}

		err := os.Mkdir(dirTrimmed, 0750) // Create project dir
		if err != nil {
			fmt.Println("error creating vault directory:", err)
			return
		}

		err = os.Mkdir(path.Join(dirTrimmed, defaultTemplatesDir), 0750)
		if err != nil {
			fmt.Println("error creating template dir:", err)
			return
		}

		cf, err := os.Create(path.Join(dirTrimmed, "config.toml"))
		if err != nil {
			fmt.Println("error creating config.toml:", err)
			return
		}
		defer cf.Close()

		cf.WriteString(configTemplate)

		tf, err := os.Create(path.Join(dirTrimmed, defaultTemplatesDir, "default.md"))
		if err != nil {
			fmt.Println("error creating default template:", err)
			return
		}
		defer tf.Close()

		tf.WriteString(defaultTemplate)

		if Git {
			cmd := exec.Command("git", "init")
			cmd.Dir = Directory
			err := cmd.Run()
			if err != nil {
				fmt.Println("error initialising git repo:", err)
				return
			}
		}

		fmt.Printf("successfully initialised in %s\n", dirTrimmed)
	},
}
