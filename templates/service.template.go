package templates

var ServiceTemplate = `
{{with $pMN := formatModuleName $.ModuleName}}
import { Injectable, NotImplementedException } from "@nestjs/common"
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

  public async create(createEntity: ICreate{{$pMN}}ServiceDTO): Promise<I{{$pMN}}ServiceDTO> {
    throw new NotImplementedException()
  }

  public async getById(id: string): Promise<I{{$pMN}}ServiceDTO> {
    throw new NotImplementedException()
  }
}
{{end}}
`
