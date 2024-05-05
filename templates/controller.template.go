package templates

var ControllerTemplate = `
{{with $pMN := formatModuleName $.ModuleName}}
import { Controller, Get, HttpCode, HttpStatus, NotImplementedException, Post, Req } from "@nestjs/common"
import { context, trace } from "@opentelemetry/api"

import { {{$pMN}}ControllerAdapters } from "@{{$.ModuleName}}/adapters/controller.adapters"
import { {{$pMN}}Service } from "@{{$.ModuleName}}/services/{{$.ModuleName}}.service"

import type { ICreate{{$pMN}}ClientDTO, I{{$pMN}}ClientDTO } from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"
import type { Context, Tracer } from "@opentelemetry/api"


@Controller("{{$.ModuleName}}")
export class {{$pMN}}Controller {
  private readonly _tracer: Tracer
  public constructor(
    private readonly _adapters: {{$pMN}}ControllerAdapters,
    private readonly _service: {{$pMN}}Service
  ) {
    this._tracer = trace.getTracer("{{$.ModuleName}}-controller")
  }

  @Post()
  @HttpCode(@HttpStatus.CREATED)
  public async create(@Req() req: ControllerRequestType<ICreate{{$pMN}}ClientDTO>): ControllerResponseType<I{{$pMN}}ClientDTO> {
    throw new NotImplementedException()
  }

  @Get(":id")
  @HttpCode(@HttpStatus.OK)
  public async getById(@Req() req: ControllerRequestType<void>): ControllerResponseType<I{{$pMN}}ClientDTO> {
    throw new NotImplementedException()
  }
}
{{end}}
`
