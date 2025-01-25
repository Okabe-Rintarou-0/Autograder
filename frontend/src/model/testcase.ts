import { BaseResp } from "./resp";

export const TestcaseStatusActive = 1;
export const TestcaseStatusInactive = 2;

export interface Testcase {
    id: number;
    name: string;
    status: number;
    content: string;
}

export interface BatchUpdateTestcaseRequest {
    data: Testcase[];
}

export interface BatchUpdateTestcaseResponse extends BaseResp { }
