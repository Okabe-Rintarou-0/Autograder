import { User } from "./user";

export interface Course {
    id: number;
    uuid: string;
    name: string;
    course_code: string;
    enrollments: Enrollment[];
    teachers: Teacher[];
    term: Term;
}

export interface Teacher {
    id: number;
    display_name: string;
}

interface Term {
    id: number;
    name: string;
    start_at?: string | null;
    end_at?: string | null;
    created_at?: string | null;
    workflow_state: string;
}

export type EnrollmentRole = "TaEnrollment" | "StudentEnrollment" | "TeacherEnrollment" | "DesignerEnrollment" | "ObserverEnrollment";

export interface Enrollment {
    type: string;
    role: EnrollmentRole;
    role_id: number;
    user_id: number;
    enrollment_state: string;
}


export interface Assignment {
    id: number;
    key: number;
    needs_grading_count: number | null;
    description: string | null;
    due_at?: string | null;
    unlock_at?: string;
    lock_at?: string;
    points_possible?: number;
    course_id: number;
    name: string;
    html_url: string;
    submission_types: string[];
    allowed_extensions: string[];
    published: boolean;
    has_submitted_submissions: boolean;
    submission?: Submission;
}

export interface AssignmentDate {
    id: number;
    base: boolean;
    title: string;
    due_at: string | null;
    unlock_at: string | null;
    lock_at: string | null;
}

export interface AssignmentOverride {
    id: number;
    pub_id: number;
    assignment_id: number;
    quiz_id: number;
    context_module_id: number;
    student_ids: number[];
    group_id: number;
    course_section_id: number;
    title: string;
    due_at: string | null;
    all_day: boolean;
    all_day_date: string;
    unlock_at: string | null;
    lock_at: string | null;
}

export type WorkflowState = "submitted" | "unsubmitted" | "graded" | "pending_review";

export interface Submission {
    id: number;
    key: number;
    grade: string | null;
    submitted_at?: string;
    assignment_id: number;
    user_id: number;
    late: boolean;
    attachments?: Attachment[];
    workflow_state: WorkflowState;
}

export interface Attachment {
    key: React.Key;
    user?: User;
    user_id: number;
    submitted_at?: string;
    grade: string | null;
    id: number;
    late: boolean;
    uuid: string;
    folder_id: number | null;
    display_name: string;
    filename: string;
    "content-type": string;
    url: string;
    size: number;
    locked: boolean;
    mime_class: string;
    preview_url: string;
}