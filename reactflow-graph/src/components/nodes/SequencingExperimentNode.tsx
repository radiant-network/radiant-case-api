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
            <NodeTooltipContent position={Position.Top} className="text-foreground px-4 py-2 shadow-md rounded-md bg-white border-1">
                <div className="text-sm font-bold">{data.aliquotId}</div>

                <div className="text-muted-foreground text-xs">
                    Platform {data.platform}
                </div>
                <div className="text-muted-foreground text-xs">
                    {data.sequencingReadTechnology}
                </div>
                <div className="absolute bottom-[-6px] left-1/2 -translate-x-1/2 w-3 h-3 bg-white border-b border-r rotate-45"></div>
            </NodeTooltipContent>
            <NodeTooltipTrigger className="flex items-center px-4 py-2 shadow-sm rounded-md bg-white border">
                <div className="rounded-md bg-muted p-2">
                    <Dna className="ml-1 w-7 h-7 text-violet-500" />
                </div>
                <div className="ml-4">
                    <div className="text-lg font-bold">{data.experimentalStrategy}</div>
                </div>
            </NodeTooltipTrigger>

            <Handle type="target" position={Position.Left} className="w-2 !bg-pink-500"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-pink-500"/>
        </NodeTooltip>
    );
}