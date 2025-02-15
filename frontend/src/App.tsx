import AppRouter from "./components/router";
import { ConfigProvider } from 'antd';
import zhCN from 'antd/locale/zh_CN';
import "./css/global.css"
import React from "react";

function App() {
  return <ConfigProvider locale={zhCN}>
    <AppRouter />
  </ConfigProvider>
}

export default App;
