import "ace-builds/src-noconflict/ext-language_tools";
import "ace-builds/src-noconflict/mode-toml";
import "ace-builds/src-noconflict/theme-github";
import { useMemoizedFn } from "ahooks";
import { Button, Card, Space, Switch, Table } from "antd";
import useMessage from "antd/es/message/useMessage";
import React, { useMemo } from "react";
import AceEditor from "react-ace";
import { useImmer } from 'use-immer';
import { Testcase, TestcaseStatusActive, TestcaseStatusInactive } from "../model/testcase";
import { batchUpdateTestcases, useTestcases } from "../service/testcase";
import { handleBaseResp } from "../utils/handle_resp";

export function TestcaseTable() {
    const [messageApi, contextHolder] = useMessage();
    const [editingTestcaseIDs, setEditingTestcaseIDs] = useImmer<Set<number>>(new Set<number>);
    const [expandedRowKeys, setExpandedRowKeys] = useImmer<React.Key[]>([]);
    const testcases = useTestcases();
    const handleEditTestcase = useMemoizedFn((testcase: Testcase) => {
        const targetID = testcase.id;
        setEditingTestcaseIDs(ids => {
            if (!ids.has(targetID)) {
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
            编辑
        </Button>
    }], []);

    const handleUpdateStatus = useMemoizedFn(async (testcase: Testcase) => {
        const resp = await batchUpdateTestcases([testcase]);
        handleBaseResp(messageApi, resp);
    });

    return <Card className="card-container"
        extra={<Button onClick={() => testcases.mutate()}>刷新</Button>}>
        {contextHolder}
        <Table columns={columns} dataSource={testcases.data}
            rowKey="id"
            scroll={{ x: "100%" }}
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
                    return <Space direction="vertical" style={{ width: "100%" }}>
                        <AceEditor mode="toml"
                            theme="github"
                            readOnly={!editing}
                            style={{ width: "100%" }}
                            onChange={(v) => testcase.content = v}
                            defaultValue={testcase.content}
                        />
                        {editing && <>
                            <Space>
                                <Button type="primary" onClick={() => {
                                    handleUpdateStatus(testcase);
                                    setEditingTestcaseIDs(ids => { ids.delete(testcase.id) });
                                }}>保存</Button>
                                <Button onClick={() => setEditingTestcaseIDs(ids => { ids.delete(testcase.id) })}>取消</Button>
                            </Space>
                        </>}
                    </Space>;
                }
            }}
        >
        </Table >
    </Card>
}
