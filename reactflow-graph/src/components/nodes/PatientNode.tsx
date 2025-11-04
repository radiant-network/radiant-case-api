import {Venus, Mars} from 'lucide-react';
import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import type {Patient} from "../../types.ts";

const iconMap = {
    F: <Venus className="w-10 h-10 text-amber-700" />,
    M: <Mars className="w-10 h-10 text-amber-700" />
};
type PatientNode = Node<Patient, 'patient'>;

export default function PatientNode({data} :  NodeProps<PatientNode>) {
    const icon = iconMap[data.sex]
    return (
        <div className="px-4 py-2 shadow-sm rounded-md bg-white border border-slate-300">
            <div className="flex items-center">
                {icon}
                <div className="ml-2">
                    <div className="text-lg font-bold">{data.orgPatientId}</div>
                    <div className="text-lg">{data.relationToProband}</div>
                    <div className="text-lg">{data.affectedStatus}</div>
                </div>
            </div>

            <Handle type="target" position={Position.Left} className="w-2 !bg-pink-500"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-pink-500"/>
        </div>
    );
}