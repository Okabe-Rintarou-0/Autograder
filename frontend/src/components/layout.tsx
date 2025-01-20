import React, { useEffect, useState } from 'react';
import { CloudUploadOutlined, UnorderedListOutlined } from '@ant-design/icons';
import { Layout, Menu, Space, theme } from 'antd';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { Header } from 'antd/lib/layout/layout';
import NavBar from './navbar';
import { User } from '../model/user';
import { getMe } from '../service/user';
const { Content, Footer, Sider } = Layout;

export interface BasicLayoutProps {
    noSider?: boolean;
    me?: User;
}

export function BasicLayout({ children, noSider, me }: React.PropsWithChildren<BasicLayoutProps>) {
    const items = [{
        key: 'submit',
        icon: <CloudUploadOutlined />,
        label: <Link to={'/submit'}> 上传任务 </Link>,
    }, {
        key: 'tasks',
        icon: <UnorderedListOutlined />,
        label: <Link to={'/tasks'}> 查看任务 </Link>,
    }];

    const parts = useLocation().pathname.split('/');
    const selectedKeys = [parts[parts.length - 1]];
    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();

    return <Layout>
        <Header><NavBar me={me} /></Header>
        <Content style={{ padding: '0 48px' }}>
            <Layout
                style={{ marginTop: "50px", padding: '24px 0', background: colorBgContainer, borderRadius: borderRadiusLG }}
            >
                {!noSider && <Sider style={{ background: colorBgContainer }} width={200} >
                    <Menu defaultSelectedKeys={selectedKeys} items={items} mode="inline" style={{ height: '100%' }} />
                </Sider>}
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

export function PrivateLayout({ children, noSider }: React.PropsWithChildren<BasicLayoutProps>) {
    const [me, setMe] = useState<User | undefined>();

    useEffect(() => {
        const fetchMe = async () => {
            try {
                const me = await getMe();
                setMe(me);
            } catch (e) {
                console.log(e);
                const navigate = useNavigate();
                navigate("/login");
            }
        }
        fetchMe();
    }, []);

    if (!me) {
        return null;
    }
    return <BasicLayout noSider={noSider} me={me}>{children}</BasicLayout>
}

