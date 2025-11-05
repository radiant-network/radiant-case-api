import Graph from './components/Graph.tsx';
import {fetchCase} from "./api.ts";
import { useEffect, useState } from "react";


export default function App() {
    const [selectedCase, setSelectedCase] = useState("trio");
    const [caseData, setCaseData] = useState<any>(null);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        setLoading(true);
        fetchCase(selectedCase)
            .then((data) => setCaseData(data))
            .finally(() => setLoading(false));
    }, [selectedCase]);

    return (
        <div className="flex flex-col h-screen">
            <header className="p-4 border-b bg-gray-100 flex items-center gap-2">
                <label htmlFor="caseSelect" className="font-semibold text-gray-700">
                    Select case:
                </label>
                <select
                    id="caseSelect"
                    value={selectedCase}
                    onChange={(e) => setSelectedCase(e.target.value)}
                    className="border rounded px-2 py-1"
                >
                    <option value="trio">Germline - Trio</option>
                    <option value="solo">Germline - Solo</option>
                    <option value="somatic">Somatic</option>
                </select>
            </header>

            <main className="flex-1">
                {loading && <div className="p-4">Loading...</div>}
                {caseData && <Graph caseData={caseData} />}
            </main>
        </div>
    );
}