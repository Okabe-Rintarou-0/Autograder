import { PrivateLayout } from "../components/layout";
import { TestcaseTable } from "../components/testcase_table";
import { Administrator } from "../model/user";

export default function TestcasesPage() {
    return (
        <PrivateLayout forRole={Administrator}>
            <TestcaseTable />
        </PrivateLayout >
    )
}