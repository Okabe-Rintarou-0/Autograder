import { PrivateLayout } from "../components/layout";
import TaskTable from "../components/task_table";

export default function TaskPage() {
    return (
        <PrivateLayout>
            <TaskTable />
        </PrivateLayout >
    )
}