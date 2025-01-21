import { Tag } from "antd";
import { Administrator } from "../model/user";

export interface RoleTagProps {
    role?: number;
}

export default function RoleTag({ role }: RoleTagProps) {
    return <>
        {role === Administrator && <Tag color="red">管理员</Tag>}
        {role !== Administrator && <Tag color="green">普通用户</Tag>}
    </>
}