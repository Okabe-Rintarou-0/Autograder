import { UploadFile } from "antd";

export interface AppInfo {
    jdkVersion: number;
    authenticationType: number;
    file: UploadFile[]
}