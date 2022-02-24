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
		if _, err := os.Stat("json_files"); os.IsNotExist(err) {
			log.Fatal("No existe el directorio json_files")
		}
		fmt.Println("La carpeta con los archivos json_files existe")
		// create directory where manipulated files will go
		if _, err := os.Stat("json_files_manipulated"); os.IsNotExist(err) {
			os.Mkdir("json_files_manipulated", 0777)
		}
		fmt.Println("Se creo o se comprobó la existencia del directorio json_files_manipulated")

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
