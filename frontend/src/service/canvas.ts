import useSWR from "swr";
import { Assignment, Course, Submission } from "../model/canvas/course";
import { User } from "../model/canvas/user";
import { fetcher } from "./common";

export function useCourses(shouldFetch: boolean = true) {
    const key = shouldFetch ? "/api/courses" : null;
    return useSWR<Course[]>(key, fetcher);
}

export function useAssignments(courseID: number, shouldFetch: boolean = true) {
    const key = shouldFetch && courseID > 0 ? `/api/assignments?course_id=${courseID}` : null;
    return useSWR<Assignment[]>(key, fetcher);
}

export function useCourseUsers(courseID: number, shouldFetch: boolean = true) {
    const key = shouldFetch && courseID > 0 ? `/api/canvas/users?course_id=${courseID}` : null;
    return useSWR<User[]>(key, fetcher);
}

export function useAssignmentSubmissions(courseID: number, assignmentID: number, shouldFetch: boolean = true) {
    const key = shouldFetch && courseID > 0 && assignmentID > 0 ? `/api/submissions?course_id=${courseID}&assignment_id=${assignmentID}` : null;
    return useSWR<Submission[]>(key, fetcher);
}