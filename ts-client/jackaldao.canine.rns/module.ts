// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgDelRecord } from "./types/rns/tx";
import { MsgList } from "./types/rns/tx";
import { MsgTransfer } from "./types/rns/tx";
import { MsgCancelBid } from "./types/rns/tx";
import { MsgBuy } from "./types/rns/tx";
import { MsgAcceptBid } from "./types/rns/tx";
import { MsgDelist } from "./types/rns/tx";
import { MsgRegister } from "./types/rns/tx";
import { MsgInit } from "./types/rns/tx";
import { MsgAddRecord } from "./types/rns/tx";
import { MsgBid } from "./types/rns/tx";


export { MsgDelRecord, MsgList, MsgTransfer, MsgCancelBid, MsgBuy, MsgAcceptBid, MsgDelist, MsgRegister, MsgInit, MsgAddRecord, MsgBid };

type sendMsgDelRecordParams = {
  value: MsgDelRecord,
  fee?: StdFee,
  memo?: string
};

type sendMsgListParams = {
  value: MsgList,
  fee?: StdFee,
  memo?: string
};

type sendMsgTransferParams = {
  value: MsgTransfer,
  fee?: StdFee,
  memo?: string
};

type sendMsgCancelBidParams = {
  value: MsgCancelBid,
  fee?: StdFee,
  memo?: string
};

type sendMsgBuyParams = {
  value: MsgBuy,
  fee?: StdFee,
  memo?: string
};

type sendMsgAcceptBidParams = {
  value: MsgAcceptBid,
  fee?: StdFee,
  memo?: string
};

type sendMsgDelistParams = {
  value: MsgDelist,
  fee?: StdFee,
  memo?: string
};

type sendMsgRegisterParams = {
  value: MsgRegister,
  fee?: StdFee,
  memo?: string
};

type sendMsgInitParams = {
  value: MsgInit,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddRecordParams = {
  value: MsgAddRecord,
  fee?: StdFee,
  memo?: string
};

type sendMsgBidParams = {
  value: MsgBid,
  fee?: StdFee,
  memo?: string
};


type msgDelRecordParams = {
  value: MsgDelRecord,
};

type msgListParams = {
  value: MsgList,
};

type msgTransferParams = {
  value: MsgTransfer,
};

type msgCancelBidParams = {
  value: MsgCancelBid,
};

type msgBuyParams = {
  value: MsgBuy,
};

type msgAcceptBidParams = {
  value: MsgAcceptBid,
};

type msgDelistParams = {
  value: MsgDelist,
};

type msgRegisterParams = {
  value: MsgRegister,
};

type msgInitParams = {
  value: MsgInit,
};

type msgAddRecordParams = {
  value: MsgAddRecord,
};

type msgBidParams = {
  value: MsgBid,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgDelRecord({ value, fee, memo }: sendMsgDelRecordParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDelRecord: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDelRecord({ value: MsgDelRecord.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDelRecord: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgList({ value, fee, memo }: sendMsgListParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgList: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgList({ value: MsgList.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgList: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgTransfer({ value, fee, memo }: sendMsgTransferParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgTransfer: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgTransfer({ value: MsgTransfer.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgTransfer: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCancelBid({ value, fee, memo }: sendMsgCancelBidParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCancelBid: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCancelBid({ value: MsgCancelBid.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCancelBid: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgBuy({ value, fee, memo }: sendMsgBuyParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgBuy: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgBuy({ value: MsgBuy.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgBuy: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAcceptBid({ value, fee, memo }: sendMsgAcceptBidParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAcceptBid: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAcceptBid({ value: MsgAcceptBid.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAcceptBid: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDelist({ value, fee, memo }: sendMsgDelistParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDelist: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDelist({ value: MsgDelist.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDelist: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRegister({ value, fee, memo }: sendMsgRegisterParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRegister: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRegister({ value: MsgRegister.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRegister: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgInit({ value, fee, memo }: sendMsgInitParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgInit: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgInit({ value: MsgInit.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgInit: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddRecord({ value, fee, memo }: sendMsgAddRecordParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddRecord: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddRecord({ value: MsgAddRecord.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddRecord: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgBid({ value, fee, memo }: sendMsgBidParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgBid: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgBid({ value: MsgBid.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgBid: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgDelRecord({ value }: msgDelRecordParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgDelRecord", value: MsgDelRecord.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDelRecord: Could not create message: ' + e.message)
			}
		},
		
		msgList({ value }: msgListParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgList", value: MsgList.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgList: Could not create message: ' + e.message)
			}
		},
		
		msgTransfer({ value }: msgTransferParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgTransfer", value: MsgTransfer.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgTransfer: Could not create message: ' + e.message)
			}
		},
		
		msgCancelBid({ value }: msgCancelBidParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgCancelBid", value: MsgCancelBid.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCancelBid: Could not create message: ' + e.message)
			}
		},
		
		msgBuy({ value }: msgBuyParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgBuy", value: MsgBuy.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBuy: Could not create message: ' + e.message)
			}
		},
		
		msgAcceptBid({ value }: msgAcceptBidParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgAcceptBid", value: MsgAcceptBid.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAcceptBid: Could not create message: ' + e.message)
			}
		},
		
		msgDelist({ value }: msgDelistParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgDelist", value: MsgDelist.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDelist: Could not create message: ' + e.message)
			}
		},
		
		msgRegister({ value }: msgRegisterParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgRegister", value: MsgRegister.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRegister: Could not create message: ' + e.message)
			}
		},
		
		msgInit({ value }: msgInitParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgInit", value: MsgInit.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgInit: Could not create message: ' + e.message)
			}
		},
		
		msgAddRecord({ value }: msgAddRecordParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgAddRecord", value: MsgAddRecord.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddRecord: Could not create message: ' + e.message)
			}
		},
		
		msgBid({ value }: msgBidParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.rns.MsgBid", value: MsgBid.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBid: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]>;

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });
		this.tx = txClient({ signer: client.signer, addr: client.env.rpcURL, prefix: client.env.prefix ?? "cosmos" });
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			JackaldaoCanineRns: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;