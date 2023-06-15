import { AuthorizationStatus } from '../models/authorization';

export interface FakeResponse {
  code: number;
  message: string;
}

export class FakeCheck {
  cardNumberRegex: RegExp = /^[\d]{16}$/;

  cvvRegex: RegExp = /^[\d]{3}$/;

  public handle(cardNumber: string, cvv: string, expireDate: number): FakeResponse {
    const fieldErrors: string[] = [];

    if (!this.cardNumberRegex.test(cardNumber)) {
      fieldErrors.push('cardNumber');
    }
    if (!this.cvvRegex.test(cvv)) {
      fieldErrors.push('cvv');
    }
    if (expireDate < 1) {
      fieldErrors.push('expireDate');
    }
    if (fieldErrors.length > 0) {
      return {
        code: AuthorizationStatus.DECLINED_INVALID_DATA,
        message: `invalid fields ${fieldErrors.join(', ')}`,
      };
    }

    const date: Date = new Date(expireDate * 1000);
    const today: Date = new Date();
    if (today > date) {
      return {
        code: AuthorizationStatus.DECLINED_EXPIRED_CARD,
        message: 'card expired',
      };
    }

    if (cvv.startsWith('1')) {
      return {
        code: AuthorizationStatus.DECLINED_UNKNOWN,
        message: 'unknown error',
      };
    }

    if (cvv.startsWith('4')) {
      return {
        code: AuthorizationStatus.DECLINED_INSUFFICIENT_FUNDS,
        message: 'insufficient funds in card',
      };
    }

    if (cvv.startsWith('5')) {
      return {
        code: AuthorizationStatus.DECLINED_SUSPECT_FRAUD,
        message: 'this transaction was declined due a fraud suspection',
      };
    }

    return {
      code: AuthorizationStatus.AUTHORIZED,
      message: 'authorized',
    };
  }
}
