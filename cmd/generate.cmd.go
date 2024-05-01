package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/ev-the-dev/bootmaker/models"
	ct "github.com/ev-the-dev/bootmaker/templates"
)

var wg sync.WaitGroup

var tempFuncs = template.FuncMap{
	"formatModuleName": formatModuleName,
}

func generateFiles(answers *models.WizardAnswers) {
	defer wg.Wait()

	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		log.Fatalf("Unable to get Current Working Directory: %v", cwdErr)
	}

	if answers.Controller {
		wg.Add(1)
		go generateController(cwd, answers)
	}
}

func generateController(cwd string, answers *models.WizardAnswers) error {
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Controller"))
	mkDirErr := os.MkdirAll(fmt.Sprintf("%s/controllers/", cwd), 0744)
	if mkDirErr != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create controller directory"), mkDirErr)
		os.Exit(1)
	}

	w, err := os.Create(fmt.Sprintf("%s/controllers/%s.controller.ts", cwd, answers.ModuleName))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create controller file"), err)
		os.Exit(1)
	}
	defer w.Close()

	cTmpl := template.Must(template.New("controller").Funcs(tempFuncs).Parse(ct.ControllerTemplate))
	cTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Controller"))
	wg.Done()
	return nil
}

func greenText(text string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", text)
}

func redText(text string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", text)
}

func yellowText(text string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", text)
}

func formatModuleName(name string) string {
	splitName := strings.Split(name, "-")
	for i, v := range splitName {
		splitName[i] = strings.Title(v)
	}
	return strings.Join(splitName, "")
}
