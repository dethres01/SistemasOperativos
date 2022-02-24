/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// agruparCmd represents the agrupar command
var agruparCmd = &cobra.Command{
	Use:   "agrupar [NombreFichero1] [ NombreFichero2] ... [NombreFichero5]",
	Short: "Este comando agrupa 5 ficheros cualquiera dentro de una carpeta con el nombre [Agrupacion]",
	Long: `Este comando agrupa 5 ficheros cualquiera dentro de una carpeta con el nombre [Agrupacion], 
	verificar si ya existe una carpeta, de ser así el nombre de las nuevas carpetas será [Agrupación2], [Agrupación3], etc...`,
	Run: func(cmd *cobra.Command, args []string) {
		// vamos a asumir que es acción de copiar en lugar de mover
		// get the flags
		//vamos a asumir que siempre son 5 ficheros
		archivo1, err := cmd.Flags().GetString("archivo1")
		if err != nil {
			log.Fatal(err)
		}
		archivo2, err := cmd.Flags().GetString("archivo2")
		if err != nil {
			log.Fatal(err)
		}
		archivo3, err := cmd.Flags().GetString("archivo3")
		if err != nil {
			log.Fatal(err)
		}
		archivo4, err := cmd.Flags().GetString("archivo4")
		if err != nil {
			log.Fatal(err)
		}
		archivo5, err := cmd.Flags().GetString("archivo5")
		if err != nil {
			log.Fatal(err)
		}
		// checkFiles
		checkFiles(archivo1, archivo2, archivo3, archivo4, archivo5)
		// if we don't have any errors means files are ok
		counter := 0
		// check if Agrupacion folder exists
		// Create flag iteration because we need to check if the folder exists and if it does not, create it
		for {
			if counter != 0 {
				// if counter is not 0, means we have to create a new folder
				// create the folder
				if _, err := os.Stat("Agrupacion" + fmt.Sprintf("%d", counter)); os.IsNotExist(err) {
					os.Mkdir("Agrupacion"+fmt.Sprintf("%d", counter), 0777)
					groupUp(archivo1, archivo2, archivo3, archivo4, archivo5, fmt.Sprintf("Agrupacion%d", counter))
					break
				} else {
					counter++
				}
			} else {
				if _, err := os.Stat("Agrupacion"); os.IsNotExist(err) {
					os.Mkdir("Agrupacion", 0777)
					//since we created the folder we can break the loop *and* add the files
					groupUp(archivo1, archivo2, archivo3, archivo4, archivo5, "Agrupacion")
					break

				} else {
					counter++
				}
			}
		}
	},
}

// groupUp will copy the files to the new folder
func groupUp(archivo1, archivo2, archivo3, archivo4, archivo5, folder string) {
	// copy the files to the new folder
	copyFile("json_files/"+archivo1+".json", folder+"/"+archivo1+".json")
	copyFile("json_files/"+archivo2+".json", folder+"/"+archivo2+".json")
	copyFile("json_files/"+archivo3+".json", folder+"/"+archivo3+".json")
	copyFile("json_files/"+archivo4+".json", folder+"/"+archivo4+".json")
	copyFile("json_files/"+archivo5+".json", folder+"/"+archivo5+".json")
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) {
	// open the source file
	source, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()
	// open the destination file
	destination, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer destination.Close()
	// copy the content of the file
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
	}
}

// checkFiles
func checkFiles(archivo1, archivo2, archivo3, archivo4, archivo5 string) {
	// files are at json_files folder and are .json
	if _, err := os.Stat("json_files/" + archivo1 + ".json"); os.IsNotExist(err) {
		log.Fatal("No existe el archivo " + archivo1)
	}
	if _, err := os.Stat("json_files/" + archivo2 + ".json"); os.IsNotExist(err) {
		log.Fatal("No existe el archivo " + archivo2)
	}
	if _, err := os.Stat("json_files/" + archivo3 + ".json"); os.IsNotExist(err) {
		log.Fatal("No existe el archivo " + archivo3)
	}
	if _, err := os.Stat("json_files/" + archivo4 + ".json"); os.IsNotExist(err) {
		log.Fatal("No existe el archivo " + archivo4)
	}
	if _, err := os.Stat("json_files/" + archivo5 + ".json"); os.IsNotExist(err) {
		log.Fatal("No existe el archivo " + archivo5)
	}
}
func init() {
	rootCmd.AddCommand(agruparCmd)
	agruparCmd.Flags().StringP("archivo1", "1", "", "nombre de archivo 1")
	agruparCmd.Flags().StringP("archivo2", "2", "", "nombre de archivo 2")
	agruparCmd.Flags().StringP("archivo3", "3", "", "nombre de archivo 3")
	agruparCmd.Flags().StringP("archivo4", "4", "", "nombre de archivo 4")
	agruparCmd.Flags().StringP("archivo5", "5", "", "nombre de archivo 5")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agruparCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agruparCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
