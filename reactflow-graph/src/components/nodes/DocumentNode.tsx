import {Database, FileText, Settings} from 'lucide-react';
// import { ReactComponent as Vcf } from "../../assets/vcf.svg";
import CramIcon from '../../assets/cram.svg?react';
import VcfIcon from '../../assets/vcf.svg?react';
import GVcfIcon from '../../assets/gvcf.svg?react';
import JsonIcon from '../../assets/json.svg?react';
import TsvIcon from '../../assets/tsv.svg?react';

import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import {type Document} from '../../types';

const iconMap = {
    json: <JsonIcon className="w-10 h-10"/>,
    cram: <CramIcon className="w-10 h-10"/>,
    vcf: <VcfIcon className="w-10 h-10"/>,
    gvcf: <GVcfIcon className="w-10 h-10"/>,
    tsv: <TsvIcon className="w-10 h-10"/>,
};
import {
    NodeTooltip,
    NodeTooltipContent,
    NodeTooltipTrigger,
} from "@/components/node-tooltip";

type DocumentNode = Node<Document, 'document'>;

export default function DocumentNode({data}: NodeProps<DocumentNode>) {
    const icon = iconMap[data.type] ?? (
        <FileText className="w-10 h-10 text-gray-500"/>
    );
    return (
        <NodeTooltip>
            <NodeTooltipContent position={Position.Top} className="text-gray-800 px-4 py-2 shadow-md rounded-md bg-white border-1 border-stone-400">
                <div className="text-sm font-bold">{data.name}</div>
                {data.size && (
                    <div className="text-gray-600 text-xs">
                        {formatFileSize(data.size)}
                    </div>
                )}
            </NodeTooltipContent>
            <NodeTooltipTrigger className="px-4 py-2 shadow-md rounded-md bg-white border-2 border-stone-400">
                <div className="flex items-center">
                    {icon}
                </div>
                {/* Optional handles for edges */}
                <Handle type="target" position={Position.Left} className="w-2 !bg-teal-500"/>
                <Handle type="source" position={Position.Right} className="w-2 !bg-teal-500"/>
            </NodeTooltipTrigger>
        </NodeTooltip>
    );
}

function formatFileSize(bytes: number) {
    if (bytes < 1024) return `${bytes} B`;
    const kb = bytes / 1024;
    if (kb < 1024) return `${kb.toFixed(1)} KB`;
    const mb = kb / 1024;
    if (mb < 1024) return `${mb.toFixed(1)} MB`;
    return `${(mb / 1024).toFixed(1)} GB`;
}