import { Button, Form, Modal } from "antd";
import { ImportCanvasUsersRequest, importCanvasUsersURL } from "../model/user";
import useMessage from "antd/es/message/useMessage";
import { submit } from "../service/common";
import { BaseResp } from "../model/resp";
import { handleBaseResp } from "../utils/handle_resp";
import CourseSelect from "./course_select";
import { useEffect, useState } from "react";
import { Course } from "../model/canvas/course";
import { listCourses } from "../service/canvas";

interface ImportCanvasUsersModal {
    open: boolean;
    onOk: () => void;
    onCancel: () => void;
}

export default function ImportCanvasUsersModal({ open, onOk, onCancel }: ImportCanvasUsersModal) {
    const [form] = Form.useForm<ImportCanvasUsersRequest>();
    const [messageApi, contextHolder] = useMessage();
    const [courses, setCourses] = useState<Course[]>([]);

    useEffect(() => {
        const getCourses = async () => {
            try {
                const courses = await listCourses();
                setCourses(courses);
            } catch (e) {
                console.log(e);
            }
        }
        getCourses();
    }, [])

    const handleSubmit = async (request: ImportCanvasUsersRequest) => {
        console.log(request);
        let resp = await submit<BaseResp>(importCanvasUsersURL, request);
        handleBaseResp(messageApi, resp, onOk);
    };

    return (
        <Modal
            destroyOnClose
            title={"从 Canvas 导入"}
            open={open}
            onOk={onOk}
            onCancel={onCancel}
            footer={null}
            width={800}
        >
            {contextHolder}
            <Form
                form={form}
                layout="vertical"
                onFinish={handleSubmit}
                preserve={false}
            >
                <Form.Item
                    name="course_id"
                    label="课程"
                    required
                    rules={[{ required: true }]}
                >
                    <CourseSelect courses={courses} />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        提交
                    </Button>
                </Form.Item>
            </Form>
        </Modal>
    );
}