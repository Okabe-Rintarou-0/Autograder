import React from 'react';
import { CloudUploadOutlined } from '@ant-design/icons';
import { Layout, Menu, Space, theme } from 'antd';
import { Link, useLocation } from 'react-router-dom';
import { Header } from 'antd/lib/layout/layout';
import NavBar from './navbar';
const { Content, Footer, Sider } = Layout;

export default function BasicLayout({ children }: React.PropsWithChildren) {
    const items = [{
        key: 'submit',
        icon: <CloudUploadOutlined />,
        label: <Link to={'/submit'}> 上传任务 </Link>,
    }];

    const parts = useLocation().pathname.split('/');
    const selectedKeys = [parts[parts.length - 1]];
    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();

    return <Layout>
        <Header><NavBar /></Header>
        <Content style={{ padding: '0 48px' }}>
            <Layout
                style={{ marginTop: "50px", padding: '24px 0', background: colorBgContainer, borderRadius: borderRadiusLG }}
            >
                <Sider style={{ background: colorBgContainer }} width={200} >
                    <Menu defaultSelectedKeys={selectedKeys} items={items} mode="inline" style={{ height: '100%' }} />
                </Sider>
                <Content style={{ padding: '0 24px', minHeight: 280 }}>{children}</Content>
            </Layout>
        </Content>
        <Footer className="footer">
            <Space direction="vertical">
                <Link target="_blank" to="https://github.com/Okabe-Rintarou-0">关于作者</Link>
                <div>电子书城 REINS 2025</div>
            </Space>
        </Footer>
    </Layout>
}