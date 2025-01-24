import { Badge, Button, Card, Select, Space, Spin, Table, Tabs, Tag, Tooltip } from "antd";
import { PrivateLayout } from "../components/layout";
import { useTasks } from "../service/task";
import useMessage from "antd/es/message/useMessage";
import { useEffect, useState } from "react";
import { LazyLog } from "@melloware/react-logviewer";
import { AppRunTask } from "../model/app";
import { AppRunTaskStatusFail, AppRunTaskStatusRunning, AppRunTaskStatusSucceed, AppRunTaskStatusWaiting } from "../model/app";
import { BadgeProps } from "antd/lib";
import ReactJson from "react-json-view-ts";
import { formatDate } from "../utils/time";
import { useDebounceFn } from "ahooks";
import { listUsers } from "../service/user";
import { User } from "../model/user";

const columns = [{
    title: '创建者',
    dataIndex: 'real_name',
    key: 'real_name',
    render: (real_name: string, task: AppRunTask) => {
        return <Tooltip title={`${task.username}\n${task.email}`}>
            <Tag color="blue">{real_name}</Tag>
        </Tooltip>
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

export default function TaskPage() {
    const pageSize = 20;
    const [messageApi, contextHolder] = useMessage();
    const [fetching, setFetching] = useState<boolean>(false);
    const [pageNo, setPageNo] = useState<number>(1);
    const [users, setUsers] = useState<User[]>([]);
    const [loadingMore, setLoadingMore] = useState(false);
    const [userPageNo, setUserPageNo] = useState<number>(1);
    const [userKeyword, setUserKeyword] = useState<string>("");
    const [userHasNextPage, setUserHasNextPage] = useState<boolean>(false);
    const tasks = useTasks(pageNo, pageSize);

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

    const { run: debounceFetcher } = useDebounceFn(
        (keyword: string) => {
            keyword = keyword.trim();
            console.log(keyword)
            if (keyword) {
                setFetching(true);
                setUserKeyword(keyword);
                setUserPageNo(1);
                listUsers(keyword, 1, pageSize).then((resp) => {
                    if (resp.data) {
                        setUserHasNextPage((userPageNo - 1) * pageSize + resp.data.length < resp.total);
                        setUsers(resp.data);
                        setUserPageNo(userPageNo => userPageNo + 1);
                    } else {
                        setUsers([]);
                        setUserHasNextPage(false);
                    }
                    setFetching(false);
                });
            }
        },
        {
            wait: 800,
        }
    );

    const loadMore = () => {
        if (!userHasNextPage) {
            return;
        }
        if (userKeyword) {
            setLoadingMore(true);
            listUsers(userKeyword, userPageNo, pageSize).then((resp) => {
                if (resp.data) {
                    setUserHasNextPage((userPageNo - 1) * pageSize + resp.data.length < resp.total);
                    setUsers(users.concat(resp.data));
                    setUserPageNo(userPageNo => userPageNo + 1);
                }
                setLoadingMore(false);
            });
        }
    }

    const onPopupScroll = (e: { target?: any }) => {
        const { target } = e;
        if (target.scrollTop + target.offsetHeight >= target.scrollHeight) {
            loadMore();
        }
    };

    return (
        <PrivateLayout>
            {contextHolder}
            <Card className="card-container"
                title={<Select
                    style={{ width: "200px" }}
                    showSearch
                    placeholder="指定用户"
                    filterOption={false}
                    onSearch={debounceFetcher}
                    onPopupScroll={onPopupScroll}
                    notFoundContent={fetching ? <Spin size="small" /> : null}
                    dropdownRender={(menu) => (
                        <>
                            {menu}
                            {loadingMore ? (
                                <Spin size="small" style={{ textAlign: 'center' }} />
                            ) : null}
                        </>
                    )}>
                    {users.map((user) => {
                        return <Select.Option key={user.id} value={user.id}>
                            {`${user.real_name}(${user.username})`}
                        </Select.Option>
                    })}
                </Select>}
                extra={<Button type="primary" onClick={() => tasks.mutate()}>刷新</Button>}>
                <Table columns={columns} dataSource={tasks.data?.data}
                    rowKey="uuid"
                    pagination={{
                        pageSize,
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
        </PrivateLayout >
    )
}