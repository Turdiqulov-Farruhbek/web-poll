import React, { useEffect, useState } from "react";
import { Button, Layout, Menu, theme } from "antd";
import Logo from "@images/Logo.jpg";
import "./style.scss";
import http from "../config";
import { Outlet, useNavigate } from "react-router-dom";
import { MenuIds } from "@store";
import {
  EnvironmentOutlined,
  LogoutOutlined,
  MailOutlined,
} from "@ant-design/icons";
import { getCookies, removeCookies } from "../utils/cocies";
import Sider from "antd/es/layout/Sider";

const { Header, Content, Footer } = Layout;

const App: React.FC = () => {
  const [data, setData] = useState<any[]>([]);
  const { changeMenu_id }: any = MenuIds();
  const navigate = useNavigate();
  const {
    token: { borderRadiusLG },
  } = theme.useToken();

  const siderStyle: React.CSSProperties = {
    height: "100vh",
    position: "fixed",
    insetInlineStart: 0,
    top: 0,
    bottom: 0,
    scrollbarWidth: "thin",
    scrollbarColor: "unset",
    marginTop: 100,
    background: "#FFF",
    zIndex: 99,
    paddingBottom: 150,
  };

  async function getMenuData() {
    try {
      const response = await http.get("/polls");
      setData(response?.data?.poll || []);
    } catch (err) {
      console.error("Error fetching menu data:", err);
    }
  }

  function logout() {
    removeCookies("access_token");
    navigate("/");
  }

  async function getUserData() {
    try {
      const response = await http.get("/profile");
      localStorage.setItem("user_id", response?.data?.id || "");
    } catch (err) {
      console.error("Error fetching user data:", err);
    }
  }

  useEffect(() => {
    const token = getCookies("access_token");
    if (!token) {
      navigate("/");
    }
    getMenuData();
    getUserData();
  }, [navigate]);

  return (
    <Layout style={{ minHeight: "100vh" }}>
      <Header
        className="header"
        style={{
          position: "fixed",
          zIndex: 999,
          width: "100%",
          background: "#0C3B2E",
        }}
      >
        <div
          onClick={() => navigate("/dashboard")}
          style={{ cursor: "pointer" }}
          className="demo-logo"
        >
          <img src={Logo} alt="" />
          <p>Researchpsy</p>
        </div>
        <div style={{ display: "flex", alignItems: "center", gap: 10 }}>
          <Button
            onClick={logout}
            className="logout"
            style={{ display: "flex", alignItems: "center" }}
          >
            <LogoutOutlined />
            <p>Chiqish</p>
          </Button>
        </div>
      </Header>
      <Layout>
        <Sider
          style={siderStyle}
          breakpoint="xl"
          width={250}
          collapsedWidth="0"
          onBreakpoint={(broken) => {
            console.log(broken);
          }}
          onCollapse={(collapsed, type) => {
            console.log(collapsed, type);
          }}
        >
          <Menu
            style={{
              height: "100%",
              fontSize: 18,
              fontWeight: "semibold",
              background: "white",
              color: "white",
            }}
            items={data.map((e: any, i: number) => ({
              key: i,
              onClick: () => changeMenu_id(e.id),
              label: e.title,
            }))}
          />
        </Sider>
        <Content
          style={{
            width: "100%",
            maxWidth: "1200px",
            margin: "0 auto",
            padding: "0 50px",
            paddingTop: 100,
            paddingBottom: 160,
            color: "black",
          }}
        >
          <div
            style={{
              marginTop: 20,
              background: "white",
              minHeight: 600,
              padding: 24,
              borderRadius: borderRadiusLG,
            }}
          >
            <Outlet />
          </div>
        </Content>
      </Layout>
      <Footer
        style={{
          textAlign: "center",
          position: "fixed",
          bottom: 0,
          left: 0,
          width: "100%",
          zIndex: 999999,
        }}
      >
        <div className="footer-wrapper">
          <div className="footer-one">
            <div>
              <img src={Logo} alt="LOGO" />
              <p>Researchpsy</p>
            </div>
            <div className="adress">
              <a
                href="mailto:nilukhanphd1@gmail.com"
                style={{
                  display: "flex",
                  gap: "5px",
                  alignItems: "center",
                  textDecoration: "none",
                  color: "inherit",
                }}
              >
                <MailOutlined style={{ fontSize: "20px", color: 'white' }} />
                <p style={{ fontSize: "12px", margin: 0 }}>
                  nilukhanphd1@gmail.com
                </p>
              </a>
              
              <p style={{fontSize: 12, display: 'flex', gap: 5, alignItems: 'center'}}>
              <EnvironmentOutlined style={{fontSize: 20}}/>
              International Nordic University, Sebzor 22A
              </p>
            </div>
          </div>
        </div>
      </Footer>
    </Layout>
  );
};

export default App;
