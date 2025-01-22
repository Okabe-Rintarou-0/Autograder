export interface BaseResp {
    error: string;
    message: string;
}

export interface ListItemResponse<T> extends BaseResp {
    total: number;
    data: T[];
}