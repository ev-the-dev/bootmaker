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

	generateModule(cwd, answers)

	if answers.Controller {
		wg.Add(1)
		go generateController(cwd, answers)
	}

	if answers.QueueConsumer {
		wg.Add(1)
		go generateQueueConsumer(cwd, answers)
	}

	if answers.Repository {
		wg.Add(1)
		go generateRepository(cwd, answers)
	}

	if answers.Service {
		wg.Add(1)
		go generateService(cwd, answers)
	}
}

func generateController(cwd string, answers *models.WizardAnswers) error {
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Controller"))
	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	mkDirErr := os.MkdirAll(fmt.Sprintf("%s/controllers/", modulePath), 0744)
	if mkDirErr != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create controller directory"), mkDirErr)
		os.Exit(1)
	}

	w, err := os.Create(fmt.Sprintf("%s/controllers/%s.controller.ts", modulePath, answers.ModuleName))
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

func generateModule(cwd string, answers *models.WizardAnswers) error {
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Module"))
	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	mkDirErr := os.MkdirAll(modulePath, 0744)
	if mkDirErr != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create module directory"), mkDirErr)
		os.Exit(1)
	}

	w, err := os.Create(fmt.Sprintf("%s/%s.module.ts", modulePath, answers.ModuleName))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create module file"), err)
		os.Exit(1)
	}
	defer w.Close()

	mTmpl := template.Must(template.New("module").Funcs(tempFuncs).Parse(ct.ModuleTemplate))
	mTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Module"))
	return nil
}

func generateQueueConsumer(cwd string, answers *models.WizardAnswers) error {
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Queue Consumer"))
	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	mkDirErr := os.MkdirAll(fmt.Sprintf("%s/queues/", modulePath), 0744)
	if mkDirErr != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create queues directory"), mkDirErr)
		os.Exit(1)
	}

	w, err := os.Create(fmt.Sprintf("%s/queues/consumers.queues.ts", modulePath))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create queue consumers file"), err)
		os.Exit(1)
	}
	defer w.Close()

	rTmpl := template.Must(template.New("queue-consumer").Funcs(tempFuncs).Parse(ct.QueueConsumerTemplate))
	rTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Queue Consumer"))
	wg.Done()
	return nil
}

func generateRepository(cwd string, answers *models.WizardAnswers) error {
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Repository"))
	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	mkDirErr := os.MkdirAll(fmt.Sprintf("%s/repository/", modulePath), 0744)
	if mkDirErr != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create repository directory"), mkDirErr)
		os.Exit(1)
	}

	w, err := os.Create(fmt.Sprintf("%s/repository/%s.repository.ts", modulePath, answers.ModuleName))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create repository file"), err)
		os.Exit(1)
	}
	defer w.Close()

	rTmpl := template.Must(template.New("repository").Funcs(tempFuncs).Parse(ct.RepositoryTemplate))
	rTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Repository"))
	wg.Done()
	return nil
}

func generateService(cwd string, answers *models.WizardAnswers) error {
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Service"))
	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	mkDirErr := os.MkdirAll(fmt.Sprintf("%s/services/", modulePath), 0744)
	if mkDirErr != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create services directory"), mkDirErr)
		os.Exit(1)
	}

	w, err := os.Create(fmt.Sprintf("%s/services/%s.service.ts", modulePath, answers.ModuleName))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create service file"), err)
		os.Exit(1)
	}
	defer w.Close()

	sTmpl := template.Must(template.New("service").Funcs(tempFuncs).Parse(ct.ServiceTemplate))
	sTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Service"))
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

func formatModuleNameEnum(name string) string {
	splitName := strings.Split(name, "-")
	for i, v := range splitName {
		splitName[i] = strings.ToUpper(v)
	}
	return strings.Join(splitName, "_")
}
