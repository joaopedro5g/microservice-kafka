import { Controller, Get, Inject, OnModuleInit, Query } from '@nestjs/common';
import { ClientKafka, MessagePattern } from '@nestjs/microservices';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { AppService } from './app.service';

@Controller()
export class AppController implements OnModuleInit {
  private kafkaProducer: Producer;
  constructor(
    @Inject('KAFKA_SERVICE')
    private kafkaClient: ClientKafka,
    private appService: AppService,
  ) {}

  async onModuleInit() {
    this.kafkaProducer = await this.kafkaClient.connect();
  }

  @Get()
  getHello() {
    this.kafkaProducer.send({
      topic: 'test',
      messages: [
        {
          value: this.appService.sendData({
            id: '1',
            description: 'Descrição',
            name: 'produto',
            price: 32.1,
          }),
        },
      ],
    });
    return `Produto comprado com sucesso`;
  }
}
