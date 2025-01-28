import 'ace-builds/src-noconflict/ace';
import "ace-builds/src-noconflict/ext-language_tools";
import "ace-builds/src-noconflict/mode-mysql";
import "ace-builds/src-noconflict/theme-xcode";
import { Button, Card, Divider, Empty, Space, Table } from "antd";
import useMessage from 'antd/es/message/useMessage';
import { Select } from 'antd/lib';
import { createRef, useMemo, useState } from 'react';
import AceEditor from 'react-ace';
import { SqlResult } from '../model/sql';
import { executeSql } from '../service/sql';

const { Option } = Select;
const databases = ["autograder", "ebookstore"];

export default function SqlConsole() {
    const [database, setDatabase] = useState<string>("autograder");
    const ref = createRef<AceEditor>();
    const [result, setResult] = useState<SqlResult | undefined>();
    const [messageApi, contextHolder] = useMessage();

    const colums = useMemo(() => {
        if (!result || !result.result) return [];

        let obj = result.result[0];
        const cols = [];
        for (let key in obj) {
            cols.push({
                dataIndex: key,
                key,
                title: key,
            });
        }
        return cols;
    }, [result])

    const handleExecSql = async () => {
        const sql = ref.current?.editor.getValue();
        if (!sql) {
            messageApi.error("SQL 语句为空！");
            return;
        }
        try {
            setResult(await executeSql(database, sql));
        } catch (e) {
            console.log(e);
            messageApi.error(`执行出错：${e}`, 0.8);
            setResult(undefined);
        }
    }

    return <Card style={{ width: "100%" }}>
        {contextHolder}
        <Space direction='vertical' style={{ width: "100%" }}>
            <Space>
                <span>选择数据库：</span>
                <Select defaultValue={"autograder"} style={{ width: "200px" }} onChange={setDatabase} >
                    {databases.map(db => <Option value={db} key={db} >{db}</Option>)}
                </Select>
            </Space>
            <AceEditor
                ref={ref}
                mode="mysql"
                theme="xcode"
                setOptions={{
                    enableBasicAutocompletion: true,
                    enableLiveAutocompletion: true,
                }}
                style={{ width: "100%", maxHeight: "100px" }}
            />
            <Button onClick={handleExecSql} type='primary'>执行</Button>
            <Divider />
            {result?.result?.length ?? 0 > 0 ? <Table
                dataSource={result?.result}
                columns={colums}
                scroll={{ x: "100%" }}
            /> : <Empty />}
        </Space>
    </Card>
}