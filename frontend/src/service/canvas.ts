import useSWR from "swr";
import { fetcher } from "./common";
import { Course } from "../model/canvas/course";

export function useCourses(shouldFetch: boolean = true) {
    const key = shouldFetch ? "/api/courses" : null;
    return useSWR<Course[]>(key, fetcher);
}