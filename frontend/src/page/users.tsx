import { PrivateLayout } from "../components/layout";
import { UserTable } from "../components/user_table";
import { Administrator } from "../model/user";

export default function UsersPage() {
    return (
        <PrivateLayout forRole={Administrator}>
            <UserTable />
        </PrivateLayout >
    )
}