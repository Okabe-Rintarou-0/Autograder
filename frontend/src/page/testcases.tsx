import { useMemoizedFn } from "ahooks";
import { Button, Card, Space, Switch, Table } from "antd";
import useMessage from "antd/es/message/useMessage";
import { PrivateLayout } from "../components/layout";
import { Testcase, TestcaseStatusActive, TestcaseStatusInactive } from "../model/testcase";
import { Administrator } from "../model/user";
import { batchUpdateTestcases, useTestcases } from "../service/testcase";
import { handleBaseResp } from "../utils/handle_resp";


export default function TestcasesPage() {
    const [messageApi, contextHolder] = useMessage();
    const testcases = useTestcases();

    const columns = [{
        title: '用例ID',
        dataIndex: 'id',
        key: 'id',
    }, {
        title: '路径',
        dataIndex: 'name',
        key: 'name',
    }, {
        title: '生效状态',
        dataIndex: 'status',
        key: 'status',
        render: (status: number, testcase: Testcase) => {
            return <Switch defaultValue={status === TestcaseStatusActive}
                onChange={(active) => {
                    testcase.status = active ? TestcaseStatusActive : TestcaseStatusInactive;
                    handleUpdateStatus(testcase);
                }}
            />
        }
    }];

    const handleUpdateStatus = useMemoizedFn(async (testcase: Testcase) => {
        const resp = await batchUpdateTestcases([testcase]);
        handleBaseResp(messageApi, resp);
    });

    return (
        <PrivateLayout forRole={Administrator}>
            {contextHolder}
            <Card className="card-container"
                extra={
                    <Space>
                        <Button onClick={() => testcases.mutate()}>刷新</Button>
                    </Space>
                }>
                <Table columns={columns} dataSource={testcases.data}
                    rowKey="id"
                    pagination={{
                        pageSize: 10,
                    }}
                >
                </Table>
            </Card>
        </PrivateLayout >
    )
}