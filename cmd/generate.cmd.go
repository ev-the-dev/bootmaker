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
	"formatModuleName":     formatModuleName,
	"formatModuleNameEnum": formatModuleNameEnum,
}

/*
 * TODO: Will implement ways for DTOs to be conditionally defined, rather than defining all of them
 * at once in the DTO template file.
 *
 * TODO: Will condtionally populate files -- like the module file -- with the selected controllers,
 * providers, imports, adapters, etc.
 */
func generateFiles(answers *models.WizardAnswers) {
	defer wg.Wait()

	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		log.Fatalf("Unable to get Current Working Directory: %v", cwdErr)
	}

	generateModule(cwd, answers)

	if answers.Controller || answers.QueueConsumer || answers.Service || answers.Repository {
		generateAdapterDirectory(cwd, answers)
		generateDtoDirectory(cwd, answers)
	}

	if answers.Controller {
		wg.Add(2)
		go generateController(cwd, answers)
		go generateControllerAdapters(cwd, answers)
	}

	if answers.QueueConsumer {
		wg.Add(2)
		go generateQueueConsumer(cwd, answers)
		go generateQueueConsumerAdapters(cwd, answers)
	}

	if answers.Repository {
		wg.Add(2)
		go generateRepository(cwd, answers)
		go generateRepositoryAdapters(cwd, answers)
	}

	if answers.Service {
		wg.Add(2)
		go generateService(cwd, answers)
		go generateServiceAdapters(cwd, answers)
	}
}

func generateAdapterDirectory(cwd string, answers *models.WizardAnswers) error {
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Creating Adapter Directory"))
	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	mkDirErr := os.MkdirAll(fmt.Sprintf("%s/adapters/", modulePath), 0744)
	if mkDirErr != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create adapter directory"), mkDirErr)
		os.Exit(1)
	}
	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Creating Adapter Directory"))
	return nil
}

func generateController(cwd string, answers *models.WizardAnswers) error {
	defer wg.Done()
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
	return nil
}

func generateControllerAdapters(cwd string, answers *models.WizardAnswers) error {
	defer wg.Done()
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Controller Adapters"))

	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	w, err := os.Create(fmt.Sprintf("%s/adapters/controller.adapters.ts", modulePath))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create controller adapter file"), err)
		os.Exit(1)
	}
	defer w.Close()

	caTmpl := template.Must(template.New("controller-adapters").Funcs(tempFuncs).Parse(ct.ControllerAdapterTemplate))
	caTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Controller Adapters"))
	return nil
}

func generateDtoDirectory(cwd string, answers *models.WizardAnswers) error {
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Creating DTO Directory"))
	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	mkDirErr := os.MkdirAll(fmt.Sprintf("%s/dtos/", modulePath), 0744)
	if mkDirErr != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create DTO directory"), mkDirErr)
		os.Exit(1)
	}
	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Creating DTO Directory"))

	// NOTE: This DTO file creation is temporary until the conditional functionality explained
	// in the above TODO is in place.
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Creating DTO File"))
	w, err := os.Create(fmt.Sprintf("%s/dtos/%s.dtos.ts", modulePath, answers.ModuleName))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create DTO file"), err)
		os.Exit(1)
	}
	defer w.Close()

	dTmpl := template.Must(template.New("dto").Funcs(tempFuncs).Parse(ct.DtoTemplate))
	dTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Creating DTO File"))
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

	qcTmpl := template.Must(template.New("queue-consumer").Funcs(tempFuncs).Parse(ct.QueueConsumerTemplate))
	qcTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Queue Consumer"))
	wg.Done()
	return nil
}

func generateQueueConsumerAdapters(cwd string, answers *models.WizardAnswers) error {
	defer wg.Done()
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Queue Consumer Adapters"))

	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	w, err := os.Create(fmt.Sprintf("%s/adapters/queue-consumer.adapters.ts", modulePath))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create queue-consumer adapter file"), err)
		os.Exit(1)
	}
	defer w.Close()

	qcaTmpl := template.Must(template.New("queue-consumer-adapters").Funcs(tempFuncs).Parse(ct.QueueConsumerAdapterTemplate))
	qcaTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Queue Consumer Adapters"))
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

func generateRepositoryAdapters(cwd string, answers *models.WizardAnswers) error {
	defer wg.Done()
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Repository Adapters"))

	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	w, err := os.Create(fmt.Sprintf("%s/adapters/repository.adapters.ts", modulePath))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create repository adapter file"), err)
		os.Exit(1)
	}
	defer w.Close()

	raTmpl := template.Must(template.New("repository-adapters").Funcs(tempFuncs).Parse(ct.RepositoryAdapterTemplate))
	raTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Repository Adapters"))
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

func generateServiceAdapters(cwd string, answers *models.WizardAnswers) error {
	defer wg.Done()
	fmt.Printf("\n#####\n%s\n#####\n", yellowText("Generating Service Adapters"))

	modulePath := fmt.Sprintf("%s/%s/", cwd, answers.ModuleName)
	w, err := os.Create(fmt.Sprintf("%s/adapters/service.adapters.ts", modulePath))
	if err != nil {
		fmt.Printf("\n!!!!!\n%s\n!!!!!\nerr: %v\n", redText("Unable to create service adapter file"), err)
		os.Exit(1)
	}
	defer w.Close()

	saTmpl := template.Must(template.New("service-adapters").Funcs(tempFuncs).Parse(ct.ServiceAdapterTemplate))
	saTmpl.Execute(w, answers)

	fmt.Printf("\n#####\n%s\n#####\n", greenText("Done Generating Service Adapters"))
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
