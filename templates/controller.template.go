package templates

var ControllerTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Controller, Get, HttpCode, HttpStatus, Post, Req } from "@nestjs/common"
import { context, trace } from "@opentelemetry/api"

import { {{$pMN}}ControllerAdapters } from "@{{$.ModuleName}}/adapters/controller.adapters"
import { Create{{$pMN}}Validator, Get{{$pMN}}ByIdValidator } from "@{{$.ModuleName}}/decorators/controller.validators"
import { {{$pMN}}Service } from "@{{$.ModuleName}}/services/{{$.ModuleName}}.service"

import type { ICreate{{$pMN}}ClientDTO, I{{$pMN}}ClientDTO } from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"
import type { Tracer } from "@opentelemetry/api"


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
  @HttpCode(HttpStatus.CREATED)
  @Create{{$pMN}}Validator()
  public async create(@Req() req: ControllerRequestType<ICreate{{$pMN}}ClientDTO>): ControllerResponseType<I{{$pMN}}ClientDTO> {
    const span = this._tracer.startSpan("create", {}, context.active())

    const entityToCreate = this._adapters.create.clientToService(req.body)
    const createdEntity = await this._service.create(conext.active(), entityToCreate)
    const entityToReturn = this._adapters.base.serviceToClient(createdEntity)

    span.end()
    return {
      data: entityToReturn
    }
  }

  @Get(":id")
  @HttpCode(HttpStatus.OK)
  @Get{{$pMN}}ByIdValidator()
  public async getById(@Req() req: ControllerRequestType<void>): ControllerResponseType<I{{$pMN}}ClientDTO> {
    const span = this._tracer.startSpan("getById", {}, context.active())

    const fetchedEntity = await this._service.getById(context.active(), req.params.id)
    const entityToReturn = this._adapters.base.serviceToClient(fetchedEntity)

    span.end()
    return {
      data: entityToReturn
    }
  }
}
{{end}}
`

var ControllerSpecTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Test } from "@nestjs/testing"

import { {{$pMN}}ControllerAdapters } from "@{{$.ModuleName}}/adapters/controller.adapters"
import { {{$pMN}}Controller } from "@{{$.ModuleName}}/controllers/{{$.ModuleName}}.controller"
import { {{$pMN}}Service } from "@{{$.ModuleName}}/services/{{$.ModuleName}}.service"

import type { TestingModule } from "@nestjs/testing"

jest.mock("@{{$.ModuleName}}/services/{{$.ModuleName}}.service")
describe("{{$pMN}}Controller", () => {
  let controller: {{$pMN}}Controller

  beforeAll(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [{{$pMN}}Controller],
      providers: [
        {{$pMN}}ControllerAdapters,
        {{$pMN}}Service,
      ]
    }).compile()

    controller = module.get({{$pMN}}Controller)
  })

  it("should be defined", () => {
    expect(controller).toBeDefined()
  })
})
{{end}}
`
