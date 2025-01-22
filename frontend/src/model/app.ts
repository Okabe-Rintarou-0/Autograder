import { UploadFile } from "antd";
import { BaseResp, ListItemResponse } from "./resp";

export interface AppInfo {
    jdkVersion: number;
    authenticationType: number;
    file: UploadFile[]
}

export interface SubmitAppResponse extends BaseResp { }


export const AppRunTaskStatusWaiting = 1
export const AppRunTaskStatusRunning = 2
export const AppRunTaskStatusSucceed = 3
export const AppRunTaskStatusFail = 4


export interface AppRunTask {
    uuid: string;
    user_id: number;
    status: number;
    created_at: string;
    pass: number;
    total: number;
    test_results: string | null;
}

export interface ListAppTasksResponse extends ListItemResponse<AppRunTask> { }