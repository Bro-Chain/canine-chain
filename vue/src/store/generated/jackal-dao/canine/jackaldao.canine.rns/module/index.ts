// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgBid } from "./types/rns/tx";
import { MsgInit } from "./types/rns/tx";
import { MsgAcceptBid } from "./types/rns/tx";
import { MsgCancelBid } from "./types/rns/tx";
import { MsgBuy } from "./types/rns/tx";
import { MsgTransfer } from "./types/rns/tx";
import { MsgDelist } from "./types/rns/tx";
import { MsgRegister } from "./types/rns/tx";
import { MsgAddRecord } from "./types/rns/tx";
import { MsgList } from "./types/rns/tx";
import { MsgDelRecord } from "./types/rns/tx";


const types = [
  ["/jackaldao.canine.rns.MsgBid", MsgBid],
  ["/jackaldao.canine.rns.MsgInit", MsgInit],
  ["/jackaldao.canine.rns.MsgAcceptBid", MsgAcceptBid],
  ["/jackaldao.canine.rns.MsgCancelBid", MsgCancelBid],
  ["/jackaldao.canine.rns.MsgBuy", MsgBuy],
  ["/jackaldao.canine.rns.MsgTransfer", MsgTransfer],
  ["/jackaldao.canine.rns.MsgDelist", MsgDelist],
  ["/jackaldao.canine.rns.MsgRegister", MsgRegister],
  ["/jackaldao.canine.rns.MsgAddRecord", MsgAddRecord],
  ["/jackaldao.canine.rns.MsgList", MsgList],
  ["/jackaldao.canine.rns.MsgDelRecord", MsgDelRecord],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgBid: (data: MsgBid): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgBid", value: MsgBid.fromPartial( data ) }),
    msgInit: (data: MsgInit): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgInit", value: MsgInit.fromPartial( data ) }),
    msgAcceptBid: (data: MsgAcceptBid): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgAcceptBid", value: MsgAcceptBid.fromPartial( data ) }),
    msgCancelBid: (data: MsgCancelBid): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgCancelBid", value: MsgCancelBid.fromPartial( data ) }),
    msgBuy: (data: MsgBuy): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgBuy", value: MsgBuy.fromPartial( data ) }),
    msgTransfer: (data: MsgTransfer): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgTransfer", value: MsgTransfer.fromPartial( data ) }),
    msgDelist: (data: MsgDelist): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgDelist", value: MsgDelist.fromPartial( data ) }),
    msgRegister: (data: MsgRegister): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgRegister", value: MsgRegister.fromPartial( data ) }),
    msgAddRecord: (data: MsgAddRecord): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgAddRecord", value: MsgAddRecord.fromPartial( data ) }),
    msgList: (data: MsgList): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgList", value: MsgList.fromPartial( data ) }),
    msgDelRecord: (data: MsgDelRecord): EncodeObject => ({ typeUrl: "/jackaldao.canine.rns.MsgDelRecord", value: MsgDelRecord.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
