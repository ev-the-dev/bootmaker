package templates

var ControllerAdapterTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Injectable } from "@nestjs/common"

import type {
  ICreate{{$pMN}}ClientDTO,
  ICreate{{$pMN}}ServiceDTO,
  I{{$pMN}}ClientDTO,
  I{{$pMN}}ServiceDTO
} from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"


@Injectable()
export class {{$pMN}}ControllerAdapters {
  public readonly base = {
  serviceToClient: (serviceEntity: I{{$pMN}}ServiceDTO): I{{$pMN}}ClientDTO => {
      return {}
    }
  }

  public readonly create = {
    clientToService: (clientEntity: ICreate{{$pMN}}ClientDTO): ICreate{{$pMN}}ServiceDTO => {
      return {}
    }
  }
}
{{end}}`

var QueueConsumerAdapterTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Injectable } from "@nestjs/common"

import type {
  ICreate{{$pMN}}MessageDTO,
  ICreate{{$pMN}}ServiceDTO,
} from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"


@Injectable()
export class {{$pMN}}QueueConsumerAdapters {
  public readonly create = {
    messageToService: (messageEntity: ICreate{{$pMN}}MessageDTO): ICreate{{$pMN}}ServiceDTO => {
      return {}
    }
  }
}
{{end}}`

var RepositoryAdapterTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Injectable } from "@nestjs/common"
import { $Enums, Prisma, {{$pMN}} as Prisma{{$pMN}} } from "@prisma/client"

import type {
  ICreate{{$pMN}}DataDTO,
  I{{$pMN}}DataDTO
} from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"


@Injectable()
export class {{$pMN}}RepositoryAdapters {
  public readonly base = {
    prismaToData: (ormEntity: Prisma{{$pMN}}): I{{$pMN}}DataDTO => {
      return {}
    }
  }

  public readonly create = {
    dataToPrisma: (dataEntity: ICreate{{$pMN}}DataDTO): Prisma.{{$pMN}}CreateInput => {
      return {}
    }
  }
}
{{end}}`

var ServiceAdapterTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Injectable } from "@nestjs/common"

import type {
  ICreate{{$pMN}}DataDTO,
  ICreate{{$pMN}}ServiceDTO,
  I{{$pMN}}DataDTO,
  I{{$pMN}}ServiceDTO
} from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"


@Injectable()
export class {{$pMN}}ServiceAdapters {
  public readonly base = {
    dataToService: (dataEntity: I{{$pMN}}DataDTO): I{{$pMN}}ServiceDTO => {
      return {}
    }
  }

  public readonly create = {
    serviceToData: (serviceEntity: ICreate{{$pMN}}ServiceDTO): ICreate{{$pMN}}DataDTO => {
      return {}
    }
  }
}
{{end}}`
