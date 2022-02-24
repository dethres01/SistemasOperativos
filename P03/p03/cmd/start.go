/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicializa los archivos y directorios para la aplicacion",
	Long:  `Copia los archivos de la carpeta json_files para poder manipular sin perder la informacion original.`,
	Run: func(cmd *cobra.Command, args []string) {
		// one part of me wants to keep data integrity but maybe I'm overthinking it and we don't need it for the moment
		if _, err := os.Stat("json_files"); os.IsNotExist(err) {
			log.Fatal("No existe el directorio json_files")
		}
		fmt.Println("La carpeta con los archivos json_files existe")

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
