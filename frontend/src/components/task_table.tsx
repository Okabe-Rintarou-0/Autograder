import { LazyLog } from "@melloware/react-logviewer";
import { Badge, Button, Card, Space, Table, Tabs, Tag } from "antd";
import { BadgeProps } from "antd/lib";
import { useContext, useEffect, useState } from "react";
import ReactJson from "react-json-view-ts";
import { PAGE_SIZE } from "../lib/config";
import { UserContext } from "../lib/context";
import { AppRunTask, AppRunTaskStatusFail, AppRunTaskStatusRunning, AppRunTaskStatusSucceed, AppRunTaskStatusWaiting, UserProfile } from "../model/app";
import { Administrator } from "../model/user";
import { useTasks } from "../service/task";
import { formatDate } from "../utils/time";
import { UserProfileDropdown } from "./user_profile_dropdown";
import UserSelect from "./user_select";

const columns = [{
    title: '创建者',
    dataIndex: 'operator',
    key: 'operator',
    render: (operator: UserProfile) => {
        return <UserProfileDropdown user={operator} />
    },
}, {
    title: '用户',
    dataIndex: 'user',
    key: 'user',
    render: (user: UserProfile) => {
        return <UserProfileDropdown user={user} />
    },
}, {
    title: '创建时间',
    dataIndex: 'created_at',
    key: 'created_at',
    render: formatDate,
}, {
    title: '任务ID',
    dataIndex: 'uuid',
    key: 'uuid',
}, {
    title: '通过率',
    dataIndex: 'pass',
    key: 'pass',
    render: (pass: number, task: AppRunTask) => {
        let status: BadgeProps["status"] = "success";
        if (pass != task.total) {
            status = "error"
        }
        return <Badge status={status} text={`${pass}/${task.total}`} />
    }
}, {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
    render: (status: number) => {
        switch (status) {
            case AppRunTaskStatusFail:
                return <Tag color="red">运行失败</Tag>;
            case AppRunTaskStatusRunning:
                return <Tag color="blue">运行中</Tag>;
            case AppRunTaskStatusWaiting:
                return <Tag color="yellow">等待运行</Tag>;
            case AppRunTaskStatusSucceed:
                return <Tag color="green">执行成功</Tag>;
            default:
                return null;
        }
    }
}];

export default function TaskTable() {
    const [pageNo, setPageNo] = useState<number>(1);
    const [selectedUserID, setSelectedUserID] = useState<number | undefined>();
    const tasks = useTasks(pageNo, PAGE_SIZE, selectedUserID);
    const me = useContext(UserContext);

    useEffect(() => {
        tasks.mutate();
    }, [pageNo]);

    const getTabs = (task: AppRunTask) => {
        const tabs = ["stdout", "stderr", "hurl"].map(logType => ({
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
        let obj;
        try {
            obj = JSON.parse(task.test_results ?? "")
        } catch (e) {
            obj = {};
        }
        tabs.push({
            key: "report",
            label: 'report.json',
            children: <ReactJson src={obj} theme={"threezerotwofour"}
                style={{ overflow: "scroll" }} collapsed={false} name={null} />
        });
        return tabs;
    }

    return (
        <Card className="card-container"
            title={me?.role === Administrator && <UserSelect onChange={setSelectedUserID} />}
            extra={<Button type="primary" onClick={() => tasks.mutate()}>刷新</Button>}>
            <Table columns={columns} dataSource={tasks.data?.data}
                rowKey="uuid"
                pagination={{
                    pageSize: PAGE_SIZE,
                    total: tasks.data?.total,
                    current: pageNo,
                    onChange: (pageNo: number) => {
                        setPageNo(pageNo);
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
    )
}