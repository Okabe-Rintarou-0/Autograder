import axios from 'axios';
import { AppInfo, SubmitAppResponse, ListAppTasksResponse } from '../model/app';
import { GetProp, UploadProps } from 'antd';
import useSWR from 'swr';
import { fetcher } from './common';

type FileType = Parameters<GetProp<UploadProps, 'beforeUpload'>>[0];


export async function submitApp(info: AppInfo) {
    const formData = new FormData();
    formData.append('jdk_version', info.jdkVersion.toString());
    formData.append('authentication_type', info.authenticationType.toString());
    formData.append('file', info.file[0].originFileObj as FileType)

    const resp = await axios.post<SubmitAppResponse>('/api/run', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
    return resp.data;
}

export function useTasks(pageNo: number, pageSize: number) {
    return useSWR<ListAppTasksResponse>(`/api/tasks?page_no=${pageNo}&page_size=${pageSize}`, fetcher);
}

export async function readLog(uuid: string, logType: string) {
    const resp = await axios.get<string>("/api/logs", {
        params: {
            uuid,
            log_type: logType
        }
    });
    return resp.data;
}