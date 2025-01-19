import { BaseResp } from "./resp";

export interface LoginRequest {
    identifier: string;
    password: string;
}

export interface LoginResponse extends BaseResp {
    token: string;
}

export interface User extends BaseResp {
    username: string;
    email: string;
}

export interface AppRunTask {
    uuid: string;
    user_id: number;
    status: number;
    created_at: string;
}

export interface ListAppTasksResponse {
    total: number;
    data: AppRunTask[];
}