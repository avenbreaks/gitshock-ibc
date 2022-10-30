import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePool } from "./types/neware/dex/tx";
import { MsgUpdatePool } from "./types/neware/dex/tx";
import { MsgDeletePool } from "./types/neware/dex/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/neware.dex.MsgCreatePool", MsgCreatePool],
    ["/neware.dex.MsgUpdatePool", MsgUpdatePool],
    ["/neware.dex.MsgDeletePool", MsgDeletePool],
    
];

export { msgTypes }