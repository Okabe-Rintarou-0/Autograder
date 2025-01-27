import { MailOutlined, UserOutlined } from "@ant-design/icons";
import { Dropdown, Space, Tag } from "antd";
import { useMemo } from "react";
import { UserProfile } from "../model/app";
import { Email } from "./email";
import RoleTag from "./role_tag";

interface UserProfileDropdownProps {
    user: UserProfile;
}

export function UserProfileDropdown(props: UserProfileDropdownProps) {
    const { user } = props;
    const items = useMemo(() => [
        {
            key: 'username',
            label: <Space>
                <span>{user.username}</span>
                <RoleTag role={user.role} />
            </Space>,
            icon: <UserOutlined />,
        },
        {
            key: 'email',
            label: <Space>
                <span>邮箱: </span>
                <Email email={user.email} />
            </Space>,
            icon: <MailOutlined />,
        },
    ], [user]);

    return <Dropdown menu={{ items }}>
        <Tag>{user.real_name}</Tag>
    </Dropdown>
}