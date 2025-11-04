import {useCallback} from 'react';
import {
    ReactFlow,
    useNodesState,
    useEdgesState,
    addEdge,
    MiniMap,
    Controls
} from '@xyflow/react';

import CustomNode from './nodes/CustomNode';

const nodeTypes = {
    custom: CustomNode,
};

const initialNodes = [
    {
        id: '1',
        type: 'custom',
        data: {name: 'Jane Doe', job: 'CEO', emoji: '😎'},
        position: {x: 0, y: 50},
    },
    {
        id: '2',
        type: 'custom',
        data: {name: 'Tyler Weary', job: 'Designer', emoji: '🤓'},
        position: {x: -200, y: 200},

    },
    {
        id: '3',
        type: 'custom',
        data: {name: 'Kristi Price', job: 'Developer', emoji: '🤩'},
        position: {x: 200, y: 200},

    },
];

const initialEdges = [
    {id: 'e1-2', source: '1', target: '2'},
    {id: 'e1-3', source: '1', target: '3'},
];



export default function OldGraph() {
    const [nodes, , onNodesChange] = useNodesState(initialNodes);
    const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);

    const onConnect = useCallback(
        (params: any) => setEdges((eds) => addEdge(params, eds)),
        [setEdges]
    );

    return (
        <div className="w-screen h-screen">
            <ReactFlow
                nodes={nodes}
                edges={edges}
                onNodesChange={onNodesChange}
                onEdgesChange={onEdgesChange}
                onConnect={onConnect}
                nodeTypes={nodeTypes}
                fitView
                className="bg-teal-50"
            >
                <Controls/>
                <MiniMap/>
            </ReactFlow>
        </div>
    );
};