import { Workflow } from 'lucide-react'; // pipeline-like icon
import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import type {Task} from "../../types.ts";

type TaskNode = Node<Task, 'task'>;

export default function TaskNode({data} :  NodeProps<TaskNode>) {
    return (
        <div className="px-4 py-2 shadow-sm rounded-md bg-white border">
            <div className="flex items-center">
                <div className="rounded-md bg-muted p-2">
                    <Workflow className="w-7 h-7 text-amber-600" />
                </div>
                <div className="ml-4">
                    <div className="text-lg font-bold">{data.name}</div>
                </div>
            </div>

            <Handle type="target" position={Position.Left} className="w-2 !bg-pink-500" id="left"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-pink-500"/>
        </div>
    );
}