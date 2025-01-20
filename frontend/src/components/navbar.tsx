import { Button, Col, Dropdown, Row } from "antd";
import React from "react";
import { Link } from "react-router-dom";
import { User } from "../model/user";
import { FormOutlined, LogoutOutlined, UserOutlined } from "@ant-design/icons";
import { logout } from "../service/user";

export default function NavBar({ me }: { me?: User }) {
    const dropMenuItems = [
        {
            key: "username",
            label: me?.username,
            icon: <UserOutlined />,
        },
        {
            key: "password",
            label: "修改密码",
            icon: <FormOutlined />,
        },
        { key: "/logout", label: "登出", icon: <LogoutOutlined />, danger: true },
    ];

    const handleMenuClick = async (e) => {
        if (e.key === "/logout") {
            logout();
            return;
        }
    };

    return (
        <Row className="navbar" justify="start">
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
        </Row>
    );
}