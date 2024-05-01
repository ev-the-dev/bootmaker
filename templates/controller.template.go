package templates

var ControllerTemplate = `
{{with $pMN := formatModuleName $.ModuleName}}
import { Controller } from "@nestjs/common"

import { {{$pMN}}ControllerAdapters } from "@{{$.ModuleName}}/adapters/controller.adapters"

@Controller("{{$.ModuleName}}")
export class {{$pMN}}Controller {
  public constructor(
    private readonly _adapter: {{$pMN}}ControllerAdapters 
  ) {}
}
{{end}}
`
