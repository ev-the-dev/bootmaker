package templates

var RepositoryTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Injectable } from "@nestjs/common"
import { SpanStatusCode, trace } from "@opentelemetry/api"

import { PrismaService } from "@database/services/prisma.service"
import { DatabaseErrorGenerator } from "@database/utils/DatabaseErrorGenerator"
import { {{$pMN}}RepositoryAdapters } from "@{{$.ModuleName}}/adapters/repository.adapters"

import type { ICreate{{$pMN}}DataDTO, I{{$pMN}}DataDTO } from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"
import type { Context, Tracer } from "@opentelemetry/api"


@Injectable()
export class {{$pMN}}Repository {
  private readonly _tracer: Tracer
  public constructor(
    private readonly _adapters: {{$pMN}}RepositoryAdapters,
    private readonly _databaseErrorGenerator: DatabaseErrorGenerator,
    private readonly _orm: PrismaService
  ) {
    this._tracer = trace.getTracer("{{$.ModuleName}}-repository")
  }

  public async create(ctx: Context, createEntity: ICreate{{$pMN}}DataDTO): Promise<I{{$pMN}}DataDTO> {
    const span = this._tracer.startSpan("create", {}, ctx)
    try {
      const entityToInsert = this._adapters.create.dataToPrisma(createEntity)
      // TODO: ORM call
      // adapt from orm -> data
      // return data DTO
    } catch(e) {
      const error = e as Error
      span.recordException(error)
      span.setStatus({ code: SpanStatusCode.ERROR })
      throw this._databaseErrorGenerator.generateError(error, { location: "{{$pMN}}Repository.create" })
    } finally {
      span.end()
    }
  }

  public async getById(ctx: Context, id: string): Promise<I{{$pMN}}DataDTO> {
    const span = this._tracer.startSpan("getById", {}, ctx)
    try { 
      // TODO: ORM call
      // adapt from orm -> data
      // return data DTO
    } catch(e) {
      const error = e as Error
      span.recordException(error)
      span.setStatus({ code: SpanStatusCode.ERROR })
      throw this._databaseErrorGenerator.generateError(error, { location: "{{$pMN}}Repository.getById" })
    } finally {
      span.end()
    }
  }
}
{{end}}
`

var RepositorySpecTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Test } from "@nestjs/testing"

import { PrismaService } from "@database/services/prisma.service"
import { {{$pMN}}Repository } from "@{{$.ModuleName}}/repository/{{$.ModuleName}}.repository"
import { {{$pMN}}RepositoryAdapters } from "@{{$.ModuleName}}/adapters/repository.adapters"

import type { TestingModule } from "@nestjs/testing"

jest.mock("@database/services/prisma.service")
describe("{{$pMN}}Repository", () => {
  let repository: {{$pMN}}Repository

  beforeAll(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        {{$pMN}}Repository,
        {{$pMN}}Repository,
        {{$pMN}}RepositoryAdapters,
        PrismaService
      ]
    }).compile()

    repository = module.get({{$pMN}}Repository)
  })

  it("should be defined", () => {
    expect(repository).toBeDefined()
  })
})
{{end}}
`
