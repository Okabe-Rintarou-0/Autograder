import axios from "axios";
import { SqlResult } from "../model/sql";

export async function executeSql(database: string, sql: string) {
    const data = {
        database,
        sql,
    }
    const resp = await axios.post<SqlResult>("/db/execute", data, {
        headers: {
            'Content-Type': 'application/json'
        }
    });
    return resp.data;
}