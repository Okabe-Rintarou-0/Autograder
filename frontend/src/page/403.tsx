import { Button, Result } from "antd";
import { useNavigate } from "react-router-dom";
import { BasicLayout } from "../components/layout";

export default function UnauthorizedPage() {
    const navigate = useNavigate();

    return <BasicLayout noSider>
        <Result
            status="403"
            title="403"
            subTitle="很抱歉，您没有访问该页面的权限。"
            extra={< Button type="primary" onClick={() => navigate("/")}> 返回主页</Button>}
        />
    </BasicLayout>
}