import { useMemoizedFn } from "ahooks";
import { Button, Card, Form, Select, UploadFile } from "antd";
import useMessage from "antd/es/message/useMessage";
import { RcFile } from "antd/es/upload";
import { PrivateLayout } from "../components/layout";
import SelectSubmissionForm from "../components/select_submission_form";
import { ZipUpload } from "../components/upload";
import { useModal } from "../lib/hooks";
import { AppInfo } from "../model/app";
import { Attachment } from "../model/canvas/course";
import { submitApp } from "../service/task";
import { urlToFile } from "../utils/file";
import { handleBaseResp } from "../utils/handle_resp";

const supportedJDKVersions = [11, 17]
const { Option } = Select;

export default function SubmitPage() {
    const [form] = Form.useForm<AppInfo>();
    const [messageApi, contextHolder] = useMessage();
    const onSubmit = useMemoizedFn(async (attachment: Attachment) => {
        console.log("attach", attachment);
        const file = await urlToFile(attachment.url, attachment.display_name);
        const uploadFile: UploadFile = {
            uid: attachment.display_name,
            name: attachment.display_name,
            url: attachment.url,
            originFileObj: file as RcFile,
        }
        form.setFieldValue("file", [uploadFile]);
        closeSelectSubmissionModal();
    })
    const {
        modal: SelectSubmissionModal,
        open: openSelectSubmissionModal,
        close: closeSelectSubmissionModal,
    } = useModal(SelectSubmissionForm, { onSubmit })
    const handleSubmit = async (appInfo: AppInfo) => {
        console.log(appInfo)
        form.setFieldValue("file", []);
        try {
            const resp = await submitApp(appInfo);
            handleBaseResp(messageApi, resp);
        } catch (e) {
            console.log(e);
            messageApi.error(`上传失败：${e}`);
        }
    }

    return (
        <PrivateLayout>
            {contextHolder}
            <SelectSubmissionModal destroyOnClose title="从 Canvas 中导入"
                footer={null} height={800} width={"80%"} onCancel={closeSelectSubmissionModal}
            />
            <Card className="card-container"
                title={<Button type="primary" onClick={openSelectSubmissionModal}>从 Canvas 中导入</Button>}
            >
                <Form
                    form={form}
                    layout="vertical"
                    onFinish={handleSubmit}
                    preserve={false}
                >
                    <Form.Item
                        key="jdkVersion"
                        name="jdkVersion"
                        label="JDK 版本"
                        rules={[{ required: true }]}
                    >
                        <Select placeholder="请选择您的 JDK 版本">
                            {supportedJDKVersions.map(version => <Option value={version} key={version}>{`JDK ${version}`}</Option>)}
                        </Select>
                    </Form.Item>
                    <Form.Item
                        key="authenticationType"
                        name="authenticationType"
                        label="鉴权方式"
                        rules={[{ required: true }]}
                    >
                        <Select placeholder="请选择您的鉴权方式">
                            <Option value={1}>Cookies</Option>
                            <Option value={2}>Token</Option>
                        </Select>
                    </Form.Item>
                    <Form.Item
                        key="file"
                        name="file"
                        label="上传项目压缩包（zip）"
                        rules={[{ required: true }]}
                    >
                        <ZipUpload />
                    </Form.Item>

                    <Form.Item>
                        <Button type="primary" htmlType="submit">
                            提交
                        </Button>
                    </Form.Item>
                </Form>
            </Card>
        </PrivateLayout>
    );
}
