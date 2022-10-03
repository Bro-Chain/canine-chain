// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgInitAll } from "./types/filetree/tx";
import { MsgDeleteFile } from "./types/filetree/tx";
import { MsgAddViewers } from "./types/filetree/tx";
import { MsgInitAccount } from "./types/filetree/tx";
import { MsgPostFile } from "./types/filetree/tx";
import { MsgPostkey } from "./types/filetree/tx";


export { MsgInitAll, MsgDeleteFile, MsgAddViewers, MsgInitAccount, MsgPostFile, MsgPostkey };

type sendMsgInitAllParams = {
  value: MsgInitAll,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeleteFileParams = {
  value: MsgDeleteFile,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddViewersParams = {
  value: MsgAddViewers,
  fee?: StdFee,
  memo?: string
};

type sendMsgInitAccountParams = {
  value: MsgInitAccount,
  fee?: StdFee,
  memo?: string
};

type sendMsgPostFileParams = {
  value: MsgPostFile,
  fee?: StdFee,
  memo?: string
};

type sendMsgPostkeyParams = {
  value: MsgPostkey,
  fee?: StdFee,
  memo?: string
};


type msgInitAllParams = {
  value: MsgInitAll,
};

type msgDeleteFileParams = {
  value: MsgDeleteFile,
};

type msgAddViewersParams = {
  value: MsgAddViewers,
};

type msgInitAccountParams = {
  value: MsgInitAccount,
};

type msgPostFileParams = {
  value: MsgPostFile,
};

type msgPostkeyParams = {
  value: MsgPostkey,
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
		
		async sendMsgInitAll({ value, fee, memo }: sendMsgInitAllParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgInitAll: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgInitAll({ value: MsgInitAll.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgInitAll: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeleteFile({ value, fee, memo }: sendMsgDeleteFileParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteFile: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteFile({ value: MsgDeleteFile.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteFile: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddViewers({ value, fee, memo }: sendMsgAddViewersParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddViewers: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddViewers({ value: MsgAddViewers.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddViewers: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgInitAccount({ value, fee, memo }: sendMsgInitAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgInitAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgInitAccount({ value: MsgInitAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgInitAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgPostFile({ value, fee, memo }: sendMsgPostFileParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgPostFile: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgPostFile({ value: MsgPostFile.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgPostFile: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgPostkey({ value, fee, memo }: sendMsgPostkeyParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgPostkey: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgPostkey({ value: MsgPostkey.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgPostkey: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgInitAll({ value }: msgInitAllParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.filetree.MsgInitAll", value: MsgInitAll.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgInitAll: Could not create message: ' + e.message)
			}
		},
		
		msgDeleteFile({ value }: msgDeleteFileParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.filetree.MsgDeleteFile", value: MsgDeleteFile.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteFile: Could not create message: ' + e.message)
			}
		},
		
		msgAddViewers({ value }: msgAddViewersParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.filetree.MsgAddViewers", value: MsgAddViewers.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddViewers: Could not create message: ' + e.message)
			}
		},
		
		msgInitAccount({ value }: msgInitAccountParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.filetree.MsgInitAccount", value: MsgInitAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgInitAccount: Could not create message: ' + e.message)
			}
		},
		
		msgPostFile({ value }: msgPostFileParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.filetree.MsgPostFile", value: MsgPostFile.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgPostFile: Could not create message: ' + e.message)
			}
		},
		
		msgPostkey({ value }: msgPostkeyParams): EncodeObject {
			try {
				return { typeUrl: "/jackaldao.canine.filetree.MsgPostkey", value: MsgPostkey.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgPostkey: Could not create message: ' + e.message)
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
			JackaldaoCanineFiletree: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;