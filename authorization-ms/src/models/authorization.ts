/* eslint-disable */
import {
  CallOptions,
  ChannelCredentials,
  Client,
  ClientOptions,
  ClientUnaryCall,
  handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  ServiceError,
  UntypedServiceImplementation,
} from "@grpc/grpc-js";
import Long from "long";
import _m0 from "protobufjs/minimal";

export enum AuthorizationStatus {
  AUTHORIZED = 0,
  DECLINED_UNKNOWN = 1,
  DECLINED_INVALID_DATA = 2,
  DECLINED_EXPIRED_CARD = 3,
  DECLINED_INSUFFICIENT_FUNDS = 4,
  DECLINED_SUSPECT_FRAUD = 5,
  UNRECOGNIZED = -1,
}

export function authorizationStatusFromJSON(object: any): AuthorizationStatus {
  switch (object) {
    case 0:
    case "AUTHORIZED":
      return AuthorizationStatus.AUTHORIZED;
    case 1:
    case "DECLINED_UNKNOWN":
      return AuthorizationStatus.DECLINED_UNKNOWN;
    case 2:
    case "DECLINED_INVALID_DATA":
      return AuthorizationStatus.DECLINED_INVALID_DATA;
    case 3:
    case "DECLINED_EXPIRED_CARD":
      return AuthorizationStatus.DECLINED_EXPIRED_CARD;
    case 4:
    case "DECLINED_INSUFFICIENT_FUNDS":
      return AuthorizationStatus.DECLINED_INSUFFICIENT_FUNDS;
    case 5:
    case "DECLINED_SUSPECT_FRAUD":
      return AuthorizationStatus.DECLINED_SUSPECT_FRAUD;
    case -1:
    case "UNRECOGNIZED":
    default:
      return AuthorizationStatus.UNRECOGNIZED;
  }
}

