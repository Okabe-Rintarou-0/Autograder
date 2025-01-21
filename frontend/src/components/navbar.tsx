import { Button, Col, Dropdown, Row, Space } from "antd";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { User } from "../model/user";
import { FormOutlined, LogoutOutlined, MailOutlined, UserOutlined } from "@ant-design/icons";
import { logout } from "../service/user";
import useMessage from "antd/es/message/useMessage";
import ChangePasswordModal from "./change_password_modal";
import RoleTag from "./role_tag";

export default function NavBar({ me }: { me?: User }) {
    const [showModal, setShowModal] = useState(false);
    const [messageApi, contextHolder] = useMessage();
    const navigate = useNavigate();

    const handleOpenModal = () => {
        setShowModal(true);
    }

    const handleCloseModal = () => {
        setShowModal(false);
    }
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
        { key: "/logout", label: "登出", icon: <LogoutOutlined />, danger: true },
    ];

    const handleMenuClick = async (e: any) => {
        if (e.key === "/logout") {
            logout();
            navigate("/");
            return;
        }
        if (e.key === "password") {
            handleOpenModal();
            return;
        }
        if (e.key.startsWith("/")) {
            const navigate = useNavigate();
            navigate(e.key);
        }
    };

    return (
        <Row className="navbar" justify="start">
            {contextHolder}
            <Col>
                <Link to="/">Book Store</Link>
            </Col>
            <Col flex="auto">
            </Col>
            {me && <Col>
                <Dropdown menu={{ onClick: handleMenuClick, items: dropMenuItems }}>
                    <Button shape="circle" icon={<UserOutlined />} />
                </Dropdown>
            </Col>}
            {me && showModal && <ChangePasswordModal onOk={handleCloseModal} onCancel={handleCloseModal} />}
        </Row>
    );
}