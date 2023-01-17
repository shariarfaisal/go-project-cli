/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Create new service",
	Run: func(cmd *cobra.Command, args []string) {
		// make a folder with the name of the first argument,
		// and make files under this folder with rest of the arguments

		folder := fmt.Sprintf("%s", args[0])

		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Println(err)
		}

		for i := 1; i < len(args); i++ {
			f, err := os.Create(folder + "/" + folder + "." + args[i] + ".go")
			if err != nil {
				fmt.Println(err)
			}

			defer f.Close()

			data := "package %s"

			f.WriteString(fmt.Sprintf(data, folder))
		}

	},
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create new module",
	Run: func(cmd *cobra.Command, args []string) {
		// make a folder with the name of the first argument,
		// and make files under this folder with rest of the arguments

		folder := fmt.Sprintf("%s", args[0])

		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Println(err)
		}

		err = os.Mkdir(folder+"/service", 0755)
		if err != nil {
			fmt.Println(err)
		}

		for i := 1; i < len(args); i++ {

			// create model file
			f, err := os.Create(fmt.Sprintf("%s/%s.model.go", folder, args[i]))
			if err != nil {
				fmt.Println(err)
			}

			defer f.Close()

			data := fmt.Sprintf(
				`package %s
				
type %s struct {}
				`, folder, args[i],
			)

			f.WriteString(data)

			// create repository file
			f, err = os.Create(fmt.Sprintf("%s/%s.repo.go", folder, args[i]))
			if err != nil {
				fmt.Println(err)
			}

			defer f.Close()

			tName := strings.Title(args[i])
			repoName := tName + "Repo"
			data = fmt.Sprintf(`package %s

import "gorm.io/gorm"

type %s struct {
	DB *gorm.DB
} 

func New%s(db *gorm.DB) *%s {
	return &%s{DB: db}
}
			
			`, folder, repoName, repoName, repoName, repoName)

			f.WriteString(data)

			// create service file
			f, err = os.Create(fmt.Sprintf("%s/service/%s.go", folder, args[i]))
			if err != nil {
				fmt.Println(err)
			}

			defer f.Close()

			tName = strings.Title(args[i])
			serviceName := tName + "Service"
			data = fmt.Sprintf(
				`package service 

type %s struct {
	%s *%s
} 

func New%s(db *gorm.DB) *%s {
	return &%s{
		%s: %s(db),
	}
}
				
				`,
				serviceName, args[i], folder+"."+repoName, serviceName, serviceName, serviceName, args[i], folder+".New"+repoName,
			)

			f.WriteString(data)
		}
	},
}

func init() {
	rootCmd.AddCommand(moduleCmd)
	rootCmd.AddCommand(serviceCmd)
}
