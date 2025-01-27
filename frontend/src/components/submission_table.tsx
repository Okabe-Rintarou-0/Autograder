import { useMemoizedFn } from "ahooks";
import { Button, Space, Table, Tag } from "antd";
import { useMemo, useState } from "react";
import { Attachment, Submission } from "../model/canvas/course";
import { User } from "../model/canvas/user";
import { useCourseUsers } from "../service/canvas";
import { formatDate } from "../utils/time";

interface SubmissionTableProps {
    courseID: number;
    assignmentID: number;
    submissions: Submission[];
    isLoading: boolean;
    onSubmit: (attachment: Attachment) => void;
}

export default function SubmissionTable(props: SubmissionTableProps) {
    const { courseID, assignmentID, submissions, isLoading, onSubmit } = props;
    const users = useCourseUsers(courseID);
    const [selectedAttachment, setSelectedAttachment] = useState<Attachment | undefined>();
    const usersMap = useMemo(() => {
        const m = new Map<number, User>();
        if (users.data) {
            for (let user of users.data) {
                m.set(user.id, user);
            }
        }
        return m;
    }, [users.data]);
    const columns = [{
        title: '学生',
        dataIndex: 'user',
        key: 'user',
        render: (user: User | undefined) => user?.name
    }, {
        title: '文件',
        dataIndex: 'display_name',
        key: 'display_name',
        render: (name: string, attachment: Attachment) => <a href={`https://oc.sjtu.edu.cn/courses/${courseID}/gradebook/speed_grader?assignment_id=${assignmentID}&student_id=${attachment.user_id}`}
            target="_blank"
        >
            {name}
        </a>
    }, {
        title: '提交时间',
        dataIndex: 'submitted_at',
        key: 'submitted_at',
        render: formatDate,
    }, {
        title: '状态',
        dataIndex: 'late',
        key: 'late',
        render: (late: boolean) => late ? <Tag color="red">迟交</Tag> : <Tag color="green">按时提交</Tag>
    }];

    const attachments = useMemo(() => {
        let attachments: Attachment[] = [];
        for (let submission of submissions) {
            let thisAttachments = submission.attachments;
            if (!thisAttachments) {
                continue;
            }
            for (let attachment of thisAttachments) {
                attachment.user = usersMap.get(submission.user_id);
                attachment.user_id = submission.user_id;
                attachment.submitted_at = submission.submitted_at;
                attachment.grade = submission.grade;
                attachment.key = attachment.id;
                attachment.late = submission.late;
            }
            attachments.push(...thisAttachments);
        }
        return attachments;
    }, [submissions]);

    const handleSubmit = useMemoizedFn(() => {
        if (selectedAttachment) {
            onSubmit(selectedAttachment)
        }
    });

    return <Space style={{ width: "100%" }} direction="vertical">
        <Table style={{ width: "100%" }}
            columns={columns}
            loading={isLoading}
            dataSource={attachments}
            pagination={{
                pageSize: 5,
            }}
            rowSelection={{
                type: 'radio',
                onChange: (_, rows) => {
                    if (rows.length > 0) {
                        setSelectedAttachment(rows[0]);
                    } else {
                        setSelectedAttachment(undefined);
                    }
                }
            }}
        />
        <Button type="primary" disabled={!selectedAttachment} onClick={handleSubmit}>选择</Button>
    </Space>
}