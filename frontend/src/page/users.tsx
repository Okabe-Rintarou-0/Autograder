import useMessage from "antd/es/message/useMessage";
import { useEffect, useState } from "react";
import { Administrator, User } from "../model/user";
import { listUsers } from "../service/user";
import RoleTag from "../components/role_tag";
import { PrivateLayout } from "../components/layout";
import { Button, Card, Input, Space, Table } from "antd";
import { formatDate } from "../utils/time";
import RegisterUserModal from "../components/register_user_modal";
import ImportCanvasUsersModal from "../components/import_canvas_users_modal";
import Search from "antd/es/input/Search";

export default function UsersPage() {
    const pageSize = 20;
    const [messageApi, contextHolder] = useMessage();
    const [pageNo, setPageNo] = useState<number>(1);
    const [total, setTotal] = useState<number>(0);
    const [users, setUsers] = useState<User[]>([]);
    const [showRegisterModal, setShowRegisterModal] = useState<boolean>(false);
    const [showImportModal, setShowImportModal] = useState<boolean>(false);
    const [keyword, setKeyword] = useState<string>("");

    const onRegisterUser = () => {
        setShowRegisterModal(false);
        getUsers();
    };

    const onCancelRegisterUser = () => {
        setShowRegisterModal(false);
    };

    const onImportUsers = () => {
        setShowImportModal(false);
        getUsers();
    };

    const onCancelImportUsers = () => {
        setShowImportModal(false);
    };

    const getUsers = async () => {
        try {
            const resp = await listUsers(keyword, pageNo, pageSize);
            setUsers(resp.data);
            setTotal(resp.total);
        } catch (e) {
            console.log(e);
        }
    }

    useEffect(() => {
        getUsers();
    }, [keyword, pageNo]);

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
            <RegisterUserModal open={showRegisterModal} onOk={onRegisterUser} onCancel={onCancelRegisterUser} />
            <ImportCanvasUsersModal open={showImportModal} onOk={onImportUsers} onCancel={onCancelImportUsers} />
            <Card className="card-container"
                title={
                    <Input.Search style={{ width: "200px" }} placeholder="输入关键词" onSearch={setKeyword} />
                }
                extra={
                    <Space>
                        <Button type="primary" onClick={() => setShowImportModal(true)}>从 Canvas 导入</Button>
                        <Button onClick={() => setShowRegisterModal(true)}>导入</Button>
                        <Button onClick={getUsers}>刷新</Button>
                    </Space>
                }>
                <Table columns={columns} dataSource={users}
                    rowKey="id"
                    pagination={{
                        pageSize,
                        total,
                        current: pageNo,
                        onChange: (pageNo: number) => {
                            setPageNo(pageNo);
                        }
                    }}
                >
                </Table>
            </Card>
        </PrivateLayout >
    )
}