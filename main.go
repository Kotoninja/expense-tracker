/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/Kotoninja/expense-tracker/cmd"
	storage "github.com/Kotoninja/expense-tracker/internal"
)

func main() {
	storageMain, err := storage.NewStorage("storage.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	storage.StorageIO = storageMain
	cmd.Execute()
}
