package templates

var ServiceTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Injectable } from "@nestjs/common"
import { context, trace } from "@opentelemetry/api"

import { {{$pMN}}ServiceAdapters } from "@{{$.ModuleName}}/adapters/service.adapters"
import { {{$pMN}}Repository } from "@{{$.ModuleName}}/repository/{{$.ModuleName}}.repository"

import type { ICreate{{$pMN}}ServiceDTO, I{{$pMN}}ServiceDTO } from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"
import type { Context, Tracer } from "@opentelemetry/api"


@Injectable()
export class {{$pMN}}Service {
  private readonly _tracer: Tracer
  public constructor(
    private readonly _adapters: {{$pMN}}ServiceAdapters,
    private readonly _repository: {{$pMN}}Repository
  ) {
    this._tracer = trace.getTracer("{{$.ModuleName}}-service")
  }

  public async create(ctx: Context, createEntity: ICreate{{$pMN}}ServiceDTO): Promise<I{{$pMN}}ServiceDTO> {
    const span = this._tracer.startSpan("create", {}, ctx)

    const entityToCreate = this._adapters.create.serviceToData(createEntity)
    const createdEntity = await this._repository.create(context.active(), entityToCreate)
    const entityToReturn = this._adapters.base.dataToService(createdEntity)

    span.end()
    return entityToReturn
  }

  public async getById(ctx: Context, id: string): Promise<I{{$pMN}}ServiceDTO> {
    const span = this._tracer.startSpan("create", {}, ctx)

    const fetchedEntity = await this._repository.getById(context.active(), id)
    const entityToReturn = this._adapters.base.dataToService(fetchedEntity)

    span.end()
    return entityToReturn
  }
}
{{end}}
`

var ServiceSpecTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Test } from "@nestjs/testing"

import { {{$pMN}}ServiceAdapters } from "@{{$.ModuleName}}/adapters/service.adapters"
import { {{$pMN}}Repository } from "@{{$.ModuleName}}/repository/{{$.ModuleName}}.repository"
import { {{$pMN}}Service } from "@{{$.ModuleName}}/services/{{$.ModuleName}}.service"

import type { TestingModule } from "@nestjs/testing"

jest.mock("@{{$.ModuleName}}/repository/{{$.ModuleName}}.repository")
describe("{{$pMN}}Service", () => {
  let service: {{$pMN}}Service

  beforeAll(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        {{$pMN}}Repository,
        {{$pMN}}Service,
        {{$pMN}}ServiceAdapters,
      ]
    }).compile()

    service = module.get({{$pMN}}Service)
  })

  it("should be defined", () => {
    expect(service).toBeDefined()
  })
})
{{end}}
`
