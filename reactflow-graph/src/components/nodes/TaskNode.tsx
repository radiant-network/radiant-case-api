import { Workflow } from 'lucide-react'; // pipeline-like icon
import {Handle, type Node, type NodeProps, Position} from '@xyflow/react';
import type {Task} from "../../types.ts";

type TaskNode = Node<Task, 'task'>;

export default function TaskNode({data} :  NodeProps<TaskNode>) {
    return (
        <div className="px-4 py-2 shadow-md rounded-md bg-white border-2 border-stone-400">
            <div className="flex items-center">
            <Workflow className="w-10 h-10 text-amber-700" />
                <div className="ml-2">
                    <div className="text-lg font-bold">{data.name}</div>
                </div>
            </div>

            <Handle type="target" position={Position.Top} className="w-2 !bg-teal-500" id="top"/>
            <Handle type="target" position={Position.Left} className="w-2 !bg-teal-500" id="left"/>
            <Handle type="source" position={Position.Right} className="w-2 !bg-teal-500"/>
        </div>
    );
}