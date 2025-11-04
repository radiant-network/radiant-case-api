export type Patient = {
    id: number;
    orgPatientId: string;
    sex: string;
    affectedStatus: string;
    relationToProband: string;
}

export type Sample = {
    id: number;
    orgSampleId: string;
    histological: string;
    patientId: number;
}

export type SequencingExperiment = {
    id: number;
    aliquotId: string;
    sampleId: number;
    experimentalStrategy: string;
    sequencingReadTechnology: string;
    platform: string;
}

export type Task = {
    id: number;
    name: string;
    sequencingExperimentId?: number;
}

export type Document = {
    id: number;
    name: string;
    size?: number; // size in bytes
    type?: string;
}

export interface Link {
    taskId: number;
    documentId: number;
    ioType: 'input' | 'output';
}

export interface GraphData {
    patients: Patient[];
    samples: Sample[];
    sequencingExperiments: SequencingExperiment[];
    tasks: Task[];
    documents: Document[];
    links: Link[];
}


