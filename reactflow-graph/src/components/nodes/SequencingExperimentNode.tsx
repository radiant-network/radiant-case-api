import { Dna } from 'lucide-react';
import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import type {SequencingExperiment} from "../../types.ts";
import {
    NodeTooltip,
    NodeTooltipContent,
    NodeTooltipTrigger,
} from "@/components/node-tooltip";

type SequencingExperimentNode = Node<SequencingExperiment, 'sequencingExperiment'>;

export default function SequencingExperimentNode({data} :  NodeProps<SequencingExperimentNode>) {
    return (
        <NodeTooltip>
            <NodeTooltipContent position={Position.Top} className="text-gray-800 px-4 py-2 shadow-md rounded-md bg-white border-1 border-stone-400">
                <div className="text-sm font-bold">{data.aliquotId}</div>

                <div className="text-gray-600 text-xs">
                    Platform {data.platform}
                </div>
                <div className="text-gray-600 text-xs">
                    {data.sequencingReadTechnology}
                </div>
                <div className="absolute bottom-[-6px] left-1/2 -translate-x-1/2 w-3 h-3 bg-white border-b border-r border-stone-400 rotate-45"></div>
            </NodeTooltipContent>
            <NodeTooltipTrigger className="flex items-center px-4 py-2 shadow-sm rounded-md bg-white border border-slate-300">
                <Dna className="ml-1 w-10 h-10 text-amber-700" />
                <div className="ml-2">
                    <div className="text-lg font-bold">{data.experimentalStrategy}</div>
                </div>
            </NodeTooltipTrigger>

            <Handle type="target" position={Position.Left} className="w-2 !bg-pink-500"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-pink-500"/>
        </NodeTooltip>
    );
}