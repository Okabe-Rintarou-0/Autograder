import axios from "axios";
import useSWR from "swr";
import { BatchUpdateTestcaseRequest, BatchUpdateTestcaseResponse, Testcase } from "../model/testcase";
import { fetcher } from "./common";

export function useTestcases() {
    return useSWR<Testcase[]>('/api/testcases', fetcher);
}

export async function batchUpdateTestcases(data: Testcase[]) {
    const request = {
        data
    } as BatchUpdateTestcaseRequest;
    const resp = await axios.post<BatchUpdateTestcaseResponse>('/api/testcases', request);
    return resp.data;
}