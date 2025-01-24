import { useMemoizedFn } from "ahooks";
import { Button, Form } from "antd";
import useMessage from "antd/es/message/useMessage";
import { ModalChildrenProps } from "../lib/hooks";
import { BaseResp } from "../model/resp";
import { ImportCanvasUsersRequest, importCanvasUsersURL } from "../model/user";
import { useCourses } from "../service/canvas";
import { submit } from "../service/common";
import { handleBaseResp } from "../utils/handle_resp";
import CourseSelect from "./course_select";

export default function ImportCanvasUsersForm(props: ModalChildrenProps) {
    const [form] = Form.useForm<ImportCanvasUsersRequest>();
    const [messageApi, contextHolder] = useMessage();
    const courses = useCourses(props.isOpen);
    const handleSubmit = useMemoizedFn(async (request: ImportCanvasUsersRequest) => {
        console.log(request);
        let resp = await submit<BaseResp>(importCanvasUsersURL, request);
        handleBaseResp(messageApi, resp, props.close);
    });

    return <Form
        form={form}
        layout="vertical"
        onFinish={handleSubmit}
        preserve={false}
    >
        {contextHolder}
        <Form.Item
            name="course_id"
            label="课程"
            required
            rules={[{ required: true }]}
        >
            <CourseSelect courses={courses.data ?? []} />
        </Form.Item>
        <Form.Item>
            <Button type="primary" htmlType="submit">
                提交
            </Button>
        </Form.Item>
    </Form >
}