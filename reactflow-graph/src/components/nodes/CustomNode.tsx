import { Handle, Position} from '@xyflow/react';

// interface CustomNodeData {
//     name: string;
//     job: string;
//     emoji: string;
// }
export default function CustomNode({data}: any) {
    return (
        <div className="px-4 py-2 shadow-sm rounded-md bg-white border">
            <div className="flex items-center">
                <div className="rounded-full w-12 h-12 flex justify-center items-center bg-gray-100">
                    {data.emoji}
                </div>
                <div className="ml-2">
                    <div className="text-lg font-bold">{data.name}</div>
                    <div className="text-gray-500">{data.job}</div>
                </div>
            </div>

            <Handle
                type="target"
                position={Position.Top}
                className="w-16 !bg-pink-500"
            />
            <Handle
                type="source"
                position={Position.Bottom}
                className="w-16 !bg-pink-500"
            />
        </div>
    );
};