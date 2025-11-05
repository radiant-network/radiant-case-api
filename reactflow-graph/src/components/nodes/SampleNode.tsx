import { TestTube } from 'lucide-react';
import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import type {Sample} from "../../types.ts";

type SampleNode = Node<Sample, 'sample'>;

export default function SampleNode({data} :  NodeProps<SampleNode>) {
    return (
        <div className="px-4 py-2 shadow-sm rounded-md bg-white border">
            <div className="flex items-center">
                <div className="rounded-md bg-muted p-2">
                    <TestTube className="w-7 h-7 text-cyan-600" />
                </div>
                <div className="ml-4">
                    <div className="text-lg font-bold">{data.orgSampleId}</div>
                    <div className="flex gap-x-2 items-baseline">
                        <div className="text-sm font-bold">{data.histological}</div>
                    </div>
                </div>

            </div>

            <Handle type="target" position={Position.Left} className="w-2 !bg-pink-500"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-pink-500"/>
        </div>
    );
}

