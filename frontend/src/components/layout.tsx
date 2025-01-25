import { CloudUploadOutlined, CodeOutlined, UnorderedListOutlined, UserOutlined } from '@ant-design/icons';
import { Layout, Menu, Space, theme } from 'antd';
import { Header } from 'antd/lib/layout/layout';
import React, { useEffect, useState } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { UserContext } from '../lib/context';
import { Administrator, User } from '../model/user';
import { getMe } from '../service/user';
import NavBar from './navbar';
const { Content, Footer, Sider } = Layout;

export interface BasicLayoutProps {
    noSider?: boolean;
    me?: User;
}

export interface PrivateLayoutProps extends BasicLayoutProps {
    forRole?: number
}

export function BasicLayout({ children, noSider, me }: React.PropsWithChildren<BasicLayoutProps>) {
    const getItems = () => {
        const items = [{
            key: 'submit',
            icon: <CloudUploadOutlined />,
            label: <Link to={'/submit'}> 上传任务 </Link>,
        }, {
            key: 'tasks',
            icon: <UnorderedListOutlined />,
            label: <Link to={'/tasks'}> 查看任务 </Link>,
        }];

        if (me?.role === Administrator) {
            const adminItems = [{
                key: 'users',
                icon: <UserOutlined />,
                label: <Link to={'/users'}> 查看用户 </Link>,
            }, {
                key: 'testcases',
                icon: <CodeOutlined />,
                label: <Link to={'/testcases'}> 测试用例 </Link>,
            }];
            items.push(...adminItems);
        }
        return items;
    }
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
                    <Menu defaultSelectedKeys={selectedKeys} items={getItems()} mode="inline" style={{ height: '100%' }} />
                </Sider>}
                <Content style={{ padding: '0 24px', minHeight: 280 }}>
                    <UserContext.Provider value={me}>
                        {children}
                    </UserContext.Provider>
                </Content>
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

export function PrivateLayout({ forRole, children, noSider }: React.PropsWithChildren<PrivateLayoutProps>) {
    const [me, setMe] = useState<User | undefined>();
    const navigate = useNavigate();
    const validateUser = (user: User) => {
        if (forRole && user.role !== forRole) {
            return false;
        }
        return true;
    }

    useEffect(() => {
        const fetchMe = async () => {
            try {
                const me = await getMe();
                setMe(me);
                if (!validateUser(me)) {
                    navigate("/403");
                }
            } catch (e) {
                console.log(e);
                navigate("/login");
            }
        }
        fetchMe();
    }, []);

    return <BasicLayout noSider={noSider} me={me}>{me && children}</BasicLayout>
}

