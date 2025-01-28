import { PrivateLayout } from "../components/layout";
import SqlConsole from "../components/sql_console";
import { Administrator } from "../model/user";

export default function SqlPage() {
    return <PrivateLayout forRole={Administrator}>
        <SqlConsole />
    </PrivateLayout>
}