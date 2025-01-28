import { Button, Form, Input } from "antd";
import useMessage from "antd/es/message/useMessage";
import { ModalChildrenProps } from "../lib/hooks";
import { BaseResp } from "../model/resp";
import { submit } from "../service/common";
import { handleBaseResp } from "../utils/handle_resp";

const { Password } = Input;

interface ChangePasswordParam {
    password: string;
    confirm: string;
}

export default function ChangePasswordForm(props: ModalChildrenProps) {
    const { close } = props;
    const [form] = Form.useForm<ChangePasswordParam>();
    const [messageApi, contextHolder] = useMessage();

    const handleSubmit = async ({ password }: ChangePasswordParam) => {
        let resp = await submit<BaseResp>("/api/me/password", { password });
        handleBaseResp(messageApi, resp, close);
    };

    return <>
        {contextHolder}
        <Form
            form={form}
            layout="vertical"
            onFinish={handleSubmit}
            preserve={false}
        >
            <Form.Item
                name="password"
                label="新密码"
                required
                rules={[{ required: true }]}
            >
                <Password placeholder="请输入新密码" />
            </Form.Item>
            <Form.Item
                name="confirm"
                label="确认新密码"
                required
                rules={[
                    { required: true },
                    ({ getFieldValue }) => ({
                        validator(_, value) {
                            if (!value || getFieldValue('password') === value) {
                                return Promise.resolve()
                            }
                            return Promise.reject("两次密码输入不一致")
                        }
                    })
                ]}
            >
                <Password placeholder="请再次输入新密码" />
            </Form.Item>
            <Form.Item>
                <Button type="primary" htmlType="submit">
                    提交
                </Button>
            </Form.Item>
        </Form>
    </>
};