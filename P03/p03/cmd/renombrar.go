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

// renombrarCmd represents the renombrar command
var renombrarCmd = &cobra.Command{
	Use:   "renombrar [NombreFichero] [NuevoNombre]",
	Short: "Este comando renombra el nombre del fichero",
	Long:  `Por favor, verifica que el nombre del estado es el correcto.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, err := cmd.Flags().GetString("nombre_fichero")
		if err != nil {
			log.Fatal(err)
		}
		newFilename, err := cmd.Flags().GetString("nuevo_nombre")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Renombrando el fichero " + filename + " a " + newFilename)
		renombrarFichero(filename, newFilename)

	},
}

//renombrarFichero
func renombrarFichero(nombreFichero string, nuevoNombre string) {
	// get the file and copy it to json_files_manipulated
	// check if original file exists
	if _, err := os.Stat("json_files/" + nombreFichero + ".json"); os.IsNotExist(err) {
		log.Fatal("No existe el fichero")
	}
	fmt.Println("El fichero " + nombreFichero + ".json existe")
	// check if new file exists
	if _, err := os.Stat("json_files_manipulated/" + nuevoNombre + ".json"); os.IsNotExist(err) {
		fmt.Println("El fichero " + nuevoNombre + ".json no existe")
		// copy the file
		copyFile("json_files/"+nombreFichero+".json", "json_files_manipulated/"+nuevoNombre+".json")
		fmt.Println("Se copio el fichero " + nombreFichero + ".json a " + nuevoNombre + ".json")
	} else {
		fmt.Println("El fichero " + nuevoNombre + ".json ya existe")
	}
}

// COPY FILE
func copyFile(src string, dest string) {
	in, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	out, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		log.Fatal(err)
	}
	err = out.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
func init() {
	rootCmd.AddCommand(renombrarCmd)
	renombrarCmd.Flags().StringP("nombre_fichero", "o", "", "Nombre del fichero")     //old name
	renombrarCmd.Flags().StringP("nuevo_nombre", "n", "", "Nuevo nombre del fichero") //new name
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renombrarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renombrarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
