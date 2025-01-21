import useMessage from "antd/es/message/useMessage";
import { useEffect, useState } from "react";
import { Administrator, User } from "../model/user";
import { listUsers } from "../service/user";
import RoleTag from "../components/role_tag";
import { PrivateLayout } from "../components/layout";
import { Button, Card, Table } from "antd";
import { formatDate } from "../utils/time";

export default function UsersPage() {
    const pageSize = 20;
    const [messageApi, contextHolder] = useMessage();
    const [pageNo, setPageNo] = useState<number>(1);
    const [total, setTotal] = useState<number>(0);
    const [users, setUsers] = useState<User[]>([]);

    const getUsers = async () => {
        try {
            const resp = await listUsers(pageNo, pageSize);
            setUsers(resp.data);
            setTotal(resp.total);
        } catch (e) {
            console.log(e);
        }
    }

    useEffect(() => {
        getUsers();
    }, []);

    const columns = [{
        title: '创建时间',
        dataIndex: 'created_at',
        key: 'created_at',
        render: formatDate
    }, {
        title: '用户ID',
        dataIndex: 'id',
        key: 'id',
    }, {
        title: '用户名',
        dataIndex: 'username',
        key: 'username',
    }, {
        title: '真名',
        dataIndex: 'real_name',
        key: 'real_name',
    }, {
        title: '邮箱',
        dataIndex: 'email',
        key: 'email',
        render: (email: string) => <a href={`mailto:${email}`}>{email}</a>
    }, {
        title: '角色',
        dataIndex: 'role',
        key: 'role',
        render: (role: number) => <RoleTag role={role} />
    }];

    return (
        <PrivateLayout forRole={Administrator}>
            {contextHolder}
            <Card className="card-container" extra={<Button type="primary" onClick={getUsers}>刷新</Button>}>
                <Table columns={columns} dataSource={users}
                    rowKey="id"
                    pagination={{
                        pageSize,
                        total,
                        current: pageNo,
                        onChange: (pageNo: number) => {
                            console.log(pageNo);
                            setPageNo(pageNo);
                            getUsers();
                        }
                    }}
                >
                </Table>
            </Card>
        </PrivateLayout >
    )
}