import { Card, Space, Table, Tabs, Tag } from "antd";
import BasicLayout from "../components/layout";
import { listTasks } from "../service/task";
import useMessage from "antd/es/message/useMessage";
import React, { useEffect, useState } from "react";
import { LazyLog } from "@melloware/react-logviewer";
import { AppRunTask } from "../model/user";
import { AppRunTaskStatusFail, AppRunTaskStatusRunning, AppRunTaskStatusSucceed, AppRunTaskStatusWaiting } from "../model/app";

export default function TaskPage() {
    const pageSize = 20;
    const [messageApi, contextHolder] = useMessage();
    const [pageNo, setPageNo] = useState<number>(1);
    const [total, setTotal] = useState<number>(0);
    const [tasks, setTasks] = useState<AppRunTask[]>([]);

    const getTasks = async () => {
        try {
            const resp = await listTasks(pageNo, pageSize);
            setTasks(resp.data);
            setTotal(resp.total);
        } catch (e) {
            console.log(e);
        }
    }

    useEffect(() => {
        getTasks();
    }, []);

    const columns = [{
        title: '创建时间',
        dataIndex: 'created_at',
        key: 'created_at',
    }, {
        title: '任务ID',
        dataIndex: 'uuid',
        key: 'uuid',
    }, {
        title: '状态',
        dataIndex: 'status',
        key: 'status',
        render: (status: number) => {
            if (status === AppRunTaskStatusFail) {
                return <Tag color="red">运行失败</Tag>
            } else if (status === AppRunTaskStatusRunning) {
                return <Tag color="blue">运行中</Tag>
            } if (status === AppRunTaskStatusWaiting) {
                return <Tag color="yellow">等待运行</Tag>
            } if (status === AppRunTaskStatusSucceed) {
                return <Tag color="green">执行成功</Tag>
            }
        }
    }];

    const getTabs = (task: AppRunTask) => {
        return ["stdout", "stderr", "hurl"].map(logType => ({
            key: logType,
            label: (logType === "hurl" ? "测试" : logType) + '日志',
            children: <LazyLog caseInsensitive
                enableHotKeys
                enableSearch
                extraLines={1}
                height="520"
                selectableLines
                url={`/api/logs?uuid=${task.uuid}&log_type=${logType}`} />
        }));
    }

    return (
        <BasicLayout>
            {contextHolder}
            <Card className="card-container">
                <Table columns={columns} dataSource={tasks}
                    rowKey="uuid"
                    pagination={{
                        pageSize,
                        total,
                        current: pageNo,
                        onChange: (pageNo: number) => {
                            console.log(pageNo);
                            setPageNo(pageNo);
                            getTasks();
                        }
                    }}
                    expandable={{
                        expandedRowRender: (task) => (
                            <Space direction="vertical" size={"large"} style={{ width: "100%" }}>
                                <Tabs items={getTabs(task)} />
                            </Space>
                        ),
                    }}
                >
                </Table>
            </Card>
        </BasicLayout >
    )
}