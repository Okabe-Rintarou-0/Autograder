import { CodeOutlined, FormOutlined, LogoutOutlined, MailOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Col, Dropdown, Row, Space } from "antd";
import { useContext } from "react";
import { Link, useNavigate } from "react-router-dom";
import { UserContext } from "../lib/context";
import { useModal } from "../lib/hooks";
import { logout } from "../service/user";
import ChangePasswordForm from "./change_password_form";
import RoleTag from "./role_tag";
import UpdateCompilationInfoForm from "./update_compilaton_info_form";

export default function NavBar() {
    const me = useContext(UserContext);
    const {
        open: openChangePasswordModal,
        close: closeChangePasswordModal,
        modal: ChangePasswordModal
    } = useModal(ChangePasswordForm, {});
    const {
        open: openCompilationInfoModal,
        close: closeCompilationInfoModal,
        modal: CompilationInfoModal
    } = useModal(UpdateCompilationInfoForm, {});
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
        if (e.key === "compile") {
            openCompilationInfoModal();
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
                onOk={closeChangePasswordModal}
                onCancel={closeChangePasswordModal}
                footer={null}
                width={800} />}
            {me && <CompilationInfoModal
                title={"修改编译/运行设置"}
                open
                onOk={closeCompilationInfoModal}
                onCancel={closeCompilationInfoModal}
                footer={null}
                width={800} />}
        </Row>
    );
}