export function authorizationStatusToJSON(object: AuthorizationStatus): string {
  switch (object) {
    case AuthorizationStatus.AUTHORIZED:
      return "AUTHORIZED";
    case AuthorizationStatus.DECLINED_UNKNOWN:
      return "DECLINED_UNKNOWN";
    case AuthorizationStatus.DECLINED_INVALID_DATA:
      return "DECLINED_INVALID_DATA";
    case AuthorizationStatus.DECLINED_EXPIRED_CARD:
      return "DECLINED_EXPIRED_CARD";
    case AuthorizationStatus.DECLINED_INSUFFICIENT_FUNDS:
      return "DECLINED_INSUFFICIENT_FUNDS";
    case AuthorizationStatus.DECLINED_SUSPECT_FRAUD:
      return "DECLINED_SUSPECT_FRAUD";
    case AuthorizationStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface CreditCardRequest {
  holderName: string;
  cardNumber: string;
  cvv: string;
  expireDate: number;
  amount: number;
  installments: number;
}

export interface CreditCardResponse {
  cardNumber: string;
  status: AuthorizationStatus;
  message: string;
}

function createBaseCreditCardRequest(): CreditCardRequest {
  return { holderName: "", cardNumber: "", cvv: "", expireDate: 0, amount: 0, installments: 0 };
}

export const CreditCardRequest = {
  encode(message: CreditCardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.holderName !== "") {
      writer.uint32(10).string(message.holderName);
    }
    if (message.cardNumber !== "") {
      writer.uint32(18).string(message.cardNumber);
    }
    if (message.cvv !== "") {
      writer.uint32(26).string(message.cvv);
    }
    if (message.expireDate !== 0) {
      writer.uint32(32).int64(message.expireDate);
    }
    if (message.amount !== 0) {
      writer.uint32(40).uint64(message.amount);
    }
    if (message.installments !== 0) {
      writer.uint32(48).uint32(message.installments);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreditCardRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreditCardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.holderName = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.cardNumber = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.cvv = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.expireDate = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.amount = longToNumber(reader.uint64() as Long);
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.installments = reader.uint32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreditCardRequest {
    return {
      holderName: isSet(object.holderName) ? String(object.holderName) : "",
      cardNumber: isSet(object.cardNumber) ? String(object.cardNumber) : "",
      cvv: isSet(object.cvv) ? String(object.cvv) : "",
      expireDate: isSet(object.expireDate) ? Number(object.expireDate) : 0,
      amount: isSet(object.amount) ? Number(object.amount) : 0,
      installments: isSet(object.installments) ? Number(object.installments) : 0,
    };
  },

  toJSON(message: CreditCardRequest): unknown {
    const obj: any = {};
    message.holderName !== undefined && (obj.holderName = message.holderName);
    message.cardNumber !== undefined && (obj.cardNumber = message.cardNumber);
    message.cvv !== undefined && (obj.cvv = message.cvv);
    message.expireDate !== undefined && (obj.expireDate = Math.round(message.expireDate));
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    message.installments !== undefined && (obj.installments = Math.round(message.installments));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreditCardRequest>, I>>(base?: I): CreditCardRequest {
    return CreditCardRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreditCardRequest>, I>>(object: I): CreditCardRequest {
    const message = createBaseCreditCardRequest();
    message.holderName = object.holderName ?? "";
    message.cardNumber = object.cardNumber ?? "";
    message.cvv = object.cvv ?? "";
    message.expireDate = object.expireDate ?? 0;
    message.amount = object.amount ?? 0;
    message.installments = object.installments ?? 0;
    return message;
  },
};

function createBaseCreditCardResponse(): CreditCardResponse {
  return { cardNumber: "", status: 0, message: "" };
}

export const CreditCardResponse = {
  encode(message: CreditCardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cardNumber !== "") {
      writer.uint32(10).string(message.cardNumber);
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status);
    }
    if (message.message !== "") {
      writer.uint32(26).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreditCardResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreditCardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.cardNumber = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreditCardResponse {
    return {
      cardNumber: isSet(object.cardNumber) ? String(object.cardNumber) : "",
      status: isSet(object.status) ? authorizationStatusFromJSON(object.status) : 0,
      message: isSet(object.message) ? String(object.message) : "",
    };
  },

  toJSON(message: CreditCardResponse): unknown {
    const obj: any = {};
    message.cardNumber !== undefined && (obj.cardNumber = message.cardNumber);
    message.status !== undefined && (obj.status = authorizationStatusToJSON(message.status));
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreditCardResponse>, I>>(base?: I): CreditCardResponse {
    return CreditCardResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreditCardResponse>, I>>(object: I): CreditCardResponse {
    const message = createBaseCreditCardResponse();
    message.cardNumber = object.cardNumber ?? "";
    message.status = object.status ?? 0;
    message.message = object.message ?? "";
    return message;
  },
};

export type AuthorizationService = typeof AuthorizationService;
export const AuthorizationService = {
  authorize: {
    path: "/service.Authorization/authorize",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreditCardRequest) => Buffer.from(CreditCardRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreditCardRequest.decode(value),
    responseSerialize: (value: CreditCardResponse) => Buffer.from(CreditCardResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => CreditCardResponse.decode(value),
  },
} as const;

export interface AuthorizationServer extends UntypedServiceImplementation {
  authorize: handleUnaryCall<CreditCardRequest, CreditCardResponse>;
}

export interface AuthorizationClient extends Client {
  authorize(
    request: CreditCardRequest,
    callback: (error: ServiceError | null, response: CreditCardResponse) => void,
  ): ClientUnaryCall;
  authorize(
    request: CreditCardRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: CreditCardResponse) => void,
  ): ClientUnaryCall;
  authorize(
    request: CreditCardRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: CreditCardResponse) => void,
  ): ClientUnaryCall;
}

export const AuthorizationClient = makeGenericClientConstructor(
  AuthorizationService,
  "service.Authorization",
) as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): AuthorizationClient;
  service: typeof AuthorizationService;
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
