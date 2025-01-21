import { Button, Result } from "antd";
import { useNavigate } from "react-router-dom";
import { BasicLayout } from "../components/layout";

export default function NotfoundPage() {
    const navigate = useNavigate();

    return <BasicLayout noSider>
        <Result
            status="404"
            title="404"
            subTitle="很抱歉，您请求的页面不存在。"
            extra={< Button type="primary" onClick={() => navigate("/")}> 返回主页</Button>}
        />
    </BasicLayout>
}