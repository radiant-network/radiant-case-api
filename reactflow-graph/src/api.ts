import trio from "./mocks/trio.json";
import solo from "./mocks/solo.json";
import somatic from "./mocks/somatic.json";

const cases = { trio, solo, somatic};
import type {Case} from "./types.ts";

export async function fetchCase(caseId: string): Promise<Case> {
    const data = cases[caseId as keyof typeof cases];
    if (!data) throw new Error(`Unknown case: ${caseId}`);
    return data;
}