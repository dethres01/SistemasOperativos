/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// eliminarCmd represents the eliminar command
var eliminarCmd = &cobra.Command{
	Use:   "eliminar [NombreFichero][CodigoPostal]",
	Short: "Este comando elimina del fichero todos los registros que coincidan con el [CodigoPostal] ingresado",
	Long:  `Este comando elimina del fichero todos los registros que coincidan con el [CodigoPostal] ingresado.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, err := cmd.Flags().GetString("nombre_fichero")
		if err != nil {
			log.Fatal(err)
		}
		postalCode, err := cmd.Flags().GetString("codigo_postal")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Eliminando del fichero " + filename + " todos los registros con el codigo postal " + postalCode)
		eliminarRegistros(filename, postalCode)

	},
}

func eliminarRegistros(filename string, postalcode string) {
	if _, err := os.Stat("json_files/" + filename + ".json"); os.IsNotExist(err) {
		log.Fatal("No existe el fichero")
	}
	// we want to read the file
	data, err := ioutil.ReadFile("json_files/" + filename + ".json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make(map[string]interface{})

	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	// test
	// print  a field
	fmt.Println(fields[postalcode])
	// mantain format of the file
	// delete the field
	delete(fields, postalcode)
	json_data, err := json.MarshalIndent(fields, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("json_files/"+filename+".json", json_data, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
func init() {
	rootCmd.AddCommand(eliminarCmd)
	eliminarCmd.Flags().StringP("nombre_fichero", "n", "", "Nombre del fichero") //name of the file
	eliminarCmd.Flags().StringP("codigo_postal", "c", "", "Codigo Postal")       //postalcode
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eliminarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eliminarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
