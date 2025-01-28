import { Form, Select } from "antd";
import useMessage from "antd/es/message/useMessage";
import { useContext, useEffect } from "react";
import { SUPPORTED_JDK_VERSIONS } from "../lib/config";
import { UserContext } from "../lib/context";
import { ModalChildrenProps } from "../lib/hooks";
import { BaseResp } from "../model/resp";
import { submit } from "../service/common";
import { handleBaseResp } from "../utils/handle_resp";

const { Option } = Select;

interface UpdateCompilationInfoParam {
    jdk_version: number;
    authentication_type: number;
}

export default function UpdateCompilationInfoForm(props: ModalChildrenProps) {
    const { close } = props;
    const me = useContext(UserContext);
    const [form] = Form.useForm<UpdateCompilationInfoParam>();
    const [messageApi, contextHolder] = useMessage();

    const handleSubmit = async (param: UpdateCompilationInfoParam) => {
        let resp = await submit<BaseResp>('/api/me/compile', param);
        handleBaseResp(messageApi, resp, close);
    };

    useEffect(() => {
        if (me) {
            form.setFieldsValue({
                jdk_version: me.jdk_version,
                authentication_type: me.authentication_type
            })
        }
    }, [me]);

    return <>
        {contextHolder}
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
                    {SUPPORTED_JDK_VERSIONS.map(version => <Option value={version} key={version}>{`JDK ${version}`}</Option>)}
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
        </Form>
    </>
};