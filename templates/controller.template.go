package templates

// NOTE: if module contains hyphen (-) need to consolidate that
// to PascalCase
var ControllerTemplate = `
{{with $pMN := call formatModuleName .ModuleName}}
import { Controller } from "@nestjs/common"

import { {{.ModuleName}}ControllerAdapters } from "@{{.ModuleName}}/adapters/controller.adapters"

@Controller("{{.ModuleName}}")
export class {{.ModuleName}}Controller {
  public constructor(
    private readonly _adapter: {{.ModuleName}}ControllerAdapters 
  ) {}
}
{{end}}
`

func formatModuleName(name string) string {
	return "TEST"
}
