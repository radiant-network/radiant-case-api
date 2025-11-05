import {useCallback, useEffect} from 'react';
import {
    Controls,
    type Edge,
    MarkerType,
    MiniMap,
    type Node,
    ReactFlow,
    useEdgesState,
    useNodesState
} from '@xyflow/react';

import dagre from '@dagrejs/dagre';
import {type Case} from '../types';
import DocumentNode from '../components/nodes/DocumentNode';
import TaskNode from '../components/nodes/TaskNode';
import SequencingExperimentNode from "./nodes/SequencingExperimentNode.tsx";
import SampleNode from "./nodes/SampleNode.tsx";
import PatientNode from "./nodes/PatientNode.tsx";
// ---- Layout Configuration ----
const dagreGraph = new dagre.graphlib.Graph().setDefaultEdgeLabel(() => ({}));
const nodeWidth = 200;
const nodeHeight = 100;

function applyLayout(nodes: Node[], edges: Edge[]) {
    dagreGraph.setGraph({rankdir: 'LR', align: 'UL', nodesep: 50, ranksep: 100});

    nodes.forEach((node) => {
        dagreGraph.setNode(node.id, {width: nodeWidth, height: nodeHeight});
    });
    edges.forEach((edge) => {
        dagreGraph.setEdge(edge.source, edge.target);
    });

    dagre.layout(dagreGraph);

    return nodes.map((node) => {
        const pos = dagreGraph.node(node.id);
        return {
            ...node,
            position: {x: pos.x - nodeWidth / 2, y: pos.y - nodeHeight / 2},
        };
    });
}

// ---- Define these outside (or memoize) ----
const nodeTypes = {
    document: DocumentNode,
    task: TaskNode,
    sequencingExperiment: SequencingExperimentNode,
    sample: SampleNode,
    patient: PatientNode
};
const edgeTypes = {}; // No custom edges yet, but stable reference


export default function Graph({ caseData }: { caseData: Case }) {

    const [nodes, setNodes, onNodesChange] = useNodesState([]);
    const [edges, setEdges, onEdgesChange] = useEdgesState([]);

    useEffect(() => {
        if (!caseData) return;

        const patientNodes: Node[] = caseData.patients.map((patient) => ({
            id: `patient-${patient.id}`,
            type: 'patient',
            data: patient,
            position: {x: 0, y: 0},
        }));

        const sampleNodes: Node[] = caseData.samples.map((sample) => ({
            id: `sample-${sample.id}`,
            type: 'sample',
            data: sample,
            position: {x: 0, y: 0},
        }));

        const sequencingExperimentNodes: Node[] = caseData.sequencingExperiments.map((sequencingExperiment) => ({
            id: `seqExp-${sequencingExperiment.id}`,
            type: 'sequencingExperiment',
            data: sequencingExperiment,
            position: {x: 0, y: 0},
        }));

        const docNodes: Node[] = caseData.documents.map((doc) => ({
            id: `doc-${doc.id}`,
            type: 'document',
            data: doc,
            position: {x: 0, y: 0},
        }));

        const taskNodes: Node[] = caseData.tasks.map((task) => ({
            id: `task-${task.id}`,
            type: 'task',
            data: task,
            position: {x: 0, y: 0},
        }));

        const samplePatientEdges: Edge[] = caseData.samples.map((sample) => ({
                id: `e-patient-${sample.patientId}-sample-${sample.id}`,
                source: `patient-${sample.patientId}`,
                target: `sample-${sample.id}`,
                markerEnd: {
                    type: MarkerType.ArrowClosed, width: 20,
                    height: 20,
                    color: 'var(--color-slate-400)',
                },
                type: 'smoothstep',
                // animated: true,
                style: {
                    strokeWidth: 2,
                    stroke: 'var(--color-slate-400)',
                }
            }
        ));

        const sequencingExperimentSampleEdges: Edge[] = caseData.sequencingExperiments.map((sequencingExperiment) => ({
                id: `e-sample-${sequencingExperiment.id}-seqExp-${sequencingExperiment.sampleId}`,
                source: `sample-${sequencingExperiment.sampleId}`,
                target: `seqExp-${sequencingExperiment.id}`,
                markerEnd: {
                    type: MarkerType.ArrowClosed, width: 20,
                    height: 20,
                    color: 'var(--color-slate-400)',
                },
                type: 'smoothstep',
                // animated: true,
                style: {
                    strokeWidth: 2,
                    stroke: 'var(--color-slate-400)',
                }
            }
        ));


        const taskSequencingExperimentEdges: Edge[] = caseData.tasks.flatMap((task) => {
            // Only create edge if there is a sequencingExperimentId and no input document link exists
            if (task.sequencingExperimentId && !caseData.links.find((link) => link.taskId === task.id && link.ioType === 'input')) {
                return [{
                    id: `e-seqExp-${task.sequencingExperimentId}-task-${task.id}`,
                    source: `seqExp-${task.sequencingExperimentId}`,
                    target: `task-${task.id}`,
                    markerEnd: {
                        type: MarkerType.ArrowClosed, width: 20,
                        height: 20,
                        color: 'var(--color-slate-400)',
                    },
                    type: 'smoothstep',
                    // animated: true,
                    style: {
                        strokeWidth: 2,
                        stroke: 'var(--color-slate-400)',
                    },
                    targetHandle: 'left'
                }]
            } else return [];
        });

        const taskDocumentEdges: Edge[] = caseData.links.map((link) =>
            link.ioType === 'input'
                ? {
                    id: `e-doc-${link.documentId}-task-${link.taskId}`,
                    source: `doc-${link.documentId}`,
                    target: `task-${link.taskId}`,
                    markerEnd: {
                        type: MarkerType.ArrowClosed, width: 20,
                        height: 20,
                        color: 'var(--color-slate-400)',
                    },
                    type: 'smoothstep',
                    // animated: true,
                    style: {
                        strokeWidth: 2,
                        stroke: 'var(--color-slate-400)',
                    },
                    targetHandle: 'left'
                }
                : {
                    id: `e-task-${link.taskId}-doc-${link.documentId}`,
                    source: `task-${link.taskId}`,
                    target: `doc-${link.documentId}`,
                    markerEnd: {
                        type: MarkerType.ArrowClosed, width: 20,
                        height: 20,
                        color: 'var(--color-slate-400)',
                    },
                    type: 'smoothstep',
                    animated: true,
                    style: {
                        strokeWidth: 2,
                        stroke: 'var(--color-slate-400)',
                    }
                }
        );

        const allEdges = [...taskSequencingExperimentEdges, ...taskDocumentEdges, ...sequencingExperimentSampleEdges, ...samplePatientEdges];
        const allNodes = [...sequencingExperimentNodes, ...docNodes, ...taskNodes, ...sampleNodes, ...patientNodes];
        const layouted = applyLayout(allNodes, allEdges);
        setNodes(layouted);
        setEdges(allEdges);

    }, [caseData]);
    const onNodeClick = useCallback((event: React.MouseEvent, node: any) => {
        console.log('node clicked', node);
        console.log('event', event);
        // here you can trigger opening the panel and set selected node id/data
    }, []);


    return (
        <div className="w-screen h-screen">
            <ReactFlow
                nodes={nodes}
                edges={edges}
                nodeTypes={nodeTypes}
                edgeTypes={edgeTypes}
                onNodesChange={onNodesChange}
                onEdgesChange={onEdgesChange}
                onNodeClick={onNodeClick}
                fitView
                className="bg-teal-50"
            >
                <Controls/>
                <MiniMap/>
            </ReactFlow>
        </div>
    );
}
