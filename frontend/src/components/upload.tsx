import { UploadOutlined } from "@ant-design/icons"
import { Button, Upload, UploadFile } from "antd"
import React from "react"
import { useEffect, useState } from "react"

interface UploadProps {
    maxCount?: number
    value?: UploadFile[]
    onChange?: (value?: UploadFile[]) => void
}

export const ZipUpload: React.FC<UploadProps> = ({ maxCount, value, onChange }) => {
    const [fileList, setFileList] = useState<UploadFile[]>([])
    useEffect(() => {
        value && setFileList(value)
    }, [value]);

    return (
        <Upload
            fileList={fileList}
            maxCount={maxCount}
            onChange={e => {
                if (Array.isArray(e)) {
                    onChange?.(e)
                    return e
                }
                setFileList(e.fileList);
                e && onChange?.(e.fileList)
                return e && e.fileList
            }}
            accept=".zip"
            customRequest={({ onSuccess }) => {
                onSuccess?.("ok");
            }}
        >
            <Button icon={<UploadOutlined />}>上传</Button>
        </Upload>
    )
}
