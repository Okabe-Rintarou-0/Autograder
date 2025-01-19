import { MessageInstance } from 'antd/lib/message/interface';
import { BaseResp } from './../model/resp';

export async function handleBaseResp(messageApi: MessageInstance, resp: BaseResp, onSuccess?: () => void, onError?: () => void, duration = 0.8) {
    console.log(resp)
    if (resp.error) {
        await messageApi.error(resp.error, duration);
        onError?.();
        return;
    }

    await messageApi.success(resp.message, duration);
    onSuccess?.();
}