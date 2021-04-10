import { Injectable } from '@nestjs/common';

interface SendDataTypes {
  id: string;
  name: string;
  description: string;
  price: number;
}

@Injectable()
export class AppService {
  sendData(data: SendDataTypes): string {
    return JSON.stringify(data);
  }
}
