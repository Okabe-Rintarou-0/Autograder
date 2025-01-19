import { UploadFile } from "antd";
import { BaseResp } from "./resp";

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


