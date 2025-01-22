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