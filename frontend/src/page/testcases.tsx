import { useMemoizedFn } from "ahooks";
import { Button, Card, Input, Space, Switch, Table } from "antd";
import useMessage from "antd/es/message/useMessage";
import React, { useMemo } from "react";
import { Prism, SyntaxHighlighterProps } from 'react-syntax-highlighter';
import { useImmer } from 'use-immer';
import { PrivateLayout } from "../components/layout";
import { Testcase, TestcaseStatusActive, TestcaseStatusInactive } from "../model/testcase";
import { Administrator } from "../model/user";
import { batchUpdateTestcases, useTestcases } from "../service/testcase";
import { handleBaseResp } from "../utils/handle_resp";

const { TextArea } = Input;

const SyntaxHighlighter = (Prism as any) as typeof React.Component<SyntaxHighlighterProps>;
export default function TestcasesPage() {
    const [messageApi, contextHolder] = useMessage();
    const [editingTestcaseIDs, setEditingTestcaseIDs] = useImmer<Set<number>>(new Set<number>);
    const [expandedRowKeys, setExpandedRowKeys] = useImmer<React.Key[]>([]);
    const testcases = useTestcases();
    const handleEditTestcase = useMemoizedFn((testcase: Testcase) => {
        const targetID = testcase.id;
        setEditingTestcaseIDs(ids => {
            if (ids.has(targetID)) {
                ids.delete(targetID);
            } else {
                ids.add(targetID);
            }
        });
        setExpandedRowKeys(keys => {
            if (!keys.find(key => key === targetID)) {
                keys.push(targetID);
            }
        });
    });

    const columns = useMemo(() => [{
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
    }, {
        title: "操作",
        key: 'action',
        render: (_: any, testcase: Testcase) => <Button onClick={() => handleEditTestcase(testcase)}>
            编辑/取消
        </Button>
    }], []);

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
                    expandable={{
                        expandedRowKeys,
                        onExpandedRowsChange: (keys) => {
                            setExpandedRowKeys([...keys]);
                        },
                        expandedRowRender: (testcase: Testcase) => {
                            const editing = editingTestcaseIDs.has(testcase.id);
                            return <>
                                {editing && <TextArea defaultValue={testcase.content} style={{ minHeight: "500px" }} />}
                                {!editing && <SyntaxHighlighter language="toml">{testcase.content}</SyntaxHighlighter>}
                            </>;
                        }
                    }}
                >
                </Table >
            </Card>
        </PrivateLayout >
    )
}