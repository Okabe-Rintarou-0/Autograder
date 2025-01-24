import {
    LockOutlined,
    UserOutlined,
} from '@ant-design/icons';
import { LoginFormPage, ProFormText } from '@ant-design/pro-components';
import useMessage from "antd/es/message/useMessage";
import { useNavigate } from "react-router-dom";
import { BasicLayout } from "../components/layout";
import { LoginRequest } from "../model/user";
import { setToken } from "../service/token";
import { login } from "../service/user";
import { handleBaseResp } from "../utils/handle_resp";

const LoginPage = () => {
    const [messageApi, contextHolder] = useMessage();
    const navigate = useNavigate();

    const onSubmit = async (request: LoginRequest) => {
        try {
            let res = await login(request);
            handleBaseResp(messageApi, res, () => {
                setToken(res.token);
                navigate("/submit");
            });
        } catch (e) {
            console.log(e);
        }
    };

    return (
        <BasicLayout noSider>
            {contextHolder}
            <LoginFormPage
                backgroundImageUrl={"/login.png"}
                logo={"/logo.webp"}
                title="Book Store"
                subTitle="电子书城"
                onFinish={onSubmit}
                style={{ height: "80vh" }}
            >
                <ProFormText
                    name="identifier"
                    fieldProps={{
                        size: 'large',
                        prefix: <UserOutlined className={'prefixIcon'} />,
                    }}
                    placeholder={'请输入用户名（邮箱）'}
                    rules={[
                        {
                            required: true,
                            message: '请输入用户名!',
                        },
                    ]}
                />
                <ProFormText.Password
                    name="password"
                    fieldProps={{
                        size: 'large',
                        prefix: <LockOutlined className={'prefixIcon'} />,
                    }}
                    placeholder={'密码'}
                    rules={[
                        {
                            required: true,
                            message: '请输入密码！',
                        },
                    ]}
                />
                <div
                    style={{
                        marginBlockEnd: 24,
                    }}
                >
                </div>
            </LoginFormPage>
        </BasicLayout>
    );
};
export default LoginPage;