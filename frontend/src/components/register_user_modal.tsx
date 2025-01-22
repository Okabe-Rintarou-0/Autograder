import { Button, Form, Input, Modal } from "antd";
import useMessage from "antd/es/message/useMessage";
import { handleBaseResp } from "../utils/handle_resp";
import { RegisterUserRequest, registerUserURL } from "../model/user";
import { submit } from "../service/common";
import { BaseResp } from "../model/resp";

const { Password } = Input;

interface RegisterUserModalProps {
    open: boolean;
    onOk: () => void;
    onCancel: () => void;
}

export default function RegisterUserModal(props: RegisterUserModalProps) {
    const { open, onOk, onCancel } = props;
    const [form] = Form.useForm<RegisterUserRequest>();
    const [messageApi, contextHolder] = useMessage();

    const handleSubmit = async (request: RegisterUserRequest) => {
        let resp = await submit<BaseResp>(registerUserURL, request);
        handleBaseResp(messageApi, resp, onOk);
    };

    return (
        <Modal
            destroyOnClose
            title={"修改密码"}
            open={open}
            onOk={onOk}
            onCancel={onCancel}
            footer={null}
            width={800}
        >
            {contextHolder}
            <Form
                form={form}
                layout="vertical"
                onFinish={handleSubmit}
                preserve={false}
            >
                <Form.Item
                    name="username"
                    label="用户名（填写学号）"
                    required
                    rules={[{ required: true }]}
                >
                    <Input placeholder="请输入用户名" />
                </Form.Item>
                <Form.Item
                    name="real_name"
                    label="真名"
                    required
                    rules={[{ required: true }]}
                >
                    <Input placeholder="请输入真名" />
                </Form.Item>
                <Form.Item
                    name="email"
                    label="邮箱"
                    required
                    rules={[{ required: true }]}
                >
                    <Input placeholder="请输入邮箱" />
                </Form.Item>
                <Form.Item
                    name="password"
                    label="密码"
                    required
                    rules={[{ required: true }]}
                >
                    <Password placeholder="请输入密码" />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        提交
                    </Button>
                </Form.Item>
            </Form>
        </Modal>
    );
};