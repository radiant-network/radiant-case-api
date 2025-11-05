import { User } from 'lucide-react';
import { Badge } from "@/components/ui/badge";
import {Venus, Mars} from 'lucide-react';
import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import type {Patient} from "../../types.ts";

const iconMap = {
    F: <Venus className="w-4 h-4" />,
    M: <Mars className="w-4 h-4"/>
};
type PatientNode = Node<Patient, 'patient'>;

export default function PatientNode({data} :  NodeProps<PatientNode>) {
    const icon = iconMap[data.sex]
    return (
        <div className="px-4 py-2 shadow-sm rounded-md bg-white border">
            <div className="flex items-center">
                <div className="rounded-md bg-sky-50 text-sky-800 p-2">
                    <User className="w-7 h-7" />
                </div>
                <div className="flex flex-col ml-4 gap-y-1">
                    <div className="text-lg font-bold">{data.orgPatientId}</div>
                    <div className="flex gap-x-2 items-center">
                        <div className="text-sm font-bold">{data.relationToProband}</div>
                        <div className="">
                            {icon}
                        </div>
                        <Badge variant="outline">{data.affectedStatus}</Badge>
                    </div>
                </div>
            </div>

            <Handle type="target" position={Position.Left} className="w-2 !bg-pink-500"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-pink-500"/>
        </div>
    );
}