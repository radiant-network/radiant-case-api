import { Separator } from "@/components/ui/separator";
import { Badge } from "@/components/ui/badge";
import {Venus, Mars} from 'lucide-react';
import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import type {Patient} from "../../types.ts";

const iconMap = {
    F: <Venus className="w-7 h-7 text-pink-400" />,
    M: <Mars className="w-7 h-7 text-blue-400"/>
};
type PatientNode = Node<Patient, 'patient'>;

export default function PatientNode({data} :  NodeProps<PatientNode>) {
    const icon = iconMap[data.sex]
    return (
        <div className="px-4 py-2 shadow-sm rounded-md bg-white border">
            <div className="flex items-center">
                <div className="rounded-md bg-muted p-2">
                    {icon}
                </div>
                <div className="flex flex-col ml-4 gap-y-1">
                    <div className="flex">
                        <div className="text-lg font-bold">{data.orgPatientId}</div>
                        {/* <Separator orientation="vertical" className="mx-2" /> */}
                        {/* <div className="text-lg font-bold">{data.relationToProband}</div> */}
                    </div>
                    <div className="flex gap-x-2 items-baseline">
                        {/* <Badge variant="secondary">{data.relationToProband}</Badge> */}
                        <div className="text-sm font-bold">{data.relationToProband}</div>
                        <Badge variant="outline">{data.affectedStatus}</Badge>
                    </div>
                    {/* <div className="text-lg text-muted-foreground">{data.affectedStatus}</div> */}
                </div>
            </div>

            <Handle type="target" position={Position.Left} className="w-2 !bg-pink-500"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-pink-500"/>
        </div>
    );
}