import { TestTube } from 'lucide-react';
import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import type {Sample} from "../../types.ts";

type SampleNode = Node<Sample, 'sample'>;

export default function SampleNode({data} :  NodeProps<SampleNode>) {
    return (
        <div className="px-4 py-2 shadow-md rounded-md bg-white border-2 border-stone-400">
            <div className="flex items-center">
            <TestTube className="w-10 h-10 text-amber-700" />
                <div className="ml-2">
                    <div className="text-lg font-bold">{data.orgSampleId}</div>
                </div>
            </div>

            <Handle type="target" position={Position.Left} className="w-2 !bg-teal-500"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-teal-500"/>
        </div>
    );
}

