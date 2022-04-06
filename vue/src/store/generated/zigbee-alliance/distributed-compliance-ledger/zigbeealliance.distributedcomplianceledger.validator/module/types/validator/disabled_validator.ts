/* eslint-disable */
import { Grant } from '../validator/grant'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface DisabledValidator {
  address: string
  creator: string
  approvals: Grant[]
  disabledByNodeAdmin: boolean
}

const baseDisabledValidator: object = { address: '', creator: '', disabledByNodeAdmin: false }

export const DisabledValidator = {
  encode(message: DisabledValidator, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    if (message.creator !== '') {
      writer.uint32(18).string(message.creator)
    }
    for (const v of message.approvals) {
      Grant.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    if (message.disabledByNodeAdmin === true) {
      writer.uint32(32).bool(message.disabledByNodeAdmin)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): DisabledValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseDisabledValidator } as DisabledValidator
    message.approvals = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        case 2:
          message.creator = reader.string()
          break
        case 3:
          message.approvals.push(Grant.decode(reader, reader.uint32()))
          break
        case 4:
          message.disabledByNodeAdmin = reader.bool()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): DisabledValidator {
    const message = { ...baseDisabledValidator } as DisabledValidator
    message.approvals = []
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(Grant.fromJSON(e))
      }
    }
    if (object.disabledByNodeAdmin !== undefined && object.disabledByNodeAdmin !== null) {
      message.disabledByNodeAdmin = Boolean(object.disabledByNodeAdmin)
    } else {
      message.disabledByNodeAdmin = false
    }
    return message
  },

  toJSON(message: DisabledValidator): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    message.creator !== undefined && (obj.creator = message.creator)
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => (e ? Grant.toJSON(e) : undefined))
    } else {
      obj.approvals = []
    }
    message.disabledByNodeAdmin !== undefined && (obj.disabledByNodeAdmin = message.disabledByNodeAdmin)
    return obj
  },

  fromPartial(object: DeepPartial<DisabledValidator>): DisabledValidator {
    const message = { ...baseDisabledValidator } as DisabledValidator
    message.approvals = []
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(Grant.fromPartial(e))
      }
    }
    if (object.disabledByNodeAdmin !== undefined && object.disabledByNodeAdmin !== null) {
      message.disabledByNodeAdmin = object.disabledByNodeAdmin
    } else {
      message.disabledByNodeAdmin = false
    }
    return message
  }
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
