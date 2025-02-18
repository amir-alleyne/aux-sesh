
import { SignInButton, SignedIn, SignedOut, UserButton } from "@clerk/clerk-react";
import Home from "./pages/Home";
import Landing from "./pages/Landing";
import { Layout, Menu, Drawer} from 'antd';
import { Link, Outlet, useLocation, useNavigate } from 'react-router-dom';

import type { MenuProps } from 'antd';
import { useState } from "react";
import {
  AppstoreOutlined,
  ContainerOutlined,
  DesktopOutlined,
  MailOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PieChartOutlined,
} from '@ant-design/icons';
import { Session } from "./pages/Session";
const { Header, Content, Footer, Sider } = Layout;


type MenuItem = Required<MenuProps>['items'][number];

const items: MenuItem[] = [
  { key: '/', icon: <PieChartOutlined />, label: 'Home' },
  { key: '/search', icon: <DesktopOutlined />, label: 'Search For Song',  },

];




export default function App() {
  const [collapsed, setCollapsed] = useState<boolean>(false);
  const location = useLocation();
  const navigate = useNavigate();

 
  const handleMenuClick: MenuProps['onClick'] = (e) => {
    navigate(e.key);
  };
  return (
    <>
    <SignedOut>
          <Landing/>
    </SignedOut>
    <SignedIn>
    <Layout style={{ minHeight: '100vh' }}>
    {/* <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}> */}
     
      <Sider theme="light" collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <Menu  defaultSelectedKeys={['1']} mode="inline" items={items} selectedKeys={[location.pathname]} onClick={handleMenuClick} />
        
        <UserButton/>

      </Sider>
      <Content>
        
       
        <Outlet />
       
      </Content>
      
    </Layout>
    </SignedIn>
    </>

  );
}