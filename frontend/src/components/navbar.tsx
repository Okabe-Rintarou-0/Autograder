import { CodeOutlined, FormOutlined, LogoutOutlined, MailOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Col, Dropdown, Row, Space } from "antd";
import { Link, useNavigate } from "react-router-dom";
import { useModal } from "../lib/hooks";
import { User } from "../model/user";
import { logout } from "../service/user";
import ChangePasswordForm from "./change_password_form";
import RoleTag from "./role_tag";

export default function NavBar({ me }: { me?: User }) {
    const {
        open: openChangePasswordModal,
        modal: ChangePasswordModal
    } = useModal(ChangePasswordForm, {});
    const navigate = useNavigate();

    const dropMenuItems = [
        {
            key: "username",
            label: <Space>
                <span>{`${me?.real_name}(${me?.username})`}</span>
                <RoleTag role={me?.role} />
            </Space>,
            icon: <UserOutlined />,
        },
        {
            key: "email",
            label: me?.email,
            icon: <MailOutlined />,
        },
        {
            key: "password",
            label: "修改密码",
            icon: <FormOutlined />,
        },
        {
            key: "compile",
            label: "修改编译/运行设置",
            icon: <CodeOutlined />,
        },
        { key: "/logout", label: "登出", icon: <LogoutOutlined />, danger: true },
    ];

    const handleMenuClick = async (e: any) => {
        if (e.key === "/logout") {
            logout();
            navigate("/");
            return;
        }
        if (e.key === "password") {
            openChangePasswordModal();
            return;
        }
        if (e.key.startsWith("/")) {
            const navigate = useNavigate();
            navigate(e.key);
        }
    };

    return (
        <Row className="navbar" justify="start">
            <Col>
                <Link to="/">Book Store</Link>
            </Col>
            <Col flex="auto" />
            {me && <Col>
                <Dropdown menu={{ onClick: handleMenuClick, items: dropMenuItems }}>
                    <Button shape="circle" icon={<UserOutlined />} />
                </Dropdown>
            </Col>}
            {me && <ChangePasswordModal
                title={"修改密码"}
                open
                onOk={close}
                onCancel={close}
                footer={null}
                width={800} />}
        </Row>
    );
}