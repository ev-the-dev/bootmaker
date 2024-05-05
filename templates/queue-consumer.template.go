package templates

var QueueConsumerTemplate = `
{{with $pMN := formatModuleName $.ModuleName}}
import { Injectable, NotImplementedException } from "@nestjs/common"
import { context, trace } from "@opentelemetry/api"

import { {{$pMN}}QueueConsumerAdapters } from "@{{$.ModuleName}}/adapters/queue-consumer.adapters"
import { {{$pMN}}Service } from "@{{$.ModuleName}}/services/{{$.ModuleName}}.service"

import type { ICreate{{$pMN}}QueueDTO } from "@{{$.ModuleName}}/dtos/{{$.ModuleName}}.dtos"
import type { Tracer } from "@opentelemetry/api"
import type { IMessage } from "@queues/types"


@Injectable()
export class {{$pMN}}QueueConsumers {
  private readonly _tracer: Tracer
  public constructor(
    private readonly _adapters: {{$pMN}}QueueConsumerAdapters,
    private readonly _service: {{$pMN}}Service
  ) {
    this._tracer = trace.getTracer("{{$.ModuleName}}-queue-consumer")
  }

  {{with $eMN := formatModuleNameEnum $.ModuleName}}
  @SqsMessageHandler({{$eMN}}_QUEUE_CONSUMER_METADATA.CREATE)
  {{end}}
  public async create(message: IMessage<ICreate{{$pMN}}QueueDTO>): Promise<void> {
    throw new NotImplementedException()
  }
}
{{end}}
`
