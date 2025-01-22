import axios from "axios";
import { BaseResp } from "../model/resp";

export async function submit<R extends BaseResp>(url: string, params: any) {
    const formData = new FormData();
    for (const key in params) {
        if (params.hasOwnProperty(key)) {
            formData.append(key, params[key]);
        }
    }
    const resp = await axios.post<R>(url, formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
    return resp.data;
} 