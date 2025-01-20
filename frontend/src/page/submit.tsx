import { Button, Card, Form, Select } from "antd";
import { PrivateLayout } from "../components/layout";
import { AppInfo } from "../model/app";
import { submitApp } from "../service/task";
import useMessage from "antd/es/message/useMessage";
import React from "react";
import { handleBaseResp } from "../utils/handle_resp";
import { ZipUpload } from "../components/upload";

const supportedJDKVersions = [11, 17]
const { Option } = Select;

export default function SubmitPage() {
    const [form] = Form.useForm<AppInfo>();
    const [messageApi, contextHolder] = useMessage();
    const handleSubmit = async (appInfo: AppInfo) => {
        form.setFieldValue("file", []);
        try {
            const resp = await submitApp(appInfo);
            handleBaseResp(messageApi, resp);
        } catch (e) {
            console.log(e);
            messageApi.error(`上传失败：${e}`);
        }
    }

    return (
        <PrivateLayout>
            {contextHolder}
            <Card className="card-container">
                <Form
                    form={form}
                    layout="vertical"
                    onFinish={handleSubmit}
                    preserve={false}
                >
                    <Form.Item
                        key="jdkVersion"
                        name="jdkVersion"
                        label="JDK 版本"
                        rules={[{ required: true }]}
                    >
                        <Select placeholder="请选择您的 JDK 版本">
                            {supportedJDKVersions.map(version => <Option value={version}>{`JDK ${version}`}</Option>)}
                        </Select>
                    </Form.Item>
                    <Form.Item
                        key="authenticationType"
                        name="authenticationType"
                        label="鉴权方式"
                        rules={[{ required: true }]}
                    >
                        <Select placeholder="请选择您的鉴权方式">
                            <Option value={1}>Cookies</Option>
                            <Option value={2}>Token</Option>
                        </Select>
                    </Form.Item>
                    <Form.Item
                        key="file"
                        name="file"
                        label="上传项目压缩包（zip）"
                        rules={[{ required: true }]}
                    >
                        <ZipUpload />
                    </Form.Item>
                    <Form.Item>
                        <Button type="primary" htmlType="submit">
                            提交
                        </Button>
                    </Form.Item>
                </Form>
            </Card>
        </PrivateLayout>
    );
}
