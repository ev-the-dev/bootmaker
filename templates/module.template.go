package templates

var ModuleTemplate = `{{with $pMN := formatModuleName $.ModuleName}}import { Module } from "@nestjs/common"

import { PrismaModule as DatabaseModule } from "@database/prisma.module"

import { {{$pMN}}ControllerAdapters } from "./adapters/controller.adapters"
import { {{$pMN}}QueueConsumerAdapters } from "./adapters/queue-consumer.adapters"
import { {{$pMN}}RepositoryAdapters } from "./adapters/repository.adapters"
import { {{$pMN}}ServiceAdapters } from "./adapters/service.adapters"
import { {{$pMN}}Controller } from "./controllers/{{$.ModuleName}}.controller"
import { {{$pMN}}QueueConsumers } from "./queues/consumers.queues"
import { {{$pMN}}Repository } from "./repository/{{$.ModuleName}}.repository"
import { {{$pMN}}Service } from "./services/{{$.ModuleName}}.service"


@Module({
  imports: [ DatabaseModule ],
  controllers: [ {{$pMN}}Controller ],
  providers: [
    {{$pMN}}ControllerAdapters,
    {{$pMN}}QueueConsumers,
    {{$pMN}}QueueConsumerAdapters,
    {{$pMN}}Repository,
    {{$pMN}}RepositoryAdapters,
    {{$pMN}}Service,
    {{$pMN}}ServiceAdapters,
  ]
})
export class {{$pMN}}Module { }
{{end}}`
