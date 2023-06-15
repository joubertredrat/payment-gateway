import { ServerUnaryCall, UntypedHandleCall, sendUnaryData } from '@grpc/grpc-js';

import { FakeCheck, FakeResponse } from './FakeCheck';
import { AuthorizationServer, AuthorizationService, CreditCardRequest, CreditCardResponse } from '../models/authorization';
import { logger } from '../utils';

class Authorization implements AuthorizationServer {
  [method: string]: UntypedHandleCall;

  public authorize(
    call: ServerUnaryCall<CreditCardRequest, CreditCardResponse>,
    callback: sendUnaryData<CreditCardResponse>,
  ): void {
    logger.info('authorize', Date.now());

    const { cardNumber, cvv, expireDate } = call.request;
    const fake: FakeCheck = new FakeCheck();
    const fakeResponse: FakeResponse = fake.handle(cardNumber, cvv, expireDate);

    const res: Partial<CreditCardResponse> = {};
    res.cardNumber = cardNumber;
    res.status = fakeResponse.code;
    res.message = fakeResponse.message;

    callback(null, CreditCardResponse.fromJSON(res));
  }
}

export {
  Authorization,
  AuthorizationService,
};
