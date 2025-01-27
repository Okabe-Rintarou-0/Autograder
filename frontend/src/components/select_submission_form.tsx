import { Space } from "antd";
import { useState } from "react";
import { ModalChildrenProps } from "../lib/hooks";
import { Attachment } from "../model/canvas/course";
import { useAssignmentSubmissions, useAssignments, useCourses } from "../service/canvas";
import AssignmentSelect from "./assignment_select";
import CourseSelect from "./course_select";
import SubmissionTable from "./submission_table";

interface SelectSubmissionFormProps {
    onSubmit: (attachment: Attachment) => void;
}

export default function SelectSubmissionForm(props: ModalChildrenProps<SelectSubmissionFormProps>) {
    const [selectedCourseID, setSelectedCourseID] = useState<number>(0);
    const [selectedAssignmentID, setSelectedAssignmentID] = useState<number>(0);
    const courses = useCourses(props.isOpen);
    const assignments = useAssignments(selectedCourseID, props.isOpen);
    const submissions = useAssignmentSubmissions(selectedCourseID, selectedAssignmentID, props.isOpen);

    return <Space direction="vertical" size={"large"} style={{ width: "100%" }}>
        <CourseSelect courses={courses.data ?? []} onChange={setSelectedCourseID} />
        <AssignmentSelect assignments={assignments.data ?? []} onChange={setSelectedAssignmentID} />
        <SubmissionTable courseID={selectedCourseID} assignmentID={selectedAssignmentID}
            submissions={submissions.data ?? []} isLoading={submissions.isLoading}
            onSubmit={props.onSubmit}
        />
    </Space>
}