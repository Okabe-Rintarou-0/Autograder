import { useMemoizedFn } from "ahooks";
import { Button, Card, Input, Space, Table } from "antd";
import { useState } from "react";
import { Email } from "../components/email";
import ImportCanvasUsersForm from "../components/import_canvas_users_form";
import RegisterUserForm from "../components/register_user_form";
import RoleTag from "../components/role_tag";
import { PAGE_SIZE } from "../lib/config";
import { useModal } from "../lib/hooks";
import { useUsers } from "../service/user";
import { formatDate } from "../utils/time";

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
    render: (email: string) => <Email email={email} />
}, {
    title: '角色',
    dataIndex: 'role',
    key: 'role',
    render: (role: number) => <RoleTag role={role} />
}];

export function UserTable() {
    const [pageNo, setPageNo] = useState<number>(1);
    const [keyword, setKeyword] = useState<string>("");
    const users = useUsers(keyword, pageNo, PAGE_SIZE);

    const { modal: RegisterUserModal,
        open: openRegisterUserModal,
        close: closeRegisterUserModal,
    } = useModal(RegisterUserForm, {});

    const { modal: ImportCanvasUsersModal,
        open: openImportCanvasUsersModal,
        close: closeImportCanvasUsersModal,
    } = useModal(ImportCanvasUsersForm, {});

    const onRegisterUser = useMemoizedFn(() => {
        closeRegisterUserModal();
        users.mutate();
    });

    return <>
        <RegisterUserModal onOk={onRegisterUser}
            destroyOnClose
            title={"导入"} footer={null} width={800}
            onCancel={closeRegisterUserModal} />
        <ImportCanvasUsersModal
            destroyOnClose
            title={"从 Canvas 导入用户"} footer={null} width={800}
            onOk={closeImportCanvasUsersModal}
            onCancel={closeImportCanvasUsersModal}
        />
        <Card className="card-container"
            title={
                <Input.Search style={{ width: "200px" }} placeholder="输入关键词" onSearch={setKeyword} allowClear />
            }
            extra={
                <Space>
                    <Button type="primary" onClick={openImportCanvasUsersModal}>从 Canvas 导入</Button>
                    <Button onClick={openRegisterUserModal}>导入</Button>
                    <Button onClick={() => users.mutate()}>刷新</Button>
                </Space>
            }>
            <Table columns={columns} dataSource={users.data?.data}
                rowKey="id"
                scroll={{ x: "100%" }}
                pagination={{
                    pageSize: PAGE_SIZE,
                    total: users.data?.total,
                    current: pageNo,
                    onChange: (pageNo: number) => {
                        setPageNo(pageNo);
                    }
                }}
            >
            </Table>
        </Card>
    </>
}