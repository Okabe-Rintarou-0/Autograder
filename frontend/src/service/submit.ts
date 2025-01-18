import axios from 'axios';
import { AppInfo } from './../model/app';
import { GetProp, UploadProps } from 'antd';

type FileType = Parameters<GetProp<UploadProps, 'beforeUpload'>>[0];

export async function submitApp(info: AppInfo) {
    const formData = new FormData();
    formData.append('jdk_version', info.jdkVersion.toString());
    formData.append('authentication_type', info.authenticationType.toString());
    formData.append('file', info.file[0].originFileObj as FileType)

    await axios.post('/api/run', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}