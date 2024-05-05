package templates

var RepositoryTemplate = `
{{with $pMN := formatModuleName $.ModuleName}}
import { Injectable, NotImplementedException } from "@nestjs/common"
import { trace } from "@opentelemetry/api"

import { PrismaService } from "@database/services/prisma.service"
import { DatabaseErrorGenerator } from "@database/utils/DatabaseErrorGenerator"
import { {{$pMN}}RepositoryAdapters } from "@{{$.ModuleName}}/adapters/service.adapters"

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
    this._tracer = trace.getTracer("{{$.ModuleName}}-service")
  }

  public async create(createEntity: ICreate{{$pMN}}DataDTO): Promise<I{{$pMN}}DataDTO> {
    throw new NotImplementedException()
  }

  public async getById(id: string): Promise<I{{$pMN}}DataDTO> {
    throw new NotImplementedException()
  }
}
{{end}}
`
