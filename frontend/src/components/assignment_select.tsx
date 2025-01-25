import { Select } from "antd";
import { Assignment } from "../model/canvas/course";

export interface AssignmentSelectProps {
    assignments: Assignment[];
    disabled?: boolean;
    onChange?: (assignmentID: number) => void;
    value?: number;
}

export default function AssignmentSelect({ assignments, disabled, onChange, value }: AssignmentSelectProps) {
    return <Select
        showSearch
        optionFilterProp="children"
        filterOption={(filter, options) => {
            if (!options || !filter) {
                return true;
            }
            let assignmentID = options.value;
            let assignment = assignments.find(assignment => assignment.id === assignmentID)!;
            // filter by course name or teacher name
            return assignment.name.includes(filter);
        }}
        value={value}
        style={{ width: 350 }}
        disabled={disabled}
        onChange={onChange}
        placeholder={"请选择作业"}
        options={assignments.map(assignment => ({
            label: assignment.name,
            value: assignment.id
        }))}
    />
